package job

import (
	"github.com/cnartlu/area-service/component/log"
)

type Daily struct {
	l *log.Logger
}

func (d *Daily) Run() {
}

func NewDaily(logger *log.Logger) *Daily {
	return &Daily{
		l: logger,
	}
}
