package main

import (
	"fmt"
	"main/DP"
	"math/rand/v2"
)

func Initializer(i DP.Factorer) {
	i.Create(max(2, rand.IntN(10)))
}

func PaintAnything(a DP.Painter) {
	a.Paint()
}

func main() {

	fmt.Println("FACTORY METHOD -----------------------------------------------------")
	var shop DP.Shop
	chairs := make([]DP.Chair, 22)
	//chairs2 := make([]DP.Chair, 22)
	for i, _ := range chairs {
		Initializer(&chairs[i])
	}
	wardrobe := make([]DP.Wardrobe, 11)
	for i, _ := range wardrobe {
		wardrobe[i].Create(max(2, rand.IntN(10)))
	}
	shop.Create(chairs, wardrobe)
	shop.ListWearhouse()
	fmt.Println("COMPOSITE METHOD ----------------------------------------------------")

	shapes := make([]DP.Shape, 4)
	shapes[0] = DP.Shape{Shape_type: "triangle"}
	shapes[1] = DP.Shape{Shape_type: "rectangle"}
	shapes[2] = DP.Shape{Shape_type: "triangle"}
	shapes[3] = DP.Shape{Shape_type: "rectangle"}
	blue := DP.Color{Setcolor: "blue"}
	red := DP.Color{Setcolor: "red"}
	orange := DP.Color{Setcolor: "orange"}
	shapes[0].SetColor(&blue)
	shapes[1].SetColor(&red)
	shapes[2].SetColor(&red)
	shapes[3].SetColor(&orange)
	for _, i := range shapes {
		PaintAnything(&i)
	}

	fmt.Println("COMPOSITE METHOD ----------------------------------------------------")
	var nokia DP.Smartphone
	fmt.Printf("smartphone price: %v\n", nokia.Price())
	fmt.Printf("its keybord costs: %v\n", nokia.K_keyboard.Product_Price)

	fmt.Println("PROTOTYPE METHOD ----------------------------------------------------")
	siemens := DP.Copy(nokia).(DP.Smartphone)
	fmt.Printf("smartphone price: %v\n", siemens.Product_Price)
	fmt.Printf("its keybord costs: %v\n", siemens.K_keyboard.Product_Price)

	fumettaro := DP.Copy(shop).(DP.Shop)
	fumettaro.ListWearhouse()

	copy_shape := DP.Copy((DP.Shape_array)(shapes)).(DP.Shape_array)
	for _, i := range copy_shape {
		PaintAnything(&i)
	}

	fmt.Println("ADAPTER METHOD -----------------------------------------------------")
	var (
		kidsgame  = DP.RoundHole{Radius: 123}
		dumbSuare = DP.SquarePeg{Width: 22}
		circle    = DP.CirclePeg{Radius: 22}
	)
	fmt.Printf("Can the toy enter?  %v\n", kidsgame.Fit(circle))
	fmt.Printf("Can the toy enter?  %v\n", kidsgame.Fit(dumbSuare))

	fmt.Println("DECORATOR METHOD ----------------------------------------------------")
	var (
		input1 = DP.ExampleDecorator{Begin: "lol", End: "#"}
		input2 = DP.ExampleDecorator{Begin: "_", End: "!"}
		input3 = DP.ExampleDecorator{Begin: "/", End: "\""}
		int1   = DP.NumericDecorator{Begin: 1, End: 22}
	)
	fmt.Println(int1.Decorator(input1.Decorator(input2.Decorator(input3.Decorator("la bella lava il fosso")))))
}
