package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hello",
		Short: "a more complicated CLI app",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("running root command")
		},
	}
	helpCmd := &cobra.Command{
		Use:   "help",
		Short: "Help command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("usage: jmp save <dir> / jmp to <dir>")
		},
	}

	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "Save command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("usage: jmp save <dir>")
			fmt.Println("save dir as json in jmp.json")
		},
	}

	toCmd := &cobra.Command{
		Use:   "to",
		Short: "Jump to command",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("jmp to <dir> using bash script?")
		},
	}
	rootCmd.AddCommand(helpCmd, toCmd, saveCmd)
	rootCmd.Execute()
}
