package geometry

import (
	"errors"
	"fmt"

	"main.go/basefunction"
)

type CartesianFigure interface {
	ObtainPlot() CartesianPlot
	GetFunc() string
}

type Point struct {
	X float32
	Y float32
}

// a*Y = m*b*X + c this allows line parallel to Y and X
type StraightLine struct {
	A bool    `json:"y present"`
	B bool    `json:"x present"`
	M float32 `json:"m"`
	C float32 `json:"c"`
}

// Cartesian plan plot
type CartesianPlot struct {
	Design []string
	MaxVal float32
	Step   float32
}

func (line StraightLine) GetFunc() string {
	functionForm := ""
	if line.A {
		functionForm = "y = "
	}
	if line.B {
		functionForm = functionForm + "x "
	}
	if int(line.C*100)/100 != 0 {
		if line.B && line.C > 0 {
			functionForm = functionForm + fmt.Sprintf("+ %f", line.C)
		} else {
			functionForm = functionForm + fmt.Sprintf(" %f", line.C)
		}
	}
	if line.A == false {
		functionForm = functionForm + " = 0"
	}
	return functionForm
}

func (line StraightLine) ObtainPlot() CartesianPlot {
	var plotRes CartesianPlot
	//find intersection with Y
	Y0 := line.C
	X0 := -line.C / line.M
	modY := Y0
	modX := X0
	if modY < 0 {
		modY = -modY
	}
	if modX < 0 {
		modX = -modX
	}
	app := basefunction.MyMax(modX, modY)
	plotRes.Design = make([]string, 0)
	plotRes.MaxVal = 2 * app
	magnify := float32(1)
	incre := float32(5)
	for magnify*app < incre {
		magnify = 10 * magnify
	}
	plotRes.Step = 1 / magnify
	for yapp := -2 * app; yapp <= 2*app; yapp = yapp + plotRes.Step {
		y := float32(int(yapp*magnify)) / magnify
		currentLine := ""
		var currentPoint Point
		currentPoint.X = float32(int((y-line.C)/line.M*magnify)) / magnify
		currentPoint.Y = y
		for xapp := -(2 * app); xapp <= 2*app; xapp = xapp + plotRes.Step {
			x := float32((int(xapp * magnify))) / magnify
			if x != currentPoint.X {
				if x == 0 {
					if y == 0 {
						currentLine = currentLine + "+"
					} else {
						currentLine = currentLine + "|"
					}

				} else if y == 0 {
					currentLine = currentLine + "-"
				} else {
					currentLine = currentLine + " "
				}
			} else {
				currentLine = currentLine + "*"
			}
		}
		plotRes.Design = append(plotRes.Design, currentLine)
	}
	return plotRes
}

func (line StraightLine) PrintLine() {
	val := line.ObtainPlot()
	val.Plot()
}

func (plotRes CartesianPlot) Plot() {
	for _, val := range plotRes.Design {
		fmt.Println(val)
	}

}

func StraightLineFrom2Points(p1 Point, p2 Point) (line StraightLine, e error) {
	//Y=mX+c -> m = (Y1-c)/X1   c = Y2-X2*(Y1-c)/X1
	e = nil
	if p1.X > p2.X-0.00001 && p1.X < p2.X+0.00001 {
		//the Point are just the same one. NOTOK!
		if p1.Y > p2.Y-0.00001 && p1.Y < p2.Y+0.00001 {
			e = errors.New("input Points are coincident, this gives a line bundle, NOTOK!")
			return
		}
		// X1 and X2 are equal, this means the line is parallel to Y aXis
		line.A = false
		line.B = true
		line.C = -p1.X
		line.M = 1

	} else if p1.Y > p2.Y-0.00001 && p1.Y < p2.Y+0.00001 {
		//we got a line parallel to X
		line.A = true
		line.B = false
		line.C = p1.Y
		line.M = 1
	} else {
		line.C = p2.Y - p2.X*(p1.Y-p2.Y)/(p2.X-p1.X)
		line.M = (p1.Y - p2.Y) / (p2.X - p1.X)
		line.A = true
		line.B = true

	}
	return
}
