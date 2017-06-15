package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	s = s[:0] // Slice to give it a length of 0
	printSlice(s)

	s = s[:4] // Extend its length
	printSlice(s)

	s = s[2:] // drop the first 2 values
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
