package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	const Emoji = "🗽"
	fmt.Println("Hello", World)
	fmt.Println("Hello", Emoji)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}