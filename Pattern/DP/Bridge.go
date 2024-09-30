package DP

import "fmt"

type Painter interface {
	Paint()
	//SetColor(c *Color)
}

type Color struct {
	Setcolor string
}

func (c *Color) Paint() {
	fmt.Printf("the color is %v\n", c.Setcolor)
}

type Shape struct {
	Shape_type string
	color      Color
}

func (s *Shape) Paint() {
	fmt.Printf("The figure is a %v\n", s.Shape_type)
	s.color.Paint()
}

func (s *Shape) SetColor(c *Color) {
	s.color = *c
}
