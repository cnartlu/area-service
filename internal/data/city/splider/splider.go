package asset

import (
	"context"
	"time"

	"github.com/cnartlu/area-service/errors"
	bizsplider "github.com/cnartlu/area-service/internal/biz/city/splider"
	"github.com/cnartlu/area-service/internal/data/data"
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/citysplider"
)

var _ bizsplider.SpliderRepo = (*SpliderRepo)(nil)

func convertSplider(model *ent.CitySplider) bizsplider.Splider {
	return bizsplider.Splider{
		ID:          model.ID,
		Source:      model.Source,
		Owner:       model.Owner,
		Repo:        model.Repo,
		SourceID:    model.SourceID,
		Title:       model.Title,
		Draft:       model.Draft,
		PreRelease:  model.PreRelease,
		PublishedAt: time.Unix(int64(model.PublisheTime), 0),
		Status:      bizsplider.Status(model.Status),
		CreatedAt:   time.Unix(int64(model.CreateTime), 0),
		UpdatedAt:   time.Unix(int64(model.UpdateTime), 0),
	}
}

type SpliderRepo struct {
	data *data.Data
}

func (r *SpliderRepo) Count(ctx context.Context, queries ...bizsplider.Query) int {
	client := r.data.GetClient(ctx)
	query := client.CitySplider.Query()
	o := NewQuery(query)
	for _, option := range queries {
		option := option
		option(o)
	}
	i, _ := query.Count(ctx)
	return i
}

func (r *SpliderRepo) FindList(ctx context.Context, queries ...bizsplider.Query) ([]*bizsplider.Splider, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySplider.Query()
	o := NewQuery(query)
	for _, option := range queries {
		option(o)
	}
	models, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	items := []*bizsplider.Splider{}
	for _, model := range models {
		model := model
		data := convertSplider(model)
		items = append(items, &data)
	}
	return items, nil
}

func (r *SpliderRepo) FindOne(ctx context.Context, options ...bizsplider.Query) (*bizsplider.Splider, error) {
	client := r.data.GetClient(ctx)
	query := client.CitySplider.Query()
	o := NewQuery(query)
	for _, option := range options {
		option(o)
	}
	model, err := query.First(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return nil, errors.ErrorDataNotFound(err.Error())
		}
		return nil, err
	}
	data := convertSplider(model)
	return &data, nil
}

func (r *SpliderRepo) Save(ctx context.Context, data *bizsplider.Splider) (*bizsplider.Splider, error) {
	client := r.data.GetClient(ctx)
	var (
		model    *ent.CitySplider
		err      error
		isUpdate bool
	)
	if data.ID > 0 {
		isUpdate = true
		model, err = client.CitySplider.Query().Where(citysplider.IDEQ(data.ID)).First(ctx)
		if err != nil {
			if !ent.IsNotFound(err) {
				return nil, err
			}
			isUpdate = false
		}
	}
	if data.PublishedAt.IsZero() {
		data.PublishedAt = time.Now()
	}
	if data.CreatedAt.IsZero() {
		data.CreatedAt = time.Now()
	}
	if data.UpdatedAt.IsZero() {
		data.UpdatedAt = time.Now()
	}
	if isUpdate {
		model, err = model.Update().
			SetSource(data.Source).
			SetOwner(data.Owner).
			SetRepo(data.Repo).
			SetSourceID(data.SourceID).
			SetTitle(data.Title).
			SetDraft(data.Draft).
			SetPreRelease(data.PreRelease).
			SetPublisheTime(uint64(data.PublishedAt.Unix())).
			SetStatus(uint8(data.Status)).
			SetUpdateTime(uint64(data.UpdatedAt.Unix())).
			Save(ctx)
	} else {
		model, err = client.CitySplider.Create().
			SetSource(data.Source).
			SetOwner(data.Owner).
			SetRepo(data.Repo).
			SetSourceID(data.SourceID).
			SetTitle(data.Title).
			SetDraft(data.Draft).
			SetPreRelease(data.PreRelease).
			SetPublisheTime(uint64(data.PublishedAt.Unix())).
			SetStatus(uint8(data.Status)).
			SetCreateTime(uint64(data.CreatedAt.Unix())).
			SetUpdateTime(uint64(data.UpdatedAt.Unix())).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	result := convertSplider(model)
	return &result, nil
}

func (r *SpliderRepo) Remove(ctx context.Context, options ...bizsplider.Query) error {
	client := r.data.GetClient(ctx)
	query := client.CitySplider.Query()
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
		if err1 := client.CitySplider.DeleteOne(result).Exec(ctx); err1 != nil {
			err = err1
			break
		}
	}
	return err
}

func NewSpliderRepo(
	d *data.Data,
) *SpliderRepo {
	return &SpliderRepo{
		data: d,
	}
}
