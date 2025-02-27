package main

import "fmt"

func main() {
	number := "DC CX1XXV2III"
	fmt.Println(fromRomanToArab(number), "expected 738")
	number = "MCMXCIV"
	fmt.Println(fromRomanToArab(number), "expected 1994")
	number = "CMXX"
	fmt.Println(fromRomanToArab(number), "expected 920")
	number = "MMXX"
	fmt.Println(fromRomanToArab(number), "expected 2020")
	number = "MMXXI"
	fmt.Println(fromRomanToArab(number), "expected 2021")
	number = "LXV"
	fmt.Println(fromRomanToArab(number), "expected 65")
	number = "LlXVv" // this is not a valid roman number cause
	//double v should be X and same for L and C
	fmt.Println(fromRomanToArab(number), "expected 120")
}
