package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var brotherCmd = &cobra.Command{
	Use:   "brother coolMan",
	Short: "br",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("brother coolMan")
		if name != "" {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Println("Hello Brother, you are COOL MAN!!!")
		}
	},
}

var talkingCmd = &cobra.Command{
	Use:   "talking",
	Short: "A command for greeting in a cool way",
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("talking")
		if name != "" {
			fmt.Printf("Hello, %s!\n", name)
		} else {
			fmt.Println("Talking.....")
		}
	},
}

func init() {
	brotherCmd.Aliases = []string{"br"}

	brotherCmd.Flags().StringP("name", "n", "", "Name to greet")
	rootCmd.AddCommand(brotherCmd)
	talkingCmd.Flags().StringP("name", "n", "", "Name to greet")
	rootCmd.AddCommand(talkingCmd)
}
