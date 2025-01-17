package main

import (
	"fmt"

	"main.go/MySQRT"
	"main.go/PrintCoditioned"
)

func main() {
	var listToCheckValue PrintCoditioned.CheckValues
	listToCheckValue.ListOfChecks = make([]PrintCoditioned.Checker, 5)
	listToCheckValue.ListOfChecks[0] = &PrintCoditioned.PrimeCheckValue{Message: "PRIME NUMBER"}
	listToCheckValue.ListOfChecks[1] = &PrintCoditioned.Divisible{Message: "boo", CheckValue: 3}
	listToCheckValue.ListOfChecks[2] = &PrintCoditioned.Divisible{Message: "baa", CheckValue: 5}
	listToCheckValue.ListOfChecks[3] = &PrintCoditioned.Divisible{Message: "two", CheckValue: 2}
	listToCheckValue.ListOfChecks[4] = &PrintCoditioned.Divisible{Message: "APEX LEGEND", CheckValue: 7}
	listToCheckValue.ExecuteAllCheckValues(100)
	fmt.Println(MySQRT.MyRad2(100, 0.0001))
	fmt.Println(MySQRT.MyRad2(2, 0.0001))
	fmt.Println(MySQRT.MyRad2(25, 0.00001))

}
