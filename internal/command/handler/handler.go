package handler

import (
	"github.com/cnartlu/area-service/internal/command/handler/github"
)

type Handler struct {
	Github github.Handler
}

func New(
	github github.Handler,
) *Handler {
	return &Handler{
		Github: github,
	}
}
