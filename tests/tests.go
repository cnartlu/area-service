package tests

import (
	"testing"

	"github.com/cnartlu/area-service/internal/config"
	"github.com/cnartlu/area-service/pkg/component/log"
)

func init() {
	testing.Init()
}

type Tests struct {
	Logger *log.Logger
	Config *config.Config
}

func New(
	logger *log.Logger,
	config *config.Config,
) *Tests {
	return &Tests{
		Logger: logger,
		Config: config,
	}
}
