package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindLocalMax(t *testing.T) {

	var (
		pointOfInterest peakAndDepth
	)

	A := []int{20, 4, 10, 1, 3, 5, 7}
	FindLocalMax(A, &pointOfInterest)
	assert.Equal(t, []int{0, 2, 6}, pointOfInterest.listOfPeak)
	assert.Equal(t, []int{0, 1, 3}, pointOfInterest.listOfDepth)

	A = []int{1, 4, 1, 5, 6, 2, 4, 3, 1, 5, 2, 1, 8}
	pointOfInterest.listOfPeak = []int{}
	pointOfInterest.listOfDepth = []int{}
	FindLocalMax(A, &pointOfInterest)
	assert.Equal(t, []int{1, 4, 6, 9, 12}, pointOfInterest.listOfPeak)
	assert.Equal(t, []int{0, 2, 5, 8, 11}, pointOfInterest.listOfDepth)

	A = []int{1, 4, 1, 5, 6, 2, 4, 3, 1, 5, 2, 1, 8, 20, 4, 10, 1, 3, 5, 7}
	pointOfInterest.listOfPeak = []int{}
	pointOfInterest.listOfDepth = []int{}
	FindLocalMax(A, &pointOfInterest)
	assert.Equal(t, []int{1, 4, 6, 9, 13, 15, 19}, pointOfInterest.listOfPeak)
	assert.Equal(t, []int{0, 2, 5, 8, 11, 14, 16}, pointOfInterest.listOfDepth)

}

func TestClearPeaks(t *testing.T) {

	var (
		pointOfInterest peakAndDepth
	)

	A := []int{1, 4, 1, 5, 6, 2, 7, 3, 4, 1, 5, 2, 2, 8, 20, 4, 10, 1, 3, 5, 7}
	pointOfInterest.listOfPeak = []int{}
	pointOfInterest.listOfDepth = []int{}
	FindLocalMax(A, &pointOfInterest)
	ClearPeaks(A, &pointOfInterest)
	assert.Equal(t, []int{1, 4, 6, 14, 16, 20}, pointOfInterest.listOfPeak)

	A = []int{1, 6, 2, 3, 1, 4, 3, 4, 3, 2, 3, 5, 2, 3, 1, 4, 3}
	pointOfInterest.listOfPeak = []int{}
	pointOfInterest.listOfDepth = []int{}
	FindLocalMax(A, &pointOfInterest)
	ClearPeaks(A, &pointOfInterest)
	assert.Equal(t, []int{1, 11, 15}, pointOfInterest.listOfPeak)
}

func TestFunctional(t *testing.T) {

	assert.Equal(t, 6, FloodDepth([]int{20, 4, 10, 1, 3, 5, 7}))
	assert.Equal(t, 5, FloodDepth([]int{1, 9, 8, 7, 4, 8, 7, 10, 31, 2}))
	assert.Equal(t, 6, FloodDepth([]int{3, 1, 10, 4, 20}))
	assert.Equal(t, 4999, FloodDepth([]int{5000, 2000, 5000, 1, 3000, 4000, 5000, 1, 2, 2, 1, 2, 5000}))
	assert.Equal(t, 2, FloodDepth([]int{1, 3, 2, 1, 2, 1, 5, 3, 3, 4, 2}))
}
