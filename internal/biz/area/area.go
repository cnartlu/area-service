package area

import (
	"context"
	"strconv"
	"strings"

	"github.com/cnartlu/area-service/errors"
	"github.com/cnartlu/area-service/internal/biz/transaction"
	pkgsort "github.com/cnartlu/area-service/pkg/data/sort"
	"github.com/mozillazg/go-pinyin"
)

type AreaRepo interface {
	Count(ctx context.Context, options ...Query) int
	// FindList 查找数据列表
	FindList(ctx context.Context, options ...Query) ([]*Area, error)
	// FindOne 查找数据
	FindOne(ctx context.Context, options ...Query) (*Area, error)
	// Save 新增或保存数据
	Save(ctx context.Context, data *Area) (*Area, error)
	// Remove 移除数据
	Remove(ctx context.Context, options ...Query) error
	// ReplaceParentListPrefix 替换父级列表前缀
	ReplaceParentListPrefix(ctx context.Context, oldPrefix, newPrefix string) (int, error)
}

type AreaUsecase struct {
	repo        AreaRepo
	transaction transaction.Transaction
}

// List 列表更新
func (m *AreaUsecase) List(ctx context.Context, params FindListParam) ([]*Area, error) {
	options := []Query{}
	if params.RegionID != "" {
		p, err := m.FindByRegionID(ctx, params.RegionID, params.Level)
		if err != nil {
			return nil, err
		}
		params.ParentID = p.ID
	}
	options = append(options, ParentIDEQ(params.ParentID))
	if params.Keyword != "" {
		options = append(options, TitleContains(params.Keyword))
	}
	if params.Order != "" {
		options = append(options, Order(pkgsort.ParseArray(params.Order)...))
	}
	return m.repo.FindList(ctx, options...)
}

// FindOne 查询ID值等价
func (m *AreaUsecase) FindOne(ctx context.Context, id uint64) (*Area, error) {
	return m.repo.FindOne(ctx, IDEQ(id))
}

func (m *AreaUsecase) FindByRegionID(ctx context.Context, regionID string, level int) (*Area, error) {
	options := []Query{RegionIDEQ(regionID)}
	if level > 0 {
		options = append(options, LevelEQ(level))
	}
	return m.repo.FindOne(ctx, options...)
}

