package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex {
		42.123, -21.234,
	}
	fmt.Println(m["Bell Labs"])
}
