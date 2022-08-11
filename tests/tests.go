package tests

import (
	"testing"

	"github.com/cnartlu/area-service/internal/component/ent"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/go-redis/redis/v8"
)

func init() {
	testing.Init()
}

func NewConfig() (*config.Config, error) {
	return config.New("")
}

type Tests struct {
	Logger *log.Logger
	Config *config.Config
	Ent    *ent.Client
	Redis  *redis.Client
}

func New(
	logger *log.Logger,
	config *config.Config,
	ent *ent.Client,
	rdb *redis.Client,
) *Tests {
	return &Tests{
		Logger: logger,
		Config: config,
		Ent:    ent,
		Redis:  rdb,
	}
}
