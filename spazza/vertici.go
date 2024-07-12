package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	x, y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.x*v.x + v.y*v.y)
}

func (p1 Vertex) Distance(p2 Vertex) float64 {
	return math.Sqrt(((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func main() {
	p1 := Vertex{1, 2}
	p2 := Vertex{3, 4}
	fmt.Println(p1.Abs())
	fmt.Println(p1.Distance(p2))
}
