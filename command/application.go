package command

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func App() *cli.App {
	return &cli.App{
		Name:  "cli",
		Usage: "fight the loneliness!",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Value:       "config/config.yaml",
				DefaultText: "config/config.yaml",
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
		Commands: []*cli.Command{},
	}
}
