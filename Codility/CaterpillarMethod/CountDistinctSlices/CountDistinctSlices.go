package CountDistinctSlices

type Caterpillar struct {
	begin int
	end   int
}

func CaterpillarSlice(A []int, delta int) Caterpillar {
	curentCaterpillar := Caterpillar{begin: delta, end: delta}
	if len(A) <= 0 {
		return curentCaterpillar
	}
	//curentCaterpillar.end++
	i := 0
	for i < len(A) {
		if i != 0 {
			if A[0] != A[i] {
				curentCaterpillar.end = i + delta
			} else {
				i = len(A) + 1
			}
		}
		i++
	}

	return curentCaterpillar
}

func sumResults(begin int, end int) int {
	result := 0
	for i := end - begin + 1; i > 0; i-- {
		result = result + i
	}
	return result
}

func Solution(M int, A []int) int {
	firstCaterpillar := CaterpillarSlice(A, 0)
	var currentCaterpillar Caterpillar
	result := 0
	if len(A) > 1 {
		for i, _ := range A {
			//if i != 0 {
			if firstCaterpillar.begin == firstCaterpillar.end {
				result = result + 1
				firstCaterpillar = CaterpillarSlice(A[i:], i)
			} else {
				currentCaterpillar = CaterpillarSlice(A[i:], i)
				if currentCaterpillar.end != firstCaterpillar.end {
					result = result + sumResults(firstCaterpillar.begin, currentCaterpillar.begin)
					firstCaterpillar = CaterpillarSlice(A[currentCaterpillar.begin+1:], currentCaterpillar.begin+1)
				}
			}
			//}
		}
		result = result + sumResults(firstCaterpillar.begin, currentCaterpillar.begin)
		return result
	} else {
		return len(A)
	}

}
