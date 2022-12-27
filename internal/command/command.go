package command

import (
	"github.com/cnartlu/area-service/internal/command/handler"
	"github.com/cnartlu/area-service/internal/command/script"
	"github.com/urfave/cli/v2"
)

type CommandFunc func(string) (*Command, func(), error)

var (
	TestFlag = &cli.BoolFlag{
		Name:    "test",
		Aliases: []string{"t"},
		Usage:   "test configuration and exit",
	}
	ConfigFlag = &cli.StringFlag{
		Name:        "config",
		Aliases:     []string{"c"},
		Value:       "config.yaml",
		DefaultText: "config.yaml",
		Usage:       "set configuration file",
	}
	SignalFlag = &cli.StringFlag{
		Name:    "signal",
		Aliases: []string{"s"},
		Usage:   "send signal to a master process: stop, quit, reload",
		Value:   "",
	}
	HelpFlag    = cli.HelpFlag
	VersionFlag = cli.VersionFlag
)

func Setup(app *cli.App, fn CommandFunc) {
	var (
		defaultCommand  *Command
		defaultCleanup  = func() {}
		initCommandfunc = func(ctx *cli.Context) error {
			command, cleanup, err := fn(ctx.String("config"))
			if err != nil {
				return err
			}
			defaultCleanup = cleanup
			defaultCommand = command
			return nil
		}
		emptyCommandFunction = func(ctx *cli.Context) error {
			return nil
		}
	)
	var registerCommand = &cli.Command{
		Name:            "register",
		HideHelpCommand: true,
		Before:          initCommandfunc,
		Action:          emptyCommandFunction,
		After: func(ctx *cli.Context) error {
			defaultCleanup()
			return nil
		},
	}

	{
		github := *registerCommand
		github.Name = "github"
		github.Before = emptyCommandFunction
		github.Action = func(ctx *cli.Context) error {
			return nil
		}
		{
			// latest 命令
			latest := *registerCommand
			latest.Name = "latest"
			latest.Action = func(ctx *cli.Context) error {
				return defaultCommand.handler.Github.Latest(ctx.Context)
			}
			// load 命令
			load := *registerCommand
			load.Name = "load"
			load.Action = func(ctx *cli.Context) error {
				return defaultCommand.handler.Github.Load(ctx.Context)
			}

			// 注册命令
			github.Subcommands = []*cli.Command{
				&latest,
				&load,
			}
		}

		// 注册命令
		app.Commands = append(app.Commands, &github)
	}

}

type Command struct {
	handler *handler.Handler
	script  *script.Script
}

func New(
	handler *handler.Handler,
	script *script.Script,
) *Command {
	var r = &Command{
		handler: handler,
		script:  script,
	}
	return r
}
