package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"main.go/MySQRT"
	"main.go/PrintCoditioned"
)

func TestPrinter(t *testing.T) {
	assert.Equal(t, float32(1.4142), MySQRT.MyRad2(2, 0.00001))
	assert.Equal(t, float32(2.0), MySQRT.MyRad2(4, 0.0001))
	assert.Equal(t, float32(1.7321), MySQRT.MyRad2(3, 0.0000001))

	var listToCheckValue PrintCoditioned.CheckValues
	listToCheckValue.ListOfChecks = make([]PrintCoditioned.Checker, 5)
	listToCheckValue.ListOfChecks[0] = &PrintCoditioned.PrimeCheckValue{Message: "PRIME NUMBER"}
	listToCheckValue.ListOfChecks[1] = &PrintCoditioned.Divisible{Message: "boo", CheckValue: 3}
	listToCheckValue.ListOfChecks[2] = &PrintCoditioned.Divisible{Message: "baa", CheckValue: 5}
	listToCheckValue.ListOfChecks[3] = &PrintCoditioned.Divisible{Message: "two", CheckValue: 2}
	listToCheckValue.ListOfChecks[4] = &PrintCoditioned.Divisible{Message: "APEX LEGEND", CheckValue: 7}

	assert.Equal(t, listToCheckValue.ExecuteAllCheckValues(14), "0,PRIME NUMBER,PRIME NUMBER,PRIME NUMBER,two,PRIME NUMBER,boo,PRIME NUMBER,two,boo,baa,PRIME NUMBER,boo,PRIME NUMBER,two,")

}
