// cmd/base.go
package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func RunApp() {
	app := &cli.App{
		Commands: append(
			[]*cli.Command{
				{
					Name:  "hello",
					Usage: "prints Hello urfave",
					Action: func(c *cli.Context) error {
						fmt.Println("Hello urfave")
						return nil
					},
				},
				{
					Name:  "goodbye",
					Usage: "prints Goodbye urfave",
					Action: func(c *cli.Context) error {
						fmt.Println("Goodbye urfave")
						return nil
					},
				},
				{
					Name:  "world",
					Usage: "prints World urfave",
					Action: func(c *cli.Context) error {
						fmt.Println("World")
						return nil
					},
				},
			},
			NewProgressiveCommands()...,
		),
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
