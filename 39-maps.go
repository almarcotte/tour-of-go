package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

func main() {
	// Keys are type `string`, values are `Vertex`
	m := make(map[string]Vertex)

	// Assign a value to a key
	m["Bell Labs"] = Vertex{40.68433, -74.39967}

	fmt.Println(m["Bell Labs"])
}
