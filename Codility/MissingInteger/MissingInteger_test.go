package MissingInteger

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// example test
func TestMissingInteger_t0(t *testing.T) {
	result := MissingInteger([]int{1, 3, 6, 4, 1, 2})
	expected := 5
	assert.Equal(t, result, expected, fmt.Sprintf("For array  [1,2,3] you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{1, 3, 6, 4, 1, 2}, " the first missing positive integer is: ", MissingInteger([]int{1, 3, 6, 4, 1, 2}))
}

// array with all value consecutive and positive
func TestMissingInteger_t1(t *testing.T) {
	result := MissingInteger([]int{1, 2, 3})
	expected := 4
	assert.Equal(t, result, expected, fmt.Sprintf("For array  [1,2,3] you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{1, 2, 3}, " the first missing positive integer is: ", MissingInteger([]int{1, 2, 3}))
}

// array with random value positive and negative
func TestMissingInteger_t2(t *testing.T) {
	result := MissingInteger([]int{1, -1000000, 7, 5, 3, -2, -7, -22, 2, 1000000})
	expected := 4
	assert.Equal(t, result, expected, fmt.Sprintf("For array  [1, -1000000, 7, 5, 3, -2, -7, -22, 2, 1000000] you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{1, -1000000, 7, 5, 3, -2, -7, -22, 1000000}, " the first missing positive integer is: ", MissingInteger([]int{1, -1000000, 7, 5, 3, -2, -7, -22, 1000000}))
}

// array with all negative value
func TestMissingInteger_allNeg(t *testing.T) {
	result := MissingInteger([]int{-1, -5, -22, -104})
	expected := 1
	assert.Equal(t, result, expected, fmt.Sprintf("For array  [-1, -5, -22, -104] you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{-1, -5, -22, -104}, " the first missing positive integer is: ", MissingInteger([]int{-1, -5, -22, -104}))
}

// array with just 1 as elements
func TestMissingInteger_justOne(t *testing.T) {
	result := MissingInteger([]int{1})
	expected := 2
	assert.Equal(t, result, expected, fmt.Sprintf("For array  [1] you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{1}, " the first missing positive integer is: ", MissingInteger([]int{1}))
}

// array with one positive element different from 1
func TestMissingInteger_OnePosElemSlice(t *testing.T) {
	result := MissingInteger([]int{2})
	expected := 1
	assert.Equal(t, result, expected, fmt.Sprintf("For array  [2] you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{2}, " the first missing positive integer is: ", MissingInteger([]int{2}))
}

// array with one element equal to the max negative value available as input
func TestMissingInteger_OneElemMaxNeg(t *testing.T) {
	result := MissingInteger([]int{-1000000})
	expected := 1
	assert.Equal(t, result, expected, "Input of the tp to MissingInteger equal to:  -1000000, 1000000 ")
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{-1000000}, " the first missing positive integer is: ", MissingInteger([]int{-1000000}))
}

// array with one element equal to the max positive value available as input
func TestMissingInteger_MinMaxVal(t *testing.T) {
	result := MissingInteger([]int{-1000000, 1000000})
	expected := 1
	assert.Equal(t, result, expected, fmt.Sprintf("For array  -1000000, 1000000 you should have as return %v not %v", expected, result))
	fmt.Println("Correct result!")
	fmt.Println("For array ", []int{-1000000, 1000000}, " the first missing positive integer is: ", MissingInteger([]int{-1000000, 1000000}))
}
