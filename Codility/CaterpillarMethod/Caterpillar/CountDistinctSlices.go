package Caterpillar

type Caterpillar struct {
	begin  int
	end    int
	number int
}

// Function that taking a slice calculate the combination of subslices with unique element
func sumResults(begin int, end int) int {
	result := 0
	for i := end - begin + 1; i > 0; i-- {
		result = result + i
	}
	return result
}

// "main" function to evaluate all slices with unique element in an array
func CountDistinctSlices(M int, A []int) int {
	occurrencies := MappingOccurrency(A)
	lenFullSlice := len(A) - 1
	firstCaterpillar := CaterpillarSlice(A, occurrencies, 0, lenFullSlice)
	var currentCaterpillar Caterpillar
	result := 0
	old_sum := 0
	if len(A) > 1 {
		i := 0
		for i < len(A) {
			currentCaterpillar = CaterpillarSlice(A[i:], occurrencies, i, lenFullSlice)
			if currentCaterpillar.end < firstCaterpillar.end || firstCaterpillar.end < i {
				if currentCaterpillar.end <= firstCaterpillar.end {
					result = result + sumResults(firstCaterpillar.begin, currentCaterpillar.end)
					old_sum = old_sum + sumResults(currentCaterpillar.begin+1, currentCaterpillar.end)
					occurrencies[firstCaterpillar.number] = occurrencies[firstCaterpillar.number][1:]
					if firstCaterpillar.number != currentCaterpillar.number {
						occurrencies[currentCaterpillar.number] = occurrencies[currentCaterpillar.number][1:]
					}
					firstCaterpillar = CaterpillarSlice(A[currentCaterpillar.begin+1:], occurrencies, currentCaterpillar.begin+1, lenFullSlice)
					i = currentCaterpillar.begin + 1

				} else {
					result = result + sumResults(firstCaterpillar.begin, firstCaterpillar.end)
					old_sum = old_sum + sumResults(firstCaterpillar.begin+1, firstCaterpillar.end)
					occurrencies[firstCaterpillar.number] = occurrencies[firstCaterpillar.number][1:]
					firstCaterpillar = CaterpillarSlice(A[firstCaterpillar.begin+1:], occurrencies, firstCaterpillar.begin+1, lenFullSlice)
					i = firstCaterpillar.begin
				}

			}
			i++
		}
		//if !(firstCaterpillar.begin == 1 && firstCaterpillar.end == lenFullSlice) {
		result = result + sumResults(firstCaterpillar.begin, firstCaterpillar.end) - old_sum
		//}
		return result
	} else {
		return len(A)
	}

}

// function that for an element of the array reports an array with its occurrencies
func MappingOccurrency(A []int) map[int][]int {
	ricorrenze := make(map[int][]int)
	for i := range A {
		ricorrenze[A[i]] = append(ricorrenze[A[i]], i)
	}
	return ricorrenze
}

// find the maximum slice length for a given number
func CaterpillarSlice(A []int, OccurrencyMap map[int][]int, delta int, lenFullSlice int) Caterpillar {
	var currentCaterpillar Caterpillar
	if len(A) > 0 {
		currentnumber := A[0]
		currentCaterpillar.number = currentnumber
		if len(OccurrencyMap[currentnumber]) > 1 {
			currentCaterpillar.begin = OccurrencyMap[currentnumber][0]
			currentCaterpillar.end = OccurrencyMap[currentnumber][1]
			if OccurrencyMap[currentnumber][1] < lenFullSlice || currentCaterpillar.begin == 0 {
				currentCaterpillar.end = currentCaterpillar.end - 1
			}
		} else {
			currentCaterpillar.begin = delta
			currentCaterpillar.end = len(A) + delta - 1
		}
	} else {
		currentCaterpillar.begin = delta
		currentCaterpillar.end = len(A) + delta - 1
	}

	return currentCaterpillar
}
