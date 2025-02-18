package main

import reflectionex "main.go/ReflectionEx"

func main() {
	printBaseResult := false // boolean disabling the print of calculated results
	// base exercise:
	// execute conditioned printing
	// fill a table given a point in the table
	baseExercise(printBaseResult)
	// exercises on multithreads
	printThreadsResult := true
	threadExercise(4, 100, 2, printThreadsResult)

	reflectionex.SendOrder()

}
