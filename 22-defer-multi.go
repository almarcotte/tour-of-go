package main

import "fmt"

func main() {
	fmt.Println("counting")
	// defer pushes the execution on a stack
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	// 9 8 7 6 5 4 3 2 1
	fmt.Println("done")
}
