package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("correct usage: tail <file name>")
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
	_, err = file1.Seek(0, 0)
	if err != nil {
		fmt.Println("Error resetting file pointer:", err)
		return
	}
	scanner1 := bufio.NewScanner(file1)
	count := 0
	for scanner1.Scan() {
		if linecount-count <= 10 {
			fmt.Println(scanner1.Text())
		}
		count++
	}
}
