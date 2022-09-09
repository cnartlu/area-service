package area

import (
	"context"

	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/go-redis/redis/v8"
)

type AreaRepo struct {
	ent *ent.Client
	rds *redis.Client
}

var _ bizArea.Manager = (*AreaRepo)(nil)

// Count 数量
func (r *AreaRepo) Count(ctx context.Context, options ...bizArea.Option) int {
	query := r.ent.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	return query.CountX(ctx)
}

// FindList 查找数据
func (r *AreaRepo) FindList(ctx context.Context, options ...bizArea.Option) (list []*bizArea.Area, err error) {
	query := r.ent.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	results, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		list = append(list, &bizArea.Area{
			ID:       result.ID,
			Title:    result.Title,
			Pinyin:   result.Pinyin,
			Ucfirst:  result.Ucfirst,
			CityCode: result.CityCode,
			ZipCode:  result.ZipCode,
		})
	}
	return list, nil
}

// FindList 查找数据
func (r *AreaRepo) FindOne(ctx context.Context, options ...bizArea.Option) (*bizArea.Area, error) {
	query := r.ent.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	data := &bizArea.Area{
		ID:       result.ID,
		Title:    result.Title,
		Pinyin:   result.Pinyin,
		Ucfirst:  result.Ucfirst,
		CityCode: result.CityCode,
		ZipCode:  result.ZipCode,
	}
	return data, err
}

// FindList 查找数据
func (r *AreaRepo) Save(ctx context.Context, x *bizArea.Area) (*bizArea.Area, error) {
	query := r.ent.Area.Query()
	// if len(options) > 0 {
	// 	q := newOption(query)
	// 	for _, option := range options {
	// 		option(q)
	// 	}
	// }
	result, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	data := &bizArea.Area{
		ID:       result.ID,
		Title:    result.Title,
		Pinyin:   result.Pinyin,
		Ucfirst:  result.Ucfirst,
		CityCode: result.CityCode,
		ZipCode:  result.ZipCode,
	}
	return data, err
}

// Remove 移除数据
func (r *AreaRepo) Remove(ctx context.Context, options ...bizArea.Option) error {
	query := r.ent.Area.Query()
	if len(options) > 0 {
		q := newOption(query)
		for _, option := range options {
			option(q)
		}
	}
	results, err := query.All(ctx)
	if err != nil {
		return err
	}

	for _, result := range results {
		r.ent.Area.DeleteOne(result).Exec(ctx)
	}

	return err
}

func NewAreaRepo(ent *ent.Client, rds *redis.Client) *AreaRepo {
	return &AreaRepo{
		ent: ent,
		rds: rds,
	}
}
