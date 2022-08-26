package tests

import (
	"github.com/cnartlu/area-service/internal/config"
	// "github.com/cnartlu/area-service/pkg/component/log"
	"github.com/go-redis/redis/v8"
)

func init() {
	// testing.Init()
}

func NewConfig() (*config.Config, error) {
	return config.New("")
}

type Tests struct {
	// Logger *log.Logger
	Config *config.Config
	Redis  *redis.Client
}

func New(
	// logger *log.Logger,
	config *config.Config,
	rdb *redis.Client,
) *Tests {
	return &Tests{
		// Logger: logger,
		Config: config,
		Redis:  rdb,
	}
}
