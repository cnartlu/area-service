package release

import (
	"context"

	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/cnartlu/area-service/internal/data/ent/arearelease"
	"github.com/go-redis/redis/v8"
)

type ReleaseRepository interface {
	// FindList 查找资源列表
	FindList(ctx context.Context, params *FindListParam) ([]*ent.AreaRelease, error)
	// FindOne 查找资源
	FindOne(ctx context.Context, id uint64) (*ent.AreaRelease, error)
	// Save 保存资源
	Save(ctx context.Context, data *ent.AreaRelease) (*ent.AreaRelease, error)
	// Remove 移除资源
	Remove(ctx context.Context, id uint64) error
}

type FindListParam struct {
	Owner      string
	Repository string
	Keyword    string
	Status     string
	Pagination bool
	Offset     int
	Limit      int
}

var _ ReleaseRepository = (*ReleaseRepo)(nil)

type ReleaseRepo struct {
	ent *ent.Client
	rdb *redis.Client
}

// FindList 查找数据列表
func (r *ReleaseRepo) FindList(ctx context.Context, params *FindListParam) ([]*ent.AreaRelease, error) {
	query := r.ent.AreaRelease.Query()
	if params != nil {
		if params.Keyword != "" {
			query.Where(arearelease.ReleaseNameContains(params.Keyword))
		}
		if params.Owner != "" {
			query.Where(arearelease.OwnerContains(params.Owner))
		}
		if params.Repository != "" {
			query.Where(arearelease.RepoContains(params.Repository))
		}
		if params.Pagination {
			query.Offset(params.Offset).Limit(params.Limit)
		}
	}
	query.Order(ent.Desc(arearelease.FieldID))
	return query.All(ctx)
}

func (r *ReleaseRepo) FindOne(ctx context.Context, id uint64) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Query().Where(arearelease.IDEQ(id)).First(ctx)
}

// FindByReleaseID 通过releaseID查找记录
func (r *ReleaseRepo) FindByReleaseID(ctx context.Context, releaseID uint64) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Query().
		Where(arearelease.ReleaseIDEQ(releaseID)).
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}

// FindOneWithLastRecord 查找最后一条记录
func (r *ReleaseRepo) FindOneWithLastRecord(ctx context.Context) (*ent.AreaRelease, error) {
	return r.ent.AreaRelease.Query().
		Order(ent.Desc(arearelease.FieldID)).
		First(ctx)
}

// Save 保存数据
func (r *ReleaseRepo) Save(ctx context.Context, data *ent.AreaRelease) (model *ent.AreaRelease, err error) {
	if data.ID == 0 {
		model, err = r.ent.AreaRelease.Create().
			SetOwner(data.Owner).
			// SetRepo(data.Repository).
			SetReleaseID(data.ReleaseID).
			SetReleaseName(data.ReleaseName).
			SetReleaseNodeID(data.ReleaseNodeID).
			SetReleasePublishedAt(data.ReleasePublishedAt).
			SetReleaseContent(data.ReleaseContent).
			SetStatus(data.Status).
			Save(ctx)
	} else {
		model, err = r.ent.AreaRelease.UpdateOne(data).
			SetOwner(data.Owner).
			// SetRepo(data.Repository).
			SetReleaseID(data.ReleaseID).
			SetReleaseName(data.ReleaseName).
			SetReleaseNodeID(data.ReleaseNodeID).
			SetReleasePublishedAt(data.ReleasePublishedAt).
			SetReleaseContent(data.ReleaseContent).
			SetStatus(data.Status).
			Save(ctx)
	}
	if err != nil {
		return nil, err
	}
	return
}

func (r *ReleaseRepo) Remove(ctx context.Context, id uint64) error {
	return r.ent.AreaRelease.DeleteOneID(id).Exec(ctx)
}

func NewRepository(ent *ent.Client, rdb *redis.Client) *ReleaseRepo {
	return &ReleaseRepo{ent, rdb}
}
