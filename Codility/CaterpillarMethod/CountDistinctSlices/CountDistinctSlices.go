package CountDistinctSlices

import "fmt"

type Caterpillar struct {
	begin  int
	end    int
	number int
}

func CaterpillarSlice(A []int, delta int) Caterpillar {
	currentCaterpillar := Caterpillar{begin: delta, end: delta}
	if len(A) <= 0 {
		return currentCaterpillar
	}
	//curentCaterpillar.end++
	i := 0
	currentCaterpillar.number = A[0]
	for i < len(A) {
		if i != 0 {
			if A[0] != A[i] {
				currentCaterpillar.end = i + delta
			} else {
				i = len(A) + 1
			}
		}
		i++
	}

	return currentCaterpillar
}

func sumResults(begin int, end int) int {
	result := 0
	for i := end - begin + 1; i > 0; i-- {
		result = result + i
	}
	return result
}

func Solution(M int, A []int) int {
	//fmt.Println(A)
	//MappingOccurrency(A)
	firstCaterpillar := CaterpillarSlice(A, 0)
	var currentCaterpillar Caterpillar
	result := 0
	old_sum := 0
	if len(A) > 1 {
		i := 0
		for i < len(A) {
			if firstCaterpillar.begin == firstCaterpillar.end {
				if firstCaterpillar.end < (len(A) - 1) {
					result = result + 1
					firstCaterpillar = CaterpillarSlice(A[firstCaterpillar.begin+1:], firstCaterpillar.begin+1)
				}
			} else {
				currentCaterpillar = CaterpillarSlice(A[i:], i)
				if currentCaterpillar.end < firstCaterpillar.end || firstCaterpillar.end < i {
					if currentCaterpillar.end <= firstCaterpillar.end {
						result = result + sumResults(firstCaterpillar.begin, currentCaterpillar.end)
						old_sum = old_sum + sumResults(currentCaterpillar.begin+1, currentCaterpillar.end)
						firstCaterpillar = CaterpillarSlice(A[currentCaterpillar.begin+1:], currentCaterpillar.begin+1)
					} else {
						result = result + sumResults(firstCaterpillar.begin, firstCaterpillar.end)
						old_sum = old_sum + sumResults(firstCaterpillar.begin+1, firstCaterpillar.end)
						firstCaterpillar = CaterpillarSlice(A[firstCaterpillar.begin+1:], firstCaterpillar.begin+1)
						i = firstCaterpillar.begin
					}

				}
			}
			i++
		}
		fmt.Println(A, result, sumResults(firstCaterpillar.begin, firstCaterpillar.end), old_sum)
		result = result + sumResults(firstCaterpillar.begin, firstCaterpillar.end) - old_sum
		return result
	} else {
		return len(A)
	}

}

//func MappingOccurrency(A []int) {
//	ricorrenze := make(map[int][]int)
//	for i, _ := range A {
//		ricorrenze[A[i]] = append(ricorrenze[A[i]], i)
//	}
//	fmt.Println(ricorrenze)
//
//}
