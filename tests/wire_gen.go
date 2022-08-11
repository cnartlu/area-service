// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package tests

import (
	component2 "github.com/cnartlu/area-service/internal/component"
	"github.com/cnartlu/area-service/internal/component/db"
	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component"
	"github.com/cnartlu/area-service/pkg/component/log"
	"github.com/cnartlu/area-service/pkg/component/redis"
	"github.com/google/wire"
)

// Injectors from wire.go:

func Init() (*Tests, func(), error) {
	logger, err := log.NewDefault()
	if err != nil {
		return nil, nil, err
	}
	config, err := NewConfig()
	if err != nil {
		return nil, nil, err
	}
	bootstrap := config.Bootstrap
	dbConfig := bootstrap.Database
	client, cleanup, err := db.NewEnt(dbConfig, logger)
	if err != nil {
		return nil, nil, err
	}
	redisConfig := bootstrap.Redis
	redisClient, cleanup2, err := redis.New(redisConfig, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	tests := New(logger, config, client, redisClient)
	return tests, func() {
		cleanup2()
		cleanup()
	}, nil
}

// wire.go:

var providerSet = wire.NewSet(
	NewConfig, log.NewDefault, config.ProviderSet, component.ProviderSet, component2.ProviderSet,
)