func (m *AreaUsecase) Create(ctx context.Context, data CreateParam) (*Area, error) {
	if data.Title == "" {
		return nil, errors.ErrorParamMissing("the title can not be blank")
	}

	var (
		level      int = 1
		parentList     = "0"
		areaModel  *Area
		parentArea *Area
	)

	{
		if data.ParentID != 0 {
			var err error
			parentArea, err = m.FindOne(ctx, data.ParentID)
			if err != nil {
				return nil, err
			}
			level += parentArea.Level
			parentList = parentArea.ParentList
		}

		if _, err := m.FindByRegionID(ctx, data.RegionID, level); err == nil {
			return nil, errors.ErrorParamFormat("identify exists")
		} else if !errors.IsDataNotFound(err) {
			return nil, err
		}
	}

	{
		py := pinyin.LazyConvert(data.Title, &pinyin.Args{
			Style: pinyin.NORMAL,
			Fallback: func(r rune, a pinyin.Args) []string {
				if r == 0 {
					return []string{}
				}
				return []string{string(r)}
			},
		})
		pyStr := strings.Join(py, " ")
		areaModel = &Area{}
		areaModel.ParentID = data.ParentID
		areaModel.RegionID = data.RegionID
		areaModel.ParentList = parentList
		areaModel.Title = data.Title
		areaModel.Pinyin = pyStr
		areaModel.Ucfirst = pyStr[0:1]
		areaModel.Lat = data.Lat
		areaModel.Lng = data.Lng
		areaModel.CityCode = data.CityCode
		areaModel.ZipCode = data.ZipCode
		areaModel.Level = level

		err := m.transaction.Transaction(ctx, func(ctx context.Context) error {
			var err error
			areaModel, err = m.repo.Save(ctx, areaModel)
			if err != nil {
				return err
			}
			areaModel.ParentList += "," + strconv.FormatUint(areaModel.ID, 10)
			areaModel, err = m.repo.Save(ctx, areaModel)
			if err != nil {
				return err
			}
			if areaModel.ParentID > 0 {
				parentArea.ChildrenNumber++
				parentArea, err = m.repo.Save(ctx, parentArea)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return areaModel, nil
}

func (m *AreaUsecase) Update(ctx context.Context, data UpdateParam) (*Area, error) {
	if data.ID == 0 {
		return nil, errors.ErrorParamMissing("unique identifier for missing data")
	}

	if data.Title == "" {
		return nil, errors.ErrorParamMissing("the title can not be blank")
	}

	areaModel, err := m.FindOne(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	var (
		level        = areaModel.Level
		parentList   = areaModel.ParentList
		oldAreaModel = areaModel
	)

	{
		// 父级 发生变更
		if areaModel.ParentID != data.ParentID {
			if data.ParentID != 0 {
				parentArea, err := m.FindOne(ctx, data.ParentID)
				if err != nil {
					return nil, err
				}
				level = parentArea.Level + 1
				parentList = parentArea.ParentList + "," + strconv.FormatUint(areaModel.ID, 10)
			}
		}
	}

	// 区域ID 发生变更
	if areaModel.RegionID != data.RegionID {
		if _, err := m.FindByRegionID(ctx, data.RegionID, level); err == nil {
			return nil, errors.ErrorParamFormat("identify exists")
		} else if !errors.IsDataNotFound(err) {
			return nil, err
		}
	}

	{
		// 更新数据
		py := pinyin.LazyConvert(data.Title, nil)
		pyStr := strings.Join(py, " ")
		areaModel.ParentID = data.ParentID
		areaModel.RegionID = data.RegionID
		areaModel.ParentList = parentList
		areaModel.Title = data.Title
		areaModel.Pinyin = pyStr
		areaModel.Ucfirst = pyStr[0:1]
		areaModel.Lat = data.Lat
		areaModel.Lng = data.Lng
		areaModel.CityCode = data.CityCode
		areaModel.ZipCode = data.ZipCode
		areaModel.Level = level
		err := m.transaction.Transaction(ctx, func(ctx context.Context) error {
			var err error
			areaModel, err = m.repo.Save(ctx, areaModel)
			if err != nil {
				return err
			}
			// 父级ID变更，则需要同步更新其子集的parentList字段
			if oldAreaModel.ParentID != areaModel.ParentID {
				if oldAreaModel.ParentID > 0 {
					oldParentModel, err := m.FindOne(ctx, oldAreaModel.ParentID)
					if err != nil {
						return err
					}
					oldParentModel.ChildrenNumber -= 1
					if _, err := m.repo.Save(ctx, oldParentModel); err != nil {
						return err
					}
				}

				if areaModel.ParentID > 0 {
					parentArea, err := m.FindOne(ctx, areaModel.ParentID)
					if err != nil {
						return err
					}
					parentArea.ChildrenNumber += 1
					if _, err := m.repo.Save(ctx, parentArea); err != nil {
						return err
					}
				}

				if areaModel.ChildrenNumber > 0 {
					if _, err := m.repo.ReplaceParentListPrefix(ctx, oldAreaModel.ParentList, areaModel.ParentList); err != nil {
						return err
					}
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}

	return areaModel, nil
}

// Delete 删除值
func (m *AreaUsecase) Delete(ctx context.Context, ids ...uint64) error {
	return m.repo.Remove(ctx, IDIn(ids...))
}

func NewAreaUsecase(repo AreaRepo) *AreaUsecase {
	return &AreaUsecase{
		repo: repo,
	}
}
