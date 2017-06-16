package main

import (
	"math"
	"fmt"
)

type Vertex struct {
	X, Y float64
}

// Receiver type of `Vertex`, method name: Abs, no arguments, returns a `float64`
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// No receiver, instead takes a `Vertex` as an argument
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs()) // dot-notation

	fmt.Println(Abs(v))
}
