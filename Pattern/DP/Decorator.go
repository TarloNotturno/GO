package DP

import (
	"strconv"
)

type Decorer interface {
	Decorator(inputS string) string
}

type Printer struct {
	start   string
	end     string
	toDecor Decorer
}

type ExampleDecorator struct {
	start   string
	end     string
	toDecor Decorer
}

type NumericDecorator struct {
	start   int
	end     int
	toDecor Decorer
}

func (s ExampleDecorator) Decorator(input string) string {
	if s.toDecor != nil {
		return s.start + s.toDecor.Decorator(input) + s.end
	} else {
		return s.start + input + s.end
	}
}

func (s *ExampleDecorator) Initialize(start string, end string, d Decorer) {
	//switch start.(type) {
	//case string:
	s.start = start
	//default:
	//}
	//switch end.(type) {
	//case string:
	s.end = end
	//default:
	//}
	s.toDecor = d
}

func (s NumericDecorator) Decorator(input string) string {
	if s.toDecor != nil {
		return strconv.Itoa(s.start) + s.toDecor.Decorator(input) + strconv.Itoa(s.end)
	} else {
		return strconv.Itoa(s.start) + input + strconv.Itoa(s.end)
	}
}

func (s *NumericDecorator) Initialize(start int, end int, d Decorer) {
	//switch start.(type) {
	//case int:
	s.start = start
	//default:
	//}
	//switch end.(type) {
	//case int:
	s.end = end
	//default:
	//}
	s.toDecor = d
}

func (s Printer) Decorator(inputS string) string {
	if s.toDecor != nil {
		return s.start + s.toDecor.Decorator(inputS) + s.end
	} else {
		return s.start + inputS + s.end
	}
}

func (s *Printer) Initialize(start string, end string, d Decorer) {
	//switch start.(type) {
	//case string:
	s.start = start
	//default:
	//}
	//switch end.(type) {
	//case string:
	s.end = end
	//default:
	//}
	s.toDecor = d
}
