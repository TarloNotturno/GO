package DP

import "strconv"

//type Decorer interface {
//	Decorator(input string) string
//}

type ExampleDecorator struct {
	Begin string
	End   string
}

type NumericDecorator struct {
	Begin int
	End   int
}

func (e ExampleDecorator) Decorator(input string) string {
	return e.Begin + input + e.End
}

func (e NumericDecorator) Decorator(input string) string {
	return strconv.Itoa(e.Begin) + input + strconv.Itoa(e.End)
}
