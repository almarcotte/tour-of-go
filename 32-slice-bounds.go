package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4] // [3 5 7]
	fmt.Println(s)

	s = s[:2] // same as s[0:2]
	fmt.Println(s)

	s = s[1:] // same as [1:UNTIL THE END]
	fmt.Println(s)
}
