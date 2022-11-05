package config

import (
	"log"

	"github.com/spf13/cobra"
)

type Handler interface {
	Save(cmd *cobra.Command, args []string)
	Set(cmd *cobra.Command, args []string)
	Get(cmd *cobra.Command, args []string)
}

type handler struct {
	logger *log.Logger
}

func (h *handler) Register(cmd *cobra.Command) {
	p := cobra.Command{
		Use:   "config",
		Short: "config",
		Long:  `config`,
		Run:   func(cmd *cobra.Command, args []string) {},
	}
	p.AddCommand(
		&cobra.Command{
			Use:   "save",
			Short: "save",
			Long:  `save`,
			Run: func(cmd *cobra.Command, args []string) {
				h.Save(cmd, args)
			},
		},
		&cobra.Command{
			Use:   "set",
			Short: "set",
			Long:  `set`,
			Run: func(cmd *cobra.Command, args []string) {
				h.Set(cmd, args)
			},
		},
		&cobra.Command{
			Use:   "get",
			Short: "get",
			Long:  `get`,
			Run: func(cmd *cobra.Command, args []string) {
				h.Get(cmd, args)
			},
		},
	)

	cmd.AddCommand(&p)
}

func (h *handler) Save(cmd *cobra.Command, args []string) {

}

func (h *handler) Get(cmd *cobra.Command, args []string) {

}

func (h *handler) Set(cmd *cobra.Command, args []string) {

}

func NewHandler(logger *log.Logger) *handler {
	return &handler{
		logger: logger,
	}
}
