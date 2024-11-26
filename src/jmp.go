package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type dirName struct {
	Name string `json:"name"`
	Dir  string `json:"dir"`
}

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
			fmt.Println("usage: jmp save <name> / jmp to <dir>")
		},
	}

	saveCmd := &cobra.Command{
		Use:   "save",
		Short: "Save command",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) != 1 {
				fmt.Println("usage: jmp save <name>")
				return
			}
			name := args[0]
			file, err := os.Open("jmp.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			decoder := json.NewDecoder(file)
			data := []dirName{}
			err = decoder.Decode(&data)
			if err != nil {
				fmt.Println(err)
				return
			}
			dir, _ := os.Getwd()
			entry := map[string]interface{}{"name": name, "dir": dir}
			data = append(dirName, entry)
			fmt.Println("data is", data)
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
