package main

import (
	"encoding/json"
	"fmt"

	"main.go/basefunction"
	"main.go/geometry"
)

func main() {
	fmt.Println(basefunction.MyOwnPwr(22, 3))
	fmt.Println(basefunction.MyOwnRoot(10648, 3))
	fmt.Println(basefunction.MyOwnRoot(8, 3))
	basefunction.SetConstants(0.0000001)
	fmt.Println(basefunction.MyOwnRoot(121, 2))
	exp := 2.7183
	fmt.Println(basefunction.MyOwnLog(148.4132, exp))

	lin1, _ := geometry.StraightLineFrom2Points(geometry.Point{X: 1, Y: 4}, geometry.Point{X: 3, Y: 1})
	lin2, _ := geometry.StraightLineFrom2Points(geometry.Point{X: 1, Y: 1}, geometry.Point{X: 3, Y: 4})
	//lin1.PrintLine()
	//lin2.PrintLine()
	//fmt.Printf("y = %vx +%v", lin1.m, lin1.c)
	var m geometry.Point
	b, err := json.Marshal(geometry.Point{X: 3, Y: 1})
	if err == nil {
		err = json.Unmarshal(b, &m)
	}
	fmt.Println(m)

	var n geometry.StraightLine
	b, err = json.Marshal(lin1)
	if err == nil {
		err = json.Unmarshal(b, &n)
	}
	if err == nil {
		fmt.Println(n.GetFunc())
	}
	fmt.Println(lin2.GetFunc())
	app, _ := geometry.StraightLineFrom2Points(geometry.Point{X: 1, Y: 4}, geometry.Point{X: 1, Y: 1})
	fmt.Println(app.GetFunc())
	app, _ = geometry.StraightLineFrom2Points(geometry.Point{X: 2, Y: 4}, geometry.Point{X: 1, Y: 4})
	fmt.Println(app.GetFunc())

}
