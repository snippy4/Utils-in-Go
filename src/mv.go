package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		fmt.Println("correct usage: mv <filename> <dir>")
		return
	}
	file, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	file.Close()
	err = os.MkdirAll(filepath.Dir(args[2]), os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
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
	err = os.Remove(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
}
