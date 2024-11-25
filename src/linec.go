package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("correct usage: linec <file name>")
		return
	}
	file1, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()
	scanner2 := bufio.NewScanner(file1)
	linecount := 0
	for scanner2.Scan() {
		linecount++
	}
	fmt.Println(linecount)
}
