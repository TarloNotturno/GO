package main

import (
	"fmt"

	"main.go/Caterpillar"
)

func main() {
	var input []int
	//fmt.Println(CountTriangles.Solution(5, input))
	//input = []int{3, 4, 5, 5, 2}
	//fmt.Println("For ", input, " solution ", CountTriangles.Solution(5, input))
	//
	//input = []int{5, 4, 7, 5, 7, 9, 2}
	////fmt.Println("For ", input, " solution ", CountTriangles.Solution(5, input))

	//input = []int{5, 4, 7, 5, 7, 9, 2, 3, 8, 7, 12, 2, 22, 15, 18, 2, 8}
	//fmt.Println("For ", input, " solution ", CountTriangles.Solution(5, input))
	input = []int{2, 2}
	fmt.Println("For ", input, " solution ", Caterpillar.CountDistinctSlices(5, input))
	A := []int{10, 2, 5, 1, 8, 12}
	fmt.Println("For ", A, " solution ", Caterpillar.CountTriangles(A))
	//fmt.Println(Caterpillar.IsATriangle(A[0], A[2], A[5]))
	fmt.Println(Caterpillar.CaterpillarTriangle(A[2:], A[0]), A[1:])
}
