package CountDistinctSlices

import (
	"fmt"
	"testing"
)

// Test function CaterpillarSlice that check the longest slice with the first element not repeated. First check is done on a slice equal to [2,2], it should result
// in a first call returning index 0 for begin and end and a second call returning 1 1. Second check done with [1,2] it should return 0,1
func TestCaterpillar(t *testing.T) {
	fmt.Println("Test caterpillar func, find slice with not first element repeted. Check with []int{2, 2}")
	input := []int{2, 2}
	OccurrencyMap := MappingOccurrency(input)
	control := CaterpillarSlice(input, OccurrencyMap, 0)
	if control.begin != 0 || control.end != 0 {
		t.Fatalf(`CaterpillarSlice([]int{2, 2},0) obtain begin = %v expected equal to 0 and obtain end = %v expected equal to 0`, control.begin, control.end)
	}
	control = CaterpillarSlice(input[1:], OccurrencyMap, 1)
	if control.begin != 1 || control.end != 1 {
		t.Fatalf(`CaterpillarSlice([]int{2},0) obtain begin = %v expected equal to 1 and obtain end = %v expected equal to 1`, control.begin, control.end)
	}

	fmt.Println("Test caterpillar func, find slice with not first element repeted. Check with []int{1, 2}")
	input = []int{1, 2}
	OccurrencyMap = MappingOccurrency(input)
	control = CaterpillarSlice(input, OccurrencyMap, 0)
	if control.begin != 0 || control.end != 1 {
		t.Fatalf(`CaterpillarSlice(%v)) obtain begin = %v expected equal to 0 and obtain end = %v expected equal to 1`, input, control.begin, control.end)
	}

}

// Test function sumResults that calculate the combination between index begin and end
func TestSumResults(t *testing.T) {

	if sumResults(0, 1) != 3 {
		t.Fatalf(`sumResults(0, 1) %v expect solution is 2`, sumResults(0, 1))
	}

	if sumResults(0, 0) != 1 {
		t.Fatalf(`sumResults(0, 0) %v expect solution is 1`, sumResults(0, 0))
	}

	if sumResults(1, 1) != 1 {
		t.Fatalf(`sumResults(1, 1) %v expect solution is 1`, sumResults(1, 1))
	}
}

// test empty array
func TestEmpty(t *testing.T) {
	var input_empty []int
	fmt.Printf("Test Zero: check empty array Solution(5, %v ) = %v\n", input_empty, Solution(5, input_empty))
	if Solution(5, input_empty) != 0 {
		t.Fatalf(`Solution(5, []int) %v expect solution is "", 0`, Solution(5, input_empty))
	}
}

// test array with just two equal element
func TestOne(t *testing.T) {
	input := []int{1}
	solution := Solution(6, input)
	fmt.Printf("First test: check empty array Solution(5, %v ) = %v\n", input, solution)
	if solution != 1 {
		t.Fatalf(`Solution(5, %v) %v expect solution is 1`, input, solution)
	}

}

// test array with just two different element
func TestDouble(t *testing.T) {
	input := []int{1, 2}
	solution := Solution(6, input)
	fmt.Printf("Second test: check empty array Solution(5, %v ) = %v\n", input, solution)
	if solution != 3 {
		t.Fatalf(`Solution(5, []int{1, 2}) %v expect solution is 3`, solution)
	}
	input = []int{2, 2}
	solution = Solution(2, input)
	fmt.Printf("Second test: check empty array Solution(5, %v ) = %v\n", input, solution)
	if solution != 2 {
		t.Fatalf(`Solution(5, %v) %v expect solution is 2`, input, solution)
	}

}

// Codility example tested
func TestBaseCase(t *testing.T) {
	input := []int{3, 4, 5, 5, 2}
	solution := Solution(6, input)
	fmt.Printf("Base test: check codility example array Solution(5, %v ) = %v\n", input, solution)
	if solution != 9 {
		t.Fatalf(`Solution(5,  %v) %v expect solution is 9`, input, solution)
	}

}

// launch a test case with a longer array
func TestFunctionalCase0(t *testing.T) {

	input := []int{5, 4, 7, 5, 7, 9, 2}
	solution := Solution(9, input)
	fmt.Printf("Check array Solution(9, %v ) = %v\n", input, solution)
	if solution != 18 {
		t.Fatalf(`Solution(5, %v) %v expect solution is 18`, input, solution)
	}
}

func TestFunctionalCase1(t *testing.T) {
	// longer array with a different condition
	input := []int{5, 4, 7, 5, 7, 9, 2, 3, 8, 7, 12, 2, 22, 15, 18, 2, 8}
	solution := Solution(22, input)
	fmt.Printf("Check array Solution(22, %v ) = %v\n", input, solution)
	fmt.Printf("Second test: check empty array Solution(5, %v ) = %v\n", input, solution)

	if Solution(6, input) != 51 {
		t.Fatalf(`Solution(5, %v) %v expect solution is 51`, input, solution)
	}
}
