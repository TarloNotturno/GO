package main

import (
	"fmt"

	"main.go/CountDistinctSlices"
)

func main() {
	var input []int
	//fmt.Println(CountDistinctSlices.Solution(5, input))
	input = []int{3, 4, 5, 5, 2}
	fmt.Println("For ", input, " solution ", CountDistinctSlices.Solution(5, input))

	input = []int{5, 4, 7, 5, 7, 9, 2}
	//fmt.Println("For ", input, " solution ", CountDistinctSlices.Solution(5, input))

	input = []int{5, 4, 7, 5, 7, 9, 2, 3, 8, 7, 12, 2, 22, 15, 18, 2, 8}
	fmt.Println("For ", input, " solution ", CountDistinctSlices.Solution(5, input))
}
