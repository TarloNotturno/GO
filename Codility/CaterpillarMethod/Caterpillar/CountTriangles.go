package Caterpillar

import "fmt"

func IsATriangle(a int, b int, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return true
	} else {
		return false
	}
}

func CaterpillarTriangle(A []int) (currentSol int) {
	a := A[0]
	i := 1
	j := 2
	for i < len(A) {
		j = i + 1
		for j < len(A) {
			if IsATriangle(a, A[i], A[j]) {
				fmt.Println(a, A[i], A[j])
				currentSol++
			}
			j++
		}
		i++
	}
	return currentSol
}

func CountTriangles(A []int) int {
	if len(A) < 3 {
		return 0
	}
	sol := 0
	app := A
	q := len(app)
	for q >= 2 {
		sol = sol + CaterpillarTriangle(app)
		app = app[1:]
		q = len(app)
	}
	return sol
}
