package main

import "fmt"

func main() {
	number := "DC CX1XXV2III"
	fmt.Println(FromRomanToArab(number), "expected 738")
	number = "MCMXCIV"
	fmt.Println(FromRomanToArab(number), "expected 1994")
	number = "CMXX"
	fmt.Println(FromRomanToArab(number), "expected 920")
	number = "MMXX"
	fmt.Println(FromRomanToArab(number), "expected 2020")
	number = "MMXXI"
	fmt.Println(FromRomanToArab(number), "expected 2021")
	number = "LXV"
	fmt.Println(FromRomanToArab(number), "expected 65")
	number = "LlXVv" // this is not a valid roman number cause
	//double v should be X and same for L and C
	fmt.Println(FromRomanToArab(number), "expected 120")
	number = "X"
	fmt.Println(FromRomanToArab(number), "expected 10")
	number = " m2C 45r3M saa23 22 3x!c5q33.i 434qv"
	fmt.Println(FromRomanToArab(number), "expected 1994")
}
