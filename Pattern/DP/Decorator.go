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



import (
	"fmt"
)

type Decorator interface {
	BasicPlot(inputS string) string
}

type Printer struct {
	start   string
	end     string
	toDecor Decorator
}

func (s Printer) BasicPlot(inputS string) string {
	if s.toDecor != nil {
		return s.start + s.toDecor.BasicPlot(inputS) + s.end
	} else {
		return s.start + inputS + s.end
	}
}

func main() {
	decor1 := Printer{start: "{", end: "}"}
	//decor2 := Printer{start: "[", end: "]", toDecor: decor1}
	//decor3 := Brachet{start: "(", end: ")"}
	fmt.Println(decor1.BasicPlot("pippo"))

}
