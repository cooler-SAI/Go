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

	var helloCmd = &cobra.Command{
		Use:   "hello",
		Short: "hl",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("name")
			if name != "" {
				fmt.Printf("Hello, %s!\n", name)
			} else {
				fmt.Println("Hello All")
			}
		},
	}

	var brotherCmd = &cobra.Command{
		Use:   "brother loh",
		Short: "hl",
		Run: func(cmd *cobra.Command, args []string) {
			name, _ := cmd.Flags().GetString("brother loh")
			if name != "" {
				fmt.Printf("Hello, %s!\n", name)
			} else {
				fmt.Println("Hello Brother, you are LOH!!!")
			}
		},
	}

	// Добавляем альтернативное имя команды "hl" для helloCmd
	helloCmd.Aliases = []string{"hl"}

	helloCmd.Flags().StringP("name", "n", "", "Name to greet")

	rootCmd.AddCommand(echoCmd)
	rootCmd.AddCommand(helloCmd)
	rootCmd.AddCommand(brotherCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		return
	}
}
