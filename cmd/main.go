package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/cnartlu/area-service/internal/command"
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
	var name string = Name
	if name == "" {
		s := filepath.Base(os.Args[0])
		if i := strings.Index(s, "."); i >= 0 {
			name = s[:i]
		} else {
			name = s
		}
	}
	var cmd = &cli.App{
		Name:                   name,
		Version:                Version,
		HideHelpCommand:        true,
		UseShortOptionHandling: true,
		Flags:                  []cli.Flag{command.TestFlag, command.SignalFlag, command.ConfigFlag},
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
			s, cleanup, err := initApp(context.TODO(), config)
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

	command.Setup(cmd, initCommand)

	if err := cmd.RunContext(ctx, os.Args); err != nil {
		fmt.Println(err)
	}
}
