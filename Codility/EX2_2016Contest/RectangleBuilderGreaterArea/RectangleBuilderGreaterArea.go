package RectangleBuilderGreaterArea

import (
	"sort"
)

func RectangleBuilderGreaterArea(A []int, X int) int {
	numberOfFences := 0
	piecesMap := make(map[int]int)
	elemMoreThanTwo := make([]int, 0)
	for _, wood := range A {
		piecesMap[wood] = piecesMap[wood] + 1
		if piecesMap[wood] == 2 {
			elemMoreThanTwo = append(elemMoreThanTwo, wood)
		}
	}

	sort.Ints(elemMoreThanTwo)

	N := len(elemMoreThanTwo)

	for i, value := range elemMoreThanTwo {
		index := N - 1
		area := value * elemMoreThanTwo[index]
		if index != i {
			for index > i && area >= X {
				numberOfFences++
				index--
				area = value * elemMoreThanTwo[index]
			}
			if area >= X && index != i {
				numberOfFences++
				if piecesMap[value] > 3 && value*value >= X {
					numberOfFences++
				}
			}
		} else if piecesMap[value] > 3 && value*value >= X {
			numberOfFences++
		}
	}

	return numberOfFences

}
