package main

/*
I = 1 (uno)
V = 5 (cinque)
X = 10 (dieci)
L = 50 (cinquanta)
C = 100 (cento)
D = 500 (cinquecento)
M = 1000 (mille)
*/

var (
	lowUpConversion map[string]string
	fromCharToInt   map[string]int
)

func init() {
	lowUpConversion = fromLowToUpperCase()
	fromCharToInt = tableConversion()
}

func tableConversion() map[string]int {
	return map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
}

func fromLowToUpperCase() map[string]string {
	return map[string]string{
		"i": "I",
		"v": "V",
		"x": "X",
		"l": "L",
		"c": "C",
		"d": "D",
		"m": "M",
		"I": "I",
		"V": "V",
		"X": "X",
		"L": "L",
		"C": "C",
		"D": "D",
		"M": "M",
	}
}
