package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func NewProgressiveCommands() []*cli.Command {
	return []*cli.Command{
		{
			Name:  "progress",
			Usage: "prints Progress message",
			Action: func(c *cli.Context) error {
				fmt.Println("Progressing...")
				return nil
			},
		},
		{
			Name:  "status",
			Usage: "prints Status message",
			Action: func(c *cli.Context) error {
				fmt.Println("Status is good!")
				return nil
			},
		},
	}
}
