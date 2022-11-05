package greet

import (
	"github.com/cnartlu/area-service/component/log"
	"github.com/spf13/cobra"
)

type Handler interface {
	Default(cmd *cobra.Command, args []string)
	To(cmd *cobra.Command, args []string)
}

type handler struct {
	logger *log.Logger
}

func NewHandler(logger *log.Logger) Handler {
	return &handler{
		logger: logger,
	}
}
