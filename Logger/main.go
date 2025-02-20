package main // package main
import (
	"main.go/mylogger"
)

func main() {

	windowInfo, err0 := mylogger.NewLogger("Info")
	if err0 != nil {
		panic(err0)
	}
	windowErr, err01 := mylogger.NewLogger("Err")
	if err0 != nil {
		panic(err01)
	}

	//list of log variable writing on a file
	destinationfile, openingErr := mylogger.OpenFile("log.txt") //open file

	if openingErr != nil {
		panic(openingErr)
	}

	fileErr, err2 := mylogger.NewLogger("Error", destinationfile)
	if err2 != nil {
		panic(err2)
	}

	fileInfo, err1 := mylogger.NewLogger("Info", destinationfile)
	if err1 != nil {
		panic(err1)
	}

	//logging to windows
	windowInfo.MyLog("This is a test message")
	windowInfo.MyLog("This is another test message")
	windowErr.MyLog("This is an error message")

	//logging to file
	fileInfo.MyLog("This is a test message")
	fileInfo.MyLog("This is second test message")
	fileErr.MyLog("And this is a bad message")

	closingErr := mylogger.CloseFile(destinationfile)
	if closingErr != nil {
		panic(openingErr)
	}

}
