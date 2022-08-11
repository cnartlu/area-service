package area

import "github.com/cnartlu/area-service/pkg/component/log"

type Servicer interface {
	Importer
}

type Service struct {
	logger *log.Logger
}

func NewService(logger *log.Logger) *Service {
	return &Service{
		logger: logger,
	}
}
