package main

import (
	"fmt"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "learning",
		Short: "learning is an example of a cobra command short",
		Long:  `learning is an example of a cobra command long`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Cobra!")
		},
	}
	var echoCmd = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo a string",
		Long:  "Echo is a subcommand that prints the input string to the terminal.",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args[0])
		},
	}
	rootCmd.AddCommand(echoCmd)
	rootCmd.Execute()

}
