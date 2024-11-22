package main

import(
	"fmt"
	"os"
)

func main(){
	dir, _ := os.Getwd()
	files, _ := os.ReadDir(dir)
	for _, file := range files{
		fmt.Println(file.Name())
}

	
}