package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("correct usage: cp <filename> <dir>")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	newfile, err := os.Create(args[2])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer newfile.Close()
	_, err = io.Copy(newfile, file)
	err = newfile.Sync()
	if err != nil {
		fmt.Println(err)
		return
	}
}
