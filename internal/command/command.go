package command

import (
	"github.com/cnartlu/area-service/internal/command/handler/greet"
	"github.com/cnartlu/area-service/internal/command/pkg/commandset"
	"github.com/cnartlu/area-service/internal/command/script"
	"github.com/spf13/cobra"
)

func Setup(rootCommand *cobra.Command, newCommand func() (*Command, func(), error)) {
	set := commandset.NewCommandSet(rootCommand)

	// TODO 编写子命令

	// 注册业务的子命令
	set.RegisterBusiness([]*commandset.Command{
		{
			Entity: &cobra.Command{
				Use:   "greet",
				Short: "示例子命令",
				Run: func(cmd *cobra.Command, args []string) {
					command, cleanup, err := newCommand()
					if err != nil {
						panic(err)
					}
					defer cleanup()
					command.greetHandler.Default(cmd, args)
				},
			},
			Option: func(command *cobra.Command) {
				command.Flags().StringP("example", "e", "foo", "示例 flag")
			},
			Children: []*commandset.Command{
				{
					Entity: &cobra.Command{
						Use:   "to",
						Short: "示例子命令",
						Run: func(cmd *cobra.Command, args []string) {
							command, cleanup, err := newCommand()
							if err != nil {
								panic(err)
							}
							defer cleanup()
							command.greetHandler.To(cmd, args)
						},
					},
				},
			},
		},
	})

	// 注册临时脚本命令
	set.RegisterScript([]*commandset.Command{
		{
			Entity: &cobra.Command{
				Use:   "S0000000000",
				Short: "示例脚本 S0000000000",
				Run: func(cmd *cobra.Command, args []string) {
					command, cleanup, err := newCommand()
					if err != nil {
						panic(err)
					}
					defer cleanup()
					command.s0000000000Script.Run(cmd, args)
				},
			},
		},
	})
}

type Command struct {
	greetHandler      greet.Handler
	s0000000000Script *script.S0000000000
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
