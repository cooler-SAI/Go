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
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "urfave",
							Usage: "name to greet",
						},
					},
					Action: func(c *cli.Context) error {
						name := c.String("name")
						fmt.Printf("Hello %s\n", name)
						return nil
					},
				},
				{
					Name:  "goodbye",
					Usage: "prints Goodbye urfave",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "urfave",
							Usage: "name to bid goodbye",
						},
					},
					Action: func(c *cli.Context) error {
						name := c.String("name")
						fmt.Printf("Goodbye %s\n", name)
						return nil
					},
				},
				{
					Name:  "world",
					Usage: "prints World urfave",
					Flags: []cli.Flag{
						&cli.StringFlag{
							Name:  "name",
							Value: "World",
							Usage: "custom message",
						},
					},
					Action: func(c *cli.Context) error {
						name := c.String("name")
						fmt.Printf("%s urfave\n", name)
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
