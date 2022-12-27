package job

import "go.uber.org/zap"

type Daily struct {
	l *zap.Logger
}

func (d *Daily) Run() {
}

func NewDaily(logger *zap.Logger) *Daily {
	return &Daily{
		l: logger,
	}
}
