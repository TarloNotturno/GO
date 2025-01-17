package PrintCoditioned

import "strconv"

type Checker interface {
	printer(input int) (string, bool)
}

type CheckValues struct {
	ListOfChecks []Checker
}

type Divisible struct {
	Message    string
	CheckValue int
}

func (v *Divisible) printer(input int) (string, bool) {
	output := ""
	if input%v.CheckValue == 0 && input != 0 {
		output = v.Message
		return output, true
	}
	return output, false
}

type PrimeCheckValue struct {
	Message string
}

func (p *PrimeCheckValue) printer(input int) (string, bool) {
	primeNumber := true
	output := ""
	if input%2 == 0 && input != 2 {
		primeNumber = false
	}
	for i := 3; i < input; i = i + 2 {
		if input%i == 0 && input != i {
			primeNumber = false
		}
	}
	if primeNumber {
		output = p.Message
		return output, true
	}
	return output, false
}

func (c *CheckValues) ExecuteAllCheckValues(maxNCheckValues int) (output string) {
	output = ""
	var app1 string
	var app2 bool
	for i := 0; i <= maxNCheckValues; i++ {
		printed := false
		for _, currentCheckValue := range c.ListOfChecks {
			if !printed {
				app1, app2 = currentCheckValue.printer(i)
				printed = printed || app2
			}
		}
		if !printed {
			output = output + strconv.Itoa(i) + ","
		} else {
			output = output + app1 + ","
		}
	}
	return
}
