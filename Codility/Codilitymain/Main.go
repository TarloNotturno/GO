package main

import (
	"fmt"
	"log"

	"example.com/GenomicRangeQuery"
	"example.com/MaxCounters"
	"example.com/MissingInteger"
)

func main() {
	fmt.Println("MaxCounters ###############################################################")
	//MaxCounters.MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4})
	fmt.Println("MaxCounter Test with ", []int{3, 4, 4, 6, 1, 4, 4}, "results in :", MaxCounters.MaxCounters(5, []int{3, 4, 4, 6, 1, 4, 4}))

	fmt.Println("GenomicRangeQuery ###############################################################")
	val, error := GenomicRangeQuery.GenomicRangeQuery("CAGCCTA", []int{2, 5, 0}, []int{4, 5, 6})
	if error == nil {
		fmt.Println("The weight related to 'CAGCCTA' for input", []int{2, 5, 0}, []int{4, 5, 6}, "is equal to:\n", val)
	} else {
		log.Fatal(error)
	}

	//the next lines are to test an error condition, so are comment out in order to test the missinginteger file
	//val, error = GenomicRangeQuery.GenomicRangeQuery("CAGCCTA", []int{2, 5, 9}, []int{4, 5, 6})
	//if error == nil {
	//	fmt.Println("The weight related to 'CAGCCTA' for input", []int{2, 5, 0}, []int{4, 5, 6}, "is equal to:\n", val)
	//} else {
	//	log.Fatal(error)
	//}

	fmt.Println("MissingInteger ###############################################################")
	fmt.Println("For array ", []int{1, 3, 6, 4, 1, 2}, " the first missing positive integer is: ", MissingInteger.MissingInteger([]int{1, 3, 6, 4, 1, 2}))
	// for more example look at MissingInteger_test.go

}
