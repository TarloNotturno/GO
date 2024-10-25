package main

import "fmt"

func BinaryGap(N int) int {
	maxBinaryGap := 0
	internal_count := 0
	binaryGapStarted := false
	for N > 0 {
		if N%2 == 0 {
			internal_count++
			N = N / 2
		} else {
			if internal_count > maxBinaryGap && binaryGapStarted {
				maxBinaryGap = internal_count
			}
			if !binaryGapStarted {
				binaryGapStarted = true
			}
			internal_count = 0
			N = N / 2

		}
	}
	return maxBinaryGap
}

func main() {
	fmt.Println(BinaryGap(32))
}
