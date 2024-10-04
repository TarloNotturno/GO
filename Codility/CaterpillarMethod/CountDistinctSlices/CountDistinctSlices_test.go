package CountDistinctSlices

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test function CaterpillarSlice that check the longest slice with the first element not repeated. First check is done on a slice equal to [2,2], it should result
// in a first call returning index 0 for begin and end and a second call returning 1 1. Second check done with [1,2] it should return 0,1
func TestCaterpillar(t *testing.T) {
	assert := assert.New(t)

	ExpectedResul := Caterpillar{begin: 0, end: 0, number: 2}
	fmt.Println(">>Test caterpillar func, find slice with not first element repeted. Check with []int{2, 2}")
	input := []int{2, 2}
	OccurrencyMap := MappingOccurrency(input)
	control := CaterpillarSlice(input, OccurrencyMap, 0, 2)
	assert.Equal(control, ExpectedResul, "Returned value must be as the one expected")

	OccurrencyMap[control.number] = OccurrencyMap[control.number][1:]
	control = CaterpillarSlice(input[1:], OccurrencyMap, 1, 2)
	ExpectedResul = Caterpillar{begin: 1, end: 1, number: 2}
	assert.Equal(control, ExpectedResul, "Returned value must be as the one expected")

	fmt.Println(">>Test caterpillar func, find slice with not first element repeted. Check with []int{1, 2}")
	input = []int{1, 2}
	OccurrencyMap = MappingOccurrency(input)
	control = CaterpillarSlice(input, OccurrencyMap, 0, 2)
	ExpectedResul = Caterpillar{begin: 0, end: 1, number: 1}
	assert.Equal(control, ExpectedResul, "Returned value must be as the one expected")

}

// Test function sumResults that calculate the combination between index begin and end
func TestSumResults(t *testing.T) {
	assert := assert.New(t)
	solution := sumResults(0, 1)
	assert.Equal(solution, 3, "Returned value must be as the one expected")

	solution = sumResults(0, 0)
	assert.Equal(solution, 1, "Returned value must be as the one expected")

	solution = sumResults(1, 1)
	assert.Equal(solution, 1, "Returned value must be as the one expected")
}

// test empty array
func TestEmpty(t *testing.T) {
	var input_empty []int
	fmt.Printf(">>Test Zero: check empty array Solution(5, %v ) = %v\n", input_empty, Solution(5, input_empty))
	solution := Solution(5, input_empty)
	assert.Equal(t, solution, 0, "Returned value must be as the one expected")
}

// test array with just two equal element
func TestOne(t *testing.T) {
	input := []int{1}
	solution := Solution(6, input)
	fmt.Printf(">>First test: check empty array Solution(5, %v ) = %v\n", input, solution)
	assert.Equal(t, solution, 1, "Returned value must be as the one expected")

}

// test array with just two different element
func TestDouble(t *testing.T) {
	assert := assert.New(t)
	input := []int{1, 2}
	solution := Solution(6, input)
	fmt.Printf(">>Second test: check empty array Solution(5, %v ) = %v\n", input, solution)
	assert.Equal(3, solution, "Returned value must be as the one expected")

	input = []int{2, 2}
	solution = Solution(2, input)
	fmt.Printf(">>Second test: check empty array Solution(5, %v ) = %v\n", input, solution)
	assert.Equal(2, solution, "Returned value must be as the one expected")

}

// Codility example tested
func TestBaseCase(t *testing.T) {
	input := []int{3, 4, 5, 5, 2}
	solution := Solution(6, input)
	fmt.Printf(">>Base test: check codility example array Solution(5, %v ) = %v\n", input, solution)
	assert.Equal(t, 9, solution, "Returned value must be as the one expected")
}

// launch a test case with a longer array
func TestFunctionalCase0(t *testing.T) {

	input := []int{5, 4, 7, 5, 7, 9, 2}
	solution := Solution(9, input)
	fmt.Printf(">>Check array Solution(9, %v ) = %v\n", input, solution)
	assert.Equal(t, 18, solution, "Returned value must be as the one expected")
}

func TestFunctionalCase1(t *testing.T) {
	// longer array with a different condition
	input := []int{5, 4, 7, 5, 7, 9, 2, 3, 8, 7, 12, 2, 22, 15, 18, 2, 8}
	solution := Solution(22, input)
	fmt.Printf(">>Check array Solution(22, %v ) = %v\n", input, solution)
	assert.Equal(t, 75, solution, "Returned value must be as the one expected")

}
