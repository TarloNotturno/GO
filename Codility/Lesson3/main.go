package main

import (
	"fmt"
)

// FrogJmp excercise ---------------------------------------------------------------------
func FrogJmp(X int, Y int, D int) int {
	if Y < X {
		return 0
	}
	if D == 0 {
		return 0
	}
	distance := Y - X
	if distance%D == 0 {
		return (Y - X) / D
	} else {
		return (Y-X)/D + 1
	}
}

// PermMissingElem excercise ---------------------------------------------------------------------
func PermMissingElem(A []int) int {

	elementPresent := make(map[int]int)
	for _, value := range A {
		elementPresent[value] = 1
	}
	/* The element must go from 1 to N+1 (N = len(A)) */
	for i := 1; i <= len(A)+1; i++ {
		_, ok := elementPresent[i]
		if !ok {
			return i
		}
	}
	return len(A)
}

// TapeEquilibrium excercise ---------------------------------------------------------------------
func myAbs(value int) int {
	if value > 0 {
		return value
	} else {
		return -value
	}
}

func weightCalc(A []int) (w []int) {
	N := len(A)
	w = make([]int, N+1)
	w[0] = 0
	for i := 0; i < N; i++ {
		w[i+1] = w[i] + A[i]
	}
	return w
}

// TapeEquilibrium excercise ---------------------------------------------------------------------
func TapeEquilibrium(A []int) int {
	N := len(A)
	var minDiff int
	if N > 1 {
		minDiff = 9223372036854775807 //max int in go
	} else {
		if N == 0 {
			return 0
		}
		return A[0]
	}
	weight := weightCalc(A)
	for i := 1; i < N; i++ {
		//for index := i + 1; index < N; index++ {
		currentDiff := myAbs(2*weight[i] - weight[N])
		if currentDiff < minDiff {
			minDiff = currentDiff
		}

	}
	return minDiff
}

func main() {
	fmt.Println("For A := []int{2,4,5,1} PermMissingElem =", PermMissingElem([]int{2, 4, 5, 1}))
	fmt.Println("For A := []int{} PermMissingElem =", PermMissingElem([]int{}))
	fmt.Println("For A := []int{1} PermMissingElem =", PermMissingElem([]int{1}))
	fmt.Println("For A := []int{3, 1, 2, 4, 3} TapeEquilibrium =", TapeEquilibrium([]int{3, 1, 2, 4, 3}))
	fmt.Println("For A := []int{} TapeEquilibrium = ", TapeEquilibrium([]int{}))
	fmt.Println("For A := []int{5} TapeEquilibrium = ", TapeEquilibrium([]int{5}))
}
