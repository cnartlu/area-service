package main

import (
	"context"
	"fmt"
	"os"
	"strings"

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
	var ctx = context.Background()
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
			var (
				help    = ctx.Bool("help")
				test    = ctx.Bool("test")
				version = ctx.Bool("version")
				signal  = strings.ToLower(strings.TrimSpace(ctx.String("signal")))
				config  = strings.TrimSpace(ctx.String("config"))
			)
			if test || help || version {
				return nil
			}
			s, cleanup, err := initApp(config)
			if err != nil {
				return err
			}
			defer cleanup()
			if err := s.Run(signal); err != nil {
				return err
			}
			return nil
		},
	}
	if err := cmd.RunContext(ctx, os.Args); err != nil {
		fmt.Println(err)
	}
}
