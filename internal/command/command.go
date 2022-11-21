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
	app.Commands = append(app.Commands, &cli.Command{
		Name:            "github",
		HideHelp:        true,
		HideHelpCommand: true,
		Hidden:          true,
		Subcommands: []*cli.Command{
			{
				Name:                   "load",
				Flags:                  []cli.Flag{ConfigFlag},
				UseShortOptionHandling: true,
				SkipFlagParsing:        true,
				Action: func(ctx *cli.Context) error {
					command, cleanup, err := fn(ctx.String("config"))
					if err != nil {
						return err
					}
					defer cleanup()
					return command.handler.Github.Load(ctx.Context)
				},
			},
		},
	})
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
