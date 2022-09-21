package area

import (
	"context"
	"time"

	bizArea "github.com/cnartlu/area-service/internal/biz/area"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/go-redis/redis/v8"
)

type AreaRepo struct {
	ent *ent.Client
	rds *redis.Client
}

var _ bizArea.Manager = (*AreaRepo)(nil)

func (r *AreaRepo) toAreaData(result *ent.Area) bizArea.Area {
	return bizArea.Area{
		ID:        result.ID,
		RegionID:  result.RegionID,
		Title:     result.Title,
		Pinyin:    result.Pinyin,
		Ucfirst:   result.Ucfirst,
		CityCode:  result.CityCode,
		ZipCode:   result.ZipCode,
		Level:     int(result.Level),
		CreateAt:  time.Unix(int64(result.CreateTime), 0),
		UpddateAt: time.Unix(int64(result.UpdateTime), 0),
	}
}

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
		d := r.toAreaData(result)
		list = append(list, &d)
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
	data := r.toAreaData(result)
	return &data, err
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
	data := r.toAreaData(result)
	return &data, err
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
