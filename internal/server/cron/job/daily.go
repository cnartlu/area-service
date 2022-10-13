package job

import (
	"github.com/cnartlu/area-service/pkg/log"
)

type Daily struct {
	l *log.Logger
}

func (d *Daily) Run() {
	d.l.Debug("daily is run")
	d.l.Debug("daily is stop")
}

func NewDaily(logger *log.Logger) *Daily {
	return &Daily{
		l: logger,
	}
}
