package main

import (
	"math"
	"fmt"
)

type MyFloat float64

// You can't use (f float64) as the received, only types defined in the current package
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
