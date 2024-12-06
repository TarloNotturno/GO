package main

import (
	"fmt"

	"main.go/DwarfsRafting"
	"main.go/RectangleBuilderGreaterArea"
)

func main() {
	RectangleBuilderGreaterArea.RectangleBuilderGreaterArea([]int{2, 5, 5, 2, 1, 1, 3, 1, 1}, 5)
	//fmt.Println(RectangleBuilderGreaterArea.RectangleBuilderGreaterArea([]int{}, 5))
	//fmt.Println(RectangleBuilderGreaterArea.RectangleBuilderGreaterArea([]int{2, 2, 1, 1, 3, 1, 1}, 5))
	//fmt.Println(RectangleBuilderGreaterArea.RectangleBuilderGreaterArea([]int{2, 2, 5, 5, 1, 1, 3, 1, 1, 5, 5}, 5))
	fmt.Println(DwarfsRafting.DwarfsRafting(4, "1B 1C 4B 1D 2A", "3B 2D"))
	//fmt.Println(DwarfsRafting.Solution(2, "", "1B"))
	fmt.Println(DwarfsRafting.DwarfsRafting(4, "1B 1A 2A", "3C 4C"))
	fmt.Println(DwarfsRafting.DwarfsRafting(26, "1B 1A 2A", "3C 4C"))
	fmt.Println(DwarfsRafting.DwarfsRafting(12, "2D 11J 3B 9G 2F 2A 4F 5B 5E 6C 7J 9K 10H ", "1E 2A 2D 2F 3B 4F 5B 6C 5E 9G 7J 9K 10H 11J "))
}
