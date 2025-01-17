package PrintCoditioned

import "fmt"

type Checker interface {
	printer(input int) bool
}

type CheckValues struct {
	ListOfChecks []Checker
}

type Divisible struct {
	Message    string
	CheckValue int
}

func (v *Divisible) printer(input int) bool {
	if input%v.CheckValue == 0 && input != 0 {
		fmt.Printf("%v", v.Message)
		return true
	}
	return false
}

type PrimeCheckValue struct {
	Message string
}

func (p *PrimeCheckValue) printer(input int) bool {
	primeNumber := true
	if input%2 == 0 && input != 2 {
		primeNumber = false
	}
	for i := 3; i < input; i = i + 2 {
		if input%i == 0 && input != i {
			primeNumber = false
		}
	}
	if primeNumber {
		fmt.Printf("%v", p.Message)
		return true
	}
	return false
}

func (c *CheckValues) ExecuteAllCheckValues(maxNCheckValues int) {
	for i := 0; i <= maxNCheckValues; i++ {
		printed := false
		for _, currentCheckValue := range c.ListOfChecks {
			if !printed {

				printed = printed || currentCheckValue.printer(i)
			}
		}
		if !printed {
			fmt.Printf("%v,", i)
		} else {
			fmt.Printf(",")
		}
	}
	fmt.Println()

}
