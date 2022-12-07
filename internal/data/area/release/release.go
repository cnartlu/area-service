package release

import (
	"context"

	"github.com/cnartlu/area-service/api"
	bizrelease "github.com/cnartlu/area-service/internal/biz/area/release"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
)

var _ bizrelease.ReleaseRepo = (*ReleaseRepo)(nil)

type ReleaseRepo struct {
	data *data.Data
}

func (r *ReleaseRepo) Count(ctx context.Context, options ...bizrelease.Query) int {
	client := r.data.GetClient(ctx)
	query := client.AreaRelease.Query()
	o := NewQuery(query)
	for _, option := range options {
		option := option
		option(o)
	}
	i, _ := query.Count(ctx)
	return i
}

func (r *ReleaseRepo) FindList(ctx context.Context, options ...bizrelease.Query) ([]*bizrelease.Release, error) {
	client := r.data.GetClient(ctx)
	query := client.AreaRelease.Query()
	o := NewQuery(query)
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

func (r *ReleaseRepo) FindOne(ctx context.Context, options ...bizrelease.Query) (*bizrelease.Release, error) {
	client := r.data.GetClient(ctx)
	query := client.AreaRelease.Query()
	o := NewQuery(query)
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
	client := r.data.GetClient(ctx)
	var (
		model    *ent.AreaRelease
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.AreaRelease.Query().
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
		model, err = client.AreaRelease.Create().
			SetOwner(data.Owner).
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

func (r *ReleaseRepo) Remove(ctx context.Context, options ...bizrelease.Query) error {
	client := r.data.GetClient(ctx)
	query := client.AreaRelease.Query()
	if len(options) > 0 {
		q := NewQuery(query)
		for _, option := range options {
			option(q)
		}
	}
	results, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, result := range results {
		if err1 := client.AreaRelease.DeleteOne(result).Exec(ctx); err1 != nil {
			err = err1
			break
		}
	}
	return err
}

func NewRepository(d *data.Data) *ReleaseRepo {
	return &ReleaseRepo{data: d}
}
