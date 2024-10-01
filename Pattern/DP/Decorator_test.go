package DP

import (
	"testing"
)

// TestPeg calls Adapter. It runs with two input, one a cilindric plug, one square.
// First it checks if adapter function do not convert the cilinder and convert the square as expected
// then check if it calculates correctly the fit func
func TestDecorator(t *testing.T) {

	var (
		input1 ExampleDecorator //{Begin: "lol", End: "#"}
		input2 ExampleDecorator //{Begin: "_", End: "!"}
		input3 ExampleDecorator //{Begin: "/", End: "\""}
		int1   NumericDecorator //{Begin: 1, End: 22}
	)
	input1.Initialize("lol", "#", &input2)
	input2.Initialize("_", "!", &input3)
	input3.Initialize("/", "\"", nil)
	int1.Initialize(1, 22, input1)
	check := int1.Decorator("la bella lava il fosso")
	if check != "1lol_/la bella lava il fosso\"!#22" {
		t.Fatalf(`input3.Decorator("la bella lava il fosso") want: "1lol_/la bella lava il fosso"!#22 obtain %v`, check)
	}
	//return nil

}
