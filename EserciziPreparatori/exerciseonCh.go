package main

import (
	"os"
	"strings"

	multithread "main.go/Multithread"
	mytimers "main.go/Mytimers"
)

func threadExercise(numberOfThreads int, numberOfSum int, callingList int, printThreadsResult bool) {
	// read files with multy thread
	multithread.OpenAndDivideFile(int64(numberOfThreads), printThreadsResult)
	multithread.Execute_Concurrency(numberOfSum)
	currentdirectory, err := os.Getwd()
	if err == nil {
		multithread.MyRead(strings.Join([]string{currentdirectory, "inputFile"}, "\\"), 3, false)
	}

	//usage of timers
	mytimers.Scheduler(callingList, printThreadsResult)
}
