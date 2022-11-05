package command

import (
	"github.com/cnartlu/area-service/internal/command/handler/greet"
	"github.com/cnartlu/area-service/internal/command/script"
	"github.com/spf13/cobra"
)

type Commander interface {
	Register(cmd *cobra.Command)
}

type Command struct {
	handlers Commander
	scripts  []script.Script
	// 注册脚本
	greetHandler      greet.Handler
	s0000000000Script *script.S0000000000
}

func (c *Command) Register(cmd *cobra.Command) {
	// c.greetHandler.Register(cmd)
	// 注册脚本命令
	cmd.AddCommand(&cobra.Command{
		Use:   "script",
		Short: "",
		Long:  ``,
		Run: func(cmd *cobra.Command, args []string) {

		},
	})
}

func New(
	greetHandler greet.Handler,
	s0000000000Script *script.S0000000000,
) *Command {
	return &Command{
		greetHandler:      greetHandler,
		s0000000000Script: s0000000000Script,
	}
}
