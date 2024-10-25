package main

import "fmt"

func CyclicRotation(A []int, K int) []int {
	if K == 0 || len(A) == 0 {
		return A
	}
	if K%len(A) == 0 {
		return A
	}
	rotatedArray := make([]int, 0)
	for K > len(A) {
		K = K - len(A)
	}
	K = len(A) - K
	rotatedArray = append(rotatedArray, A[K:]...)
	rotatedArray = append(rotatedArray, A[:K]...)
	return rotatedArray
}

func OddOccurrencesInArray(A []int) int {
	occurrencyMap := make(map[int]int)
	for _, value := range A {
		_, ok := occurrencyMap[value]
		if ok {
			occurrencyMap[value]++
		} else {
			occurrencyMap[value] = 1
		}
	}
	for value, occ := range occurrencyMap {
		if occ%2 != 0 {
			return value
		}
	}
	return 0
}

func main() {
	fmt.Println(CyclicRotation([]int{1, 2, 3, 4}, 4))
	fmt.Println(CyclicRotation([]int{1, 2, 3, 4, 5}, 21))
	fmt.Println(CyclicRotation([]int{3, 8, 9, 7, 6}, 3))
	fmt.Println(CyclicRotation([]int{3, 8, 9, 7, 6}, 2))
	fmt.Println(CyclicRotation([]int{3, 8, 9, 7, 6}, 1))
	fmt.Println(OddOccurrencesInArray([]int{9, 3, 9, 3, 9, 7, 9}))
	fmt.Println(OddOccurrencesInArray([]int{}))
}
