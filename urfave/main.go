package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	fmt.Println("testing urfave...")

	app := &cli.App{
		Name:  "urfave",
		Usage: "testing urfave",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Value:   "config.json",
				Usage:   "config file",
			},
		},
		Action: func(c *cli.Context) error {
			configFile := c.String("config")
			if configFile == "" {
				configFile = "config.json"
			}
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "greet",
				Aliases: []string{"g"},
				Usage:   "Greet someone",
				Action: func(c *cli.Context) error {
					name := c.Args().Get(0)
					if name == "" {
						name = "World"
					}
					fmt.Printf("Hello, %s!\n", name)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		_, err := fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		if err != nil {
			return
		}
		os.Exit(1)

	}
}
