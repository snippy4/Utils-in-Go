package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("correct usage: rmdir <folder name>")
		return
	}
	dir := args[1]
	files, _ := os.ReadDir(dir)
	if len(files) != 0 {
		fmt.Println("Folder not empty")
		return
	}
	os.Remove(dir)
}
