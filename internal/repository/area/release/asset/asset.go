package asset

import (
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/go-redis/redis/v8"
)

type RepositoryManager interface {
	Querier
	Creator
	Updater
	Deleter
}

type Repository struct {
	ent *ent.Client
	rdb *redis.Client
}

var _ RepositoryManager = (*Repository)(nil)

func NewRepository(ent *ent.Client, rdb *redis.Client) *Repository {
	return &Repository{ent, rdb}
}
