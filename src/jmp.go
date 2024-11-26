package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

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
			fmt.Println("usage: jmp save <name> / jmp to <saved name>")
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

			// open jmp.json

			name := args[0]
			executable, err := os.Executable()
			if err != nil {
				fmt.Println(err)
				return
			}
			absoluteDir := filepath.Dir(executable)
			file, err := os.Open(absoluteDir + "\\jmp.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			defer file.Close()

			// load data from jmp.json

			decoder := json.NewDecoder(file)
			data := []dirName{}
			err = decoder.Decode(&data)
			if err != nil {
				fmt.Println(err)
				return
			}

			// add new data to jmp.json and save

			dir, _ := os.Getwd()
			entry := dirName{name, dir}
			data = append(data, entry)
			file.Close()
			file, err = os.Create(absoluteDir + "\\jmp.json")
			if err != nil {
				fmt.Println(err)
				return
			}
			encoder := json.NewEncoder(file)
			encoder.SetIndent("", "	 ")
			err = encoder.Encode(data)
			if err != nil {
				fmt.Println(err)
				return
			}
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
