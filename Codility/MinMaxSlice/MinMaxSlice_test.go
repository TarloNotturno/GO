package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmptyArray(t *testing.T) {
	assert.Equal(t, Solution1(3, 5, []int{}), Solution(3, 5, []int{}))
	assert.Equal(t, Solution1(3, 5, []int{2}), Solution(3, 5, []int{2}))
}

func TestFunctional(t *testing.T) {
	A := []int{2, 1, 5, 1, 2, 2, 2}
	assert.Equal(t, 6, Solution(3, 5, A))
	assert.Equal(t, 5, Solution(4, 5, A))
	//
	B := []int{11, 11, 2, 1, 12, 5, 12, 1, 2, 2, 2}
	assert.Equal(t, 24, Solution(4, 12, B))
	assert.Equal(t, 17, Solution(5, 12, B))
	B = []int{11, 11, 2, 1, 12, 5, 1, 2, 2, 2}
	assert.Equal(t, Solution1(5, 12, B), Solution(5, 12, B))
	c := []int{1000}
	assert.Equal(t, Solution1(3, 5, c), Solution(3, 5, c))
	D := []int{5, 3}
	assert.Equal(t, Solution1(1, 5, D), Solution(1, 5, D))
	assert.Equal(t, Solution1(3, 5, c), Solution(3, 5, c))
	assert.Equal(t, Solution1(4, 10, []int{3, 4, 5, 6, 7, 8}), Solution(4, 10, []int{3, 4, 5, 6, 7, 8}))
}

func TestSlicer(t *testing.T) {
	assert.Equal(t, 7, reduceSliceWithBin([]int{0, 11, 22, 24, 25, 37, 42, 54, 55, 57, 59, 61}, 5, 11))
}
