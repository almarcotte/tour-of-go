package main

import (
	"time"
	"fmt"
)

func main() {
	t := time.Now()
	// useful to replace big if-then-else chains
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon!")
	default:
		fmt.Println("Good evening!")
	}
}
