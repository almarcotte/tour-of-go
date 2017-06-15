package main

import "fmt"

func main() {
	// defer pushes the execution of the function until the surrounding function returns
	defer fmt.Println("world")
	fmt.Println("Hello")
}
