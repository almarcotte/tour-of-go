package main

import "fmt"

func main() {
	powers := []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range powers {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Skip the index or value by assigning it to _
	for _, value := range powers {
		fmt.Printf("%d\n", value)
	}
}
