package release

import (
	"context"

	"github.com/cnartlu/area-service/api"
	bizrelease "github.com/cnartlu/area-service/internal/biz/area/release"

	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
	"github.com/go-redis/redis/v8"
)

var _ bizrelease.ManageRepo = (*ReleaseRepo)(nil)

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

func (r *ReleaseRepo) FindList(ctx context.Context, options ...bizrelease.Option) ([]*bizrelease.Release, error) {
	query := r.ent.AreaRelease.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	models, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	items := make([]*bizrelease.Release, len(models))
	for i, model := range models {
		model := model
		items[i] = &bizrelease.Release{
			ID:                 model.ID,
			Owner:              model.Owner,
			Repo:               model.Repo,
			ReleaseID:          model.ReleaseID,
			ReleaseName:        model.ReleaseName,
			ReleaseNodeID:      model.ReleaseNodeID,
			ReleasePublishedAt: model.ReleasePublishedAt,
			ReleaseContent:     model.ReleaseContent,
			Status:             bizrelease.Status(model.Status),
			CreateTime:         model.CreateTime,
			UpdateTime:         model.UpdateTime,
		}
	}
	return items, nil
}

func (r *ReleaseRepo) FindOne(ctx context.Context, options ...bizrelease.Option) (*bizrelease.Release, error) {
	query := r.ent.AreaRelease.Query()
	o := newOption(query)
	for _, option := range options {
		option := option
		option(o)
	}
	model, err := query.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, api.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	var result = bizrelease.Release{
		ID:                 model.ID,
		Owner:              model.Owner,
		Repo:               model.Repo,
		ReleaseID:          model.ReleaseID,
		ReleaseName:        model.ReleaseName,
		ReleaseNodeID:      model.ReleaseNodeID,
		ReleasePublishedAt: model.ReleasePublishedAt,
		ReleaseContent:     model.ReleaseContent,
		Status:             bizrelease.Status(model.Status),
		CreateTime:         model.CreateTime,
		UpdateTime:         model.UpdateTime,
	}
	return &result, nil
}

func (r *ReleaseRepo) Save(ctx context.Context, data *bizrelease.Release) (*bizrelease.Release, error) {
	var (
		model    *ent.AreaRelease
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = r.ent.AreaRelease.Query().
			Where(arearelease.IDEQ(data.ID)).
			First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if isUpdate {
		model, err = model.Update().
			SetOwner(data.Owner).
			SetRepo(data.Repo).
			SetReleaseID(data.ReleaseID).
			SetReleaseName(data.ReleaseName).
			SetReleaseNodeID(data.ReleaseNodeID).
			SetReleasePublishedAt(data.ReleasePublishedAt).
			SetReleaseContent(data.ReleaseContent).
			SetStatus(uint8(data.Status)).
			Save(ctx)
	} else {
		model, err = r.ent.AreaRelease.Create().
			SetRepo(data.Repo).
			SetReleaseID(data.ReleaseID).
			SetReleaseName(data.ReleaseName).
			SetReleaseNodeID(data.ReleaseNodeID).
			SetReleasePublishedAt(data.ReleasePublishedAt).
			SetReleaseContent(data.ReleaseContent).
			SetStatus(uint8(data.Status)).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	return &bizrelease.Release{
		ID:                 model.ID,
		Owner:              model.Owner,
		Repo:               model.Repo,
		ReleaseID:          model.ReleaseID,
		ReleaseName:        model.ReleaseName,
		ReleaseNodeID:      model.ReleaseNodeID,
		ReleasePublishedAt: model.ReleasePublishedAt,
		ReleaseContent:     model.ReleaseContent,
		Status:             bizrelease.Status(model.Status),
		CreateTime:         model.CreateTime,
		UpdateTime:         model.UpdateTime,
	}, nil
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
