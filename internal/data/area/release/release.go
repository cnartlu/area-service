package release

import (
	"context"

	bizrelease "github.com/cnartlu/area-service/internal/biz/area/release"

	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/go-redis/redis/v8"
)

var _ bizrelease.ManageRepo = (*ReleaseRepo)(nil)

type FindListParam struct {
	Owner      string
	Repository string
	Keyword    string
	Status     string
	Pagination bool
	Offset     int
	Limit      int
}

type ReleaseRepo struct {
	ent *ent.Client
	rdb *redis.Client
}

func (r *ReleaseRepo) Count(ctx context.Context, options ...bizrelease.Option) int {
	query := r.ent.AreaRelease.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	i, _ := query.Count(ctx)
	return i
}

// FindList 查找数据列表
func (r *ReleaseRepo) FindList(ctx context.Context, options ...bizrelease.Option) ([]*bizrelease.Release, error) {
	query := r.ent.AreaRelease.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	_, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *ReleaseRepo) FindOne(ctx context.Context, options ...bizrelease.Option) (*bizrelease.Release, error) {
	query := r.ent.AreaRelease.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	_, err := query.First(ctx)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Save 保存数据
func (r *ReleaseRepo) Save(ctx context.Context, data *bizrelease.Release) (model *bizrelease.Release, err error) {
	// if data.ID == 0 {
	// 	model, err = r.ent.AreaRelease.Create().
	// 		SetOwner(data.Owner).
	// 		// SetRepo(data.Repository).
	// 		SetReleaseID(data.ReleaseID).
	// 		SetReleaseName(data.ReleaseName).
	// 		SetReleaseNodeID(data.ReleaseNodeID).
	// 		SetReleasePublishedAt(data.ReleasePublishedAt).
	// 		SetReleaseContent(data.ReleaseContent).
	// 		SetStatus(data.Status).
	// 		Save(ctx)
	// } else {
	// 	model, err = r.ent.AreaRelease.UpdateOne(data).
	// 		SetOwner(data.Owner).
	// 		// SetRepo(data.Repository).
	// 		SetReleaseID(data.ReleaseID).
	// 		SetReleaseName(data.ReleaseName).
	// 		SetReleaseNodeID(data.ReleaseNodeID).
	// 		SetReleasePublishedAt(data.ReleasePublishedAt).
	// 		SetReleaseContent(data.ReleaseContent).
	// 		SetStatus(data.Status).
	// 		Save(ctx)
	// }
	// if err != nil {
	// 	return nil, err
	// }
	return
}

func (r *ReleaseRepo) Remove(ctx context.Context, options ...bizrelease.Option) error {
	query := r.ent.AreaRelease.Query()
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
		r.ent.AreaRelease.DeleteOne(result).Exec(ctx)
	}
	return err
}

func NewRepository(ent *ent.Client, rdb *redis.Client) *ReleaseRepo {
	return &ReleaseRepo{ent, rdb}
}
