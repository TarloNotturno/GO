package main

import (
	"fmt"
	"math"
	"slices"
)

type mySlice struct {
	begin  int
	end    int
	weight int
}

/* calculcate the increment for each element of the vector */
func calcWeight(A []int) []int {
	weight := make([]int, len(A)+1)
	weight[0] = 0
	for i, value := range A {
		weight[i+1] = weight[i] + value
	}
	return weight

}

/* calculate the optimal division of a slice into two */
func reduceSliceWithBin(weight []int, begin int, end int) int {
	i := end
	sol := end
	i_old := end
	weightA := weight[end] - weight[begin]
	weightB := 0
	weightOldA := weightA + 1
	weightOldB := weightB + 1
	lengthFind := end
	begin_curr := begin
	end_curr := end
	for lengthFind > 1 {
		i_old = i
		if weight[i]-weight[begin_curr] > weight[end_curr]-weight[i] {
			i = (i + begin_curr) / 2
			end_curr = i
		} else {
			i = (i + end_curr) / 2
			begin_curr = i
		}
		weightOldA = weightA
		weightOldB = weightB
		weightA = weight[i] - weight[begin]
		weightB = weight[end] - weight[i]
		lengthFind = i_old - i
		if math.Max((float64)(weightA), (float64)(weightB)) < math.Max((float64)(weightOldA), (float64)(weightOldB)) {
			sol = i
		}
	}
	return sol
}

func Solution1(K int, M int, A []int) int {
	end := len(A)
	if end < 1 {
		return end
	}
	weight := calcWeight(A)
	N := len(A)
	var arraySliced []mySlice
	number_of_slice := 0

	end = int(math.Max(float64(N-K+1), 0))
	begin := end

	arraySliced = append(arraySliced, mySlice{begin: 0, end: end, weight: (weight[end] - weight[0])})

	for number_of_slice < K-1 {
		if end < N {
			end++
		}
		arraySliced = append(arraySliced, mySlice{begin: begin, end: end, weight: (weight[end] - weight[begin])})
		number_of_slice++
		begin = end
	}

	update := true
	for update {
		update = false
		for j, value := range arraySliced {
			/* reduce with previous one */
			if j > 0 {
				i := reduceSliceWithBin(weight, arraySliced[j-1].begin, arraySliced[j].end)
				oldBegin := arraySliced[j].begin
				arraySliced[j].begin = i
				arraySliced[j].weight = weight[arraySliced[j].end] - weight[i]

				arraySliced[j-1].end = i
				arraySliced[j-1].weight = weight[i] - weight[arraySliced[j-1].begin]
				update = update || oldBegin != i
			}

			/* reduce with next one */
			if j < len(arraySliced)-1 {
				oldBegin := arraySliced[j+1].begin
				i := reduceSliceWithBin(weight, arraySliced[j].begin, arraySliced[j+1].end)
				arraySliced[j+1].begin = i
				arraySliced[j+1].weight = weight[arraySliced[j+1].end] - weight[i]
				arraySliced[j].end = i
				arraySliced[j].weight = weight[i] - weight[value.begin]
				update = update || oldBegin != i
			}
		}
	}

	result := 0
	for _, value := range arraySliced {
		if value.weight > result {
			result = value.weight
		}
	}

	return result
}

func Solution(K int, M int, A []int) (result int) {
	i := 0
	begin := 0
	end := len(A)
	if end < 1 {
		return end
	}
	weight := calcWeight(A)

	var (
		arraySliced []mySlice
	)

	arraySliced = append(arraySliced, mySlice{begin: begin, end: end, weight: (weight[end] - weight[begin])})
	maxSliceId := 0
	for i < K-1 {
		result = 0
		middle := reduceSliceWithBin(weight, begin, end)
		if i < K-1 {
			arraySliced = slices.Insert(arraySliced, maxSliceId+1, mySlice{begin: middle, end: end, weight: (weight[end] - weight[middle])})
			arraySliced[maxSliceId].end = middle
			arraySliced[maxSliceId].weight = (weight[middle] - weight[begin])
			if len(arraySliced) >= maxSliceId+3 {
				arraySliced[maxSliceId+2].begin = end
			}
		}

		for j, slice_curr := range arraySliced {
			if result <= slice_curr.weight {
				begin = slice_curr.begin
				end = slice_curr.end
				result = slice_curr.weight
				maxSliceId = j
			}
		}
		i++
	}

	//update := true
	//for update {
	//	update = false
	for j := (len(arraySliced) - 1); j > 0; j-- {
		//for k := range arraySliced {
		/* reduce with previous one */
		if j > 0 {
			i := reduceSliceWithBin(weight, arraySliced[j-1].begin, arraySliced[j].end)
			//oldBegin := arraySliced[j].begin
			arraySliced[j].begin = i
			arraySliced[j].weight = weight[arraySliced[j].end] - weight[i]

			arraySliced[j-1].end = i
			arraySliced[j-1].weight = weight[i] - weight[arraySliced[j-1].begin]
			//update = update || oldBegin != i
		}

		/* reduce with next one */
		if j < len(arraySliced)-1 {
			//oldBegin := arraySliced[j+1].begin
			i := reduceSliceWithBin(weight, arraySliced[j].begin, arraySliced[j+1].end)
			arraySliced[j+1].begin = i
			arraySliced[j+1].weight = weight[arraySliced[j+1].end] - weight[i]
			arraySliced[j].end = i
			arraySliced[j].weight = weight[i] - weight[arraySliced[j].begin]
			//update = update || oldBegin != i
		}
	}
	//}

	result = 0
	for _, slice_curr := range arraySliced {
		if result <= slice_curr.weight {
			result = slice_curr.weight
		}
	}

	return result
}

func main() {
	fmt.Println("opt", Solution1(3, 5, []int{2}), "working", Solution(3, 5, []int{2}))
	A := []int{2, 1, 5, 1, 2, 2, 2}
	fmt.Println("opt", Solution(3, 5, A), "expected", 6)
	fmt.Println("opt", Solution(4, 5, A), "expected", 5)
	//
	B := []int{11, 11, 2, 1, 12, 5, 12, 1, 2, 2, 2}
	fmt.Println("opt Solution(4, 12, B)", Solution1(4, 12, B), "working Solution(4, 12, B)", Solution(4, 12, B))
	fmt.Println("opt Solution(5, 12, B)", Solution(5, 12, B), "working Solution(5, 12, B)", Solution(5, 12, B))
	B = []int{11, 11, 2, 1, 12, 5, 1, 2, 2, 2}
	fmt.Println("opt Solution(5, 12, B)", Solution(5, 12, B), "working Solution(5, 12, B)", Solution(5, 12, B))
	c := []int{1000}
	fmt.Println("opt Solution(3, 5, c)", Solution(3, 5, c), "working Solution(3, 5, c)", Solution(3, 5, c))
	D := []int{5, 3}
	fmt.Println("opt Solution(1, 10, D)", Solution(1, 10, D), "working Solution(1, 10, D)", Solution(1, 10, D))
	fmt.Println("opt Solution(3, 5, c)", Solution(3, 5, c), "working Solution(3, 5, c)", Solution(3, 5, c))
	fmt.Println("opt Solution(4, 10, []int{3, 4, 5, 6, 7, 8})", Solution(4, 10, []int{3, 4, 5, 6, 7, 8}), "working Solution(4, 10, []int{3, 4, 5, 6, 7, 8})", Solution(4, 10, []int{3, 4, 5, 6, 7, 8}))

}
