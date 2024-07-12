package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	fibonacci_numb := 1
	var fibonacci_old int
	var fibonacci_app int
	return func() int {
		fibonacci_numb += fibonacci_old
		fibonacci_old = fibonacci_app
		fibonacci_app = fibonacci_numb
		return fibonacci_numb
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
