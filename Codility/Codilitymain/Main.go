package main

import (
	"fmt"

	"example.com/GenomicRangeQuery"
	"example.com/MaxCounters"
)

func main() {
	MaxCounters.MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4})
	//fmt.Println("MaxCounter Test with ", []int{3, 4, 4, 6, 1, 4, 4}, "results in :", MaxCounters.MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4}))
	fmt.Println("CAGCCTA", GenomicRangeQuery.GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6}))
	//GenomicRangeQuery.GenomicMinWeight("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6})
}
