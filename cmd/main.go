package main

import (
	"context"
	"fmt"
	"os"

	"github.com/cnartlu/area-service/component/app"
	"github.com/urfave/cli/v2"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name string = ""
	// Version is the version of the compiled software.
	Version string = "0.0.1beta"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("recover", err)
		}
	}()
	var ctx = context.Background()
	var application *app.App
	var cmd = cli.App{
		Name:    Name,
		Version: Version,
		// Usage:                  fmt.Sprintf("%s [-hvt] [-s signal] [-c filename]", Name),
		HideHelpCommand:        true,
		UseShortOptionHandling: true,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "test",
				Aliases: []string{"t"},
				Usage:   "test configuration and exit",
			},
			&cli.StringFlag{
				Name:    "signal",
				Aliases: []string{"s"},
				Usage:   "send signal to a master process: stop, quit, reload",
				Value:   "",
			},
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "",
				DefaultText: "config.yaml",
				Usage:       "set configuration file",
			},
		},
		Action: func(ctx *cli.Context) error {
			application = app.New(
				app.WithFlagHelp(ctx.Bool("help")),
				app.WithFlagVersion(ctx.Bool("version")),
				app.WithFlagTest(ctx.Bool("test")),
				app.WithFlagSignal(ctx.String("signal")),
				app.WithFlagConfig(ctx.String("config")),
				app.WithStartFunc(func() error {
					appServ, appCleanup, err := initApp(application.Config())
					if err != nil {
						return err
					}
					defer appCleanup()
					if err := appServ.Start(); err != nil {
						return err
					}
					return nil
				}),
			)
			return application.Run()
		},
	}
	if err := cmd.RunContext(ctx, os.Args); err != nil {
		fmt.Println(err)
	}
}
