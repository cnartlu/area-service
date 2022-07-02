package casbin

import (
	"github.com/cnartlu/area-service/pkg/component/casbin/adapter"
	"github.com/cnartlu/area-service/pkg/component/casbin/model"

	"github.com/casbin/casbin/v2"
	"gorm.io/gorm"
)

type Config struct {
	Model   *model.Config
	Adapter *adapter.Config
}

func New(config *Config, db *gorm.DB) (*casbin.Enforcer, error) {
	if config == nil {
		return nil, nil
	}

	mod, err := model.New(config.Model)
	if err != nil {
		return nil, err
	}

	adp, err := adapter.New(config.Adapter, db)
	if err != nil {
		return nil, err
	}

	ef, err := casbin.NewEnforcer(mod, adp)
	if err != nil {
		return nil, err
	}

	return ef, nil
}
