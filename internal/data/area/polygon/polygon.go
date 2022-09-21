package polygon

import (
	"github.com/cnartlu/area-service/internal/data/ent"
	"github.com/go-redis/redis/v8"
)

type RepositoryManager interface {
}

var _ RepositoryManager = (*RepositoryManager)(nil)

type Repository struct {
	ent *ent.Client
	rds *redis.Client
}

func NewRepository(ent *ent.Client, rds *redis.Client) *Repository {
	return &Repository{
		ent: ent,
		rds: rds,
	}
}
