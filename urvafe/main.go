package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
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
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
