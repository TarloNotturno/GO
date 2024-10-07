package Caterpillar

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var A []int = []int{10, 2, 5, 1, 8, 12}

func TestTriangleCheck(t *testing.T) {
	assert.Equal(t, true, IsATriangle(A[0], A[2], A[5]))
	assert.Equal(t, false, IsATriangle(A[0], A[1], A[2]))
}

func TestCodilityEx(t *testing.T) {
	assert.Equal(t, 4, CountTriangles(A))

}

func TestSmallTriangle(t *testing.T) {
	assert.Equal(t, 0, CountTriangles([]int{}))
	assert.Equal(t, 0, CountTriangles([]int{1, 2, 3}))
	assert.Equal(t, 1, CountTriangles([]int{4, 2, 3}))
}
