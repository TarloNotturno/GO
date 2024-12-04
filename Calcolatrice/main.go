package main

import (
	"encoding/json"
	"fmt"

	"main.go/SaveData"
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

	var m geometry.Point
	b, err := json.Marshal(geometry.Point{X: 3, Y: 1})
	if err == nil {
		err = json.Unmarshal(b, &m)
	}
	fmt.Println(m)

	listOfFigure := []geometry.StraightLine{}
	app, err := geometry.StraightLineFrom2Points(geometry.Point{X: 1, Y: 4}, geometry.Point{X: 1, Y: 1})
	if err == nil {
		listOfFigure = append(listOfFigure, app)
	}
	app, err = geometry.StraightLineFrom2Points(geometry.Point{X: 2, Y: 4}, geometry.Point{X: 1, Y: 4})
	if err == nil {
		listOfFigure = append(listOfFigure, app)
	}
	app, err = geometry.StraightLineFrom2Points(geometry.Point{X: 1, Y: 4}, geometry.Point{X: 3, Y: 1})
	if err == nil {
		listOfFigure = append(listOfFigure, app)
	}
	app, err = geometry.StraightLineFrom2Points(geometry.Point{X: 1, Y: 1}, geometry.Point{X: 3, Y: 4})
	if err == nil {
		listOfFigure = append(listOfFigure, app)
	}

	var savingData SaveData.FileData
	savingData.StraightLines = listOfFigure
	savingData.Points = make([]geometry.Point, 0)
	savingData.Points = append(savingData.Points, m)
	savingData.Plots = make([]geometry.CartesianPlot, 0)
	savingData.Plots = append(savingData.Plots, listOfFigure[0].ObtainPlot())
	SaveData.WriteData("calculatorSavings.json", savingData, 2, listOfFigure, listOfFigure[2])
	fmt.Println(app.GetFunc())
	SaveData.ReadData("calculatorSavings.json")
}
