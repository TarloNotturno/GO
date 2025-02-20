package main // package main
import (
	"main.go/mylogger"
)

func main() {

	windowLog, err0 := mylogger.NewLogger()
	if err0 != nil {
		panic(err0)
	}
	//windowErr, err01 := mylogger.NewLogger()
	//if err0 != nil {
	//	panic(err01)
	//}

	//list of log variable writing on a file
	destinationfile, openingErr := mylogger.OpenFile("log.txt") //open file

	if openingErr != nil {
		panic(openingErr)
	}

	fileLog, err2 := mylogger.NewLogger(destinationfile)
	if err2 != nil {
		panic(err2)
	}

	//fileInfo, err1 := mylogger.NewLogger("Info", destinationfile)
	//if err1 != nil {
	//	panic(err1)
	//}

	//logging to windows
	windowLog.LogInfo("This is a test message")
	windowLog.LogInfo("This is another test message")
	windowLog.LogError("This is an error message")

	//logging to file
	fileLog.LogInfo("This is a test message")
	fileLog.LogInfo("This is second test message")
	fileLog.LogError("And this is a bad message")

	closingErr := mylogger.CloseFile(destinationfile)
	if closingErr != nil {
		panic(openingErr)
	}

}
