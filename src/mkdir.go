package main

import(
	"fmt"
	"os"
)

func main(){
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Correct usage: mkdir <folder name>")
		return
	}
	os.Mkdir(args[1], 0755)
}