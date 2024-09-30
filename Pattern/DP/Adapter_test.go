package DP

import (
	"math"
	"testing"
)

// TestPeg calls Adapter. It runs with two input, one a cilindric plug, one square.
// First it checks if adapter function do not convert the cilinder and convert the square as expected
// then check if it calculates correctly the fit func
func TestPeg(t *testing.T) {

	var (
		kidsgame  = RoundHole{Radius: 21}
		dumbSuare = SquarePeg{Width: 22}
		circle    = CirclePeg{Radius: 22}
	)
	if kidsgame.Adapt(circle) != 22 {
		t.Fatalf(`kidsgame.Adapt( circle ) = %v expect "", 22`, kidsgame.Adapt(circle))
	}
	if kidsgame.Fit(circle) == true {
		t.Fatalf(`kidsgame.Fit( %v,) want "", false`, circle)
	}
	if kidsgame.Adapt(dumbSuare) != ((float32)(math.Sqrt(2)) * 22 / 2) {
		t.Fatalf(`kidsgame.Adapt( dumbSuare ) = %v expect "", %v`, kidsgame.Adapt(dumbSuare), ((float32)(math.Sqrt(2)) * 22 / 2))
	}
	if kidsgame.Fit(dumbSuare) == false {
		t.Fatalf(`kidsgame.Fit( %v,) want "", true`, dumbSuare)
	}
	//return nil

}
