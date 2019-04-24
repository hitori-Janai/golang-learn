package main

import (
	"os"
	"fmt"
)

var cwd string

func main() {
	var err error
	cwd, err = os.Getwd()
	// cwd, err := os.Getwd()
	if err != nil {
		fmt.Printf("os.Getwd failed: %v\n", err)
	}
	fmt.Println(cwd)

	printcwd()
}

func printcwd(){
	fmt.Print(cwd)

}
