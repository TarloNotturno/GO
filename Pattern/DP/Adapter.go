package DP

import (
	"fmt"
	"math"
)

type Adapter interface {
	//Adapt(Adapter) float32
	GetDimension() float32
}

type RoundHole struct {
	Radius float32
}

type CirclePeg struct {
	Radius float32
}

type SquarePeg struct {
	Width float32
}

func (c CirclePeg) GetDimension() float32 {
	return c.Radius
}

func (s SquarePeg) GetDimension() float32 {
	return s.Width
}

func (r RoundHole) Adapt(input Adapter) float32 {
	switch input.(type) {
	case CirclePeg:
		return input.GetDimension()
	case SquarePeg:
		return ((float32)(math.Sqrt(2)) * input.GetDimension() / 2)
	default:
		return input.GetDimension()
	}
}

func (r RoundHole) Fit(input Adapter) bool {
	ToyDimension := r.Adapt((input))
	fmt.Printf(" - hole dimension :%v\n", r.Radius)
	fmt.Printf(" - toy dimension :%v\n", ToyDimension)
	return r.Radius > r.Adapt(input)
}
