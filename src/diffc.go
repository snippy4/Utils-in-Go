package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	args := os.Args
	if len(args) != 3{
		fmt.Println("correct usage: diff <file name> <file name>")
		return
	}
	file1, err := os.Open(args[1])
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file1.Close()
	file2, err := os.Open(args[2])
	if err != nil{
		fmt.Println(err)
		return
	}
	defer file2.Close()
	scanner1 := bufio.NewScanner(file1)
	scanner2 := bufio.NewScanner(file2)
	linecount := 0
	for scanner1.Scan() {
		scanner2.Scan()
		if scanner1.Text() != scanner2.Text(){
			linecount++
		}
    }
	for scanner2.Scan() {
		linecount++
	}
	fmt.Println(linecount, "lines differ")
}