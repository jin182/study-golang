package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["bell Labs"])
}