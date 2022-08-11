package release

import (
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/go-redis/redis/v8"
)

type RepositoryManager interface {
	Querier
	Creator
	Updater
}

var _ RepositoryManager = (*Repository)(nil)

type Repository struct {
	ent *ent.Client
	rdb *redis.Client
}

func NewRepository(ent *ent.Client, rdb *redis.Client) *Repository {
	return &Repository{ent, rdb}
}
