package polygon

import (
	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/go-redis/redis/v8"
)

type RepositoryInterface interface {
}

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
