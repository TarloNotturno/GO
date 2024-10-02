package main

import (
	"fmt"

	"main.go/CountDistinctSlices"
)

func main() {
	var input1 []int
	fmt.Println(CountDistinctSlices.Solution(5, input1))
	//fmt.Println(input1)
	input2 := []int{3, 4, 5, 5, 2}
	//fmt.Println(input2)
	fmt.Println("For ", input2, " solution ", CountDistinctSlices.Solution(5, input2))

	input3 := []int{5, 4, 7, 5, 7, 9, 2}
	fmt.Println("For ", input3, " solution ", CountDistinctSlices.Solution(5, input3))
	//A[0] = 3
	//A[1] = 4
	//A[2] = 5cd
	//A[3] = 5
	//A[4] = 2
}
