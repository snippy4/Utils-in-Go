package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("correct usage: rm <file>")
		return
	}
	dir := args[1]
	err := os.Remove(dir)
	if err != nil {
		fmt.Println(err)
	}
}
