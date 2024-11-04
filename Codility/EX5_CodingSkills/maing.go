package main

import (
	"fmt"
	"strconv"
)

func ParkingBill(E string, L string) int {
	var (
		Ehour, Emin, Lhour, Lmin int
		e                        error
	)
	Ehour, e = strconv.Atoi(E[:2])
	if e == nil {
		Emin, e = strconv.Atoi(E[3:5])
		if e == nil {
			Lhour, e = strconv.Atoi(L[:2])
			if e == nil {
				Lmin, e = strconv.Atoi(L[3:5])
				if e != nil {
					return 0
				}
			} else {
				return 0
			}

		} else {
			return 0
		}
	} else {
		return 0
	}
	hourPerm := Lhour - Ehour
	if Emin < Lmin {
		hourPerm++
	}
	payment := 5
	hourPerm--
	if hourPerm > 0 {
		payment = payment + hourPerm*4
	}
	return payment
}

func ParityDegree(N int) int {
	solution := 0
	for N%2 == 0 {
		solution++
		N = N / 2
	}
	return solution
}

func ThreeLetters(A int, B int) string {
	var (
		sol string
	)
	a := "a"
	b := "b"
	if A < B {
		app := A
		A = B
		B = app
		a = "b"
		b = "a"
	}
	diffAB := A - B
	for A > 0 {
		for diffAB > 0 {
			sol = sol + a
			A--
			if A > 0 {
				sol = sol + a
				A--
			}
			if B > 0 {
				sol = sol + b
				B--
			}
			diffAB = A - B
		}
		if A > 0 {
			sol = sol + a
			A--
			if B > 0 {
				sol = sol + b
				B--
			}
		}
	}
	return sol
}

func main() {
	fmt.Println(ParkingBill("09:20", "21:15"))
	fmt.Println(ParkingBill("09:00", "21:15"))
	fmt.Println(ParkingBill("09:00", "09:15"))
	fmt.Println(ParityDegree(2313124536))
	fmt.Println(ParityDegree(8))
	fmt.Println(ThreeLetters(5, 3))
	fmt.Println(ThreeLetters(3, 3))
	fmt.Println(ThreeLetters(1, 4))
	fmt.Println(ThreeLetters(2, 6))
	fmt.Println(ThreeLetters(3, 1))
}
