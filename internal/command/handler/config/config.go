package config

import "log"

type Handler interface {
}

type handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) *handler {
	return &handler{
		logger: logger,
	}
}
