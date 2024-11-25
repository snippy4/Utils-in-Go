package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("correct usage: head <file name>")
		return
	}
	file1, err := os.Open(args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file1.Close()
	scanner1 := bufio.NewScanner(file1)
	count := 0
	for scanner1.Scan() && count < 10 {
		fmt.Println(scanner1.Text())
		count++
	}
}
