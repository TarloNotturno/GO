package main

import "fmt"

//"main.go/DwarfsRafting"

func myMin(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	fmt.Println("2017 Contest")
	fmt.Println("1.TennisTournament")
	TennisTournament(5, 3)
	TennisTournament(10, 3)
	fmt.Println("2.SocksLaundering")
	SocksLaundering(2, []int{1, 2, 1, 1}, []int{1, 4, 3, 2, 4})
	//fmt.Println(SocksLaundering(2, []int{1, 2, 1, 1}, []int{1, 4, 3, 2, 4}), "expected 3")
	//fmt.Println(SocksLaundering(5, []int{1, 1, 2}, []int{2, 2, 3}), "expected 2")
	//fmt.Println(SocksLaundering(5, []int{1, 1, 3, 4}, []int{3, 3, 3, 5}), "expected 3")
	//fmt.Println(SocksLaundering(5, []int{}, []int{3, 3, 3, 5, 4, 4, 5, 4, 3}), "expected 2")
	//fmt.Println(SocksLaundering(3, []int{}, []int{3, 3, 3, 5, 4, 4, 5, 4, 3}), "expected 1")
	//fmt.Println(SocksLaundering(5, []int{}, []int{3, 5, 4}))
}
