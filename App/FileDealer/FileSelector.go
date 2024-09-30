package FileDealer

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

type SWCInterfaces struct {
	fileName string
	rpList   []string
	ppList   []string
}

func (s SWCInterfaces) PrintInterface() {
	fmt.Println("Software Component: " + s.fileName)
	fmt.Print("Provided port : {")
	for _, val := range s.ppList {
		fmt.Print(val + " ")
	}
	fmt.Print("}\nRequired port :{")
	for _, val := range s.rpList {
		fmt.Print(val + " ")
	}
	fmt.Println("}")
}

func LaunchFileRead(ListOfFile string, ListOfInterfaces *[]SWCInterfaces, wg *sync.WaitGroup, scheduler *sync.WaitGroup, destinationFile string) {
	//open file and read it
	*ListOfInterfaces = nil
	f, err := os.Open(ListOfFile)
	ReadCheck(err)
	Scanner := bufio.NewScanner(f)
	Scanner.Split(bufio.ScanWords)
	var listOfFiles []string
	for Scanner.Scan() {
		app := (Scanner.Text())
		listOfFiles = append(listOfFiles, app)
	}
	NUMBER_OF_FILE := len(listOfFiles)
	listOfDone := make(chan bool, NUMBER_OF_FILE)
	portList := make(chan SWCInterfaces, NUMBER_OF_FILE/2)
	currentInterface := make(chan SWCInterfaces, NUMBER_OF_FILE/2)

	wg.Add(2)

	for _, val := range listOfFiles {
		wg.Add(1)
		go ReadInterfaces(val, wg, portList, currentInterface, listOfDone)
	}
	go SaveInterface(ListOfInterfaces, wg, currentInterface, NUMBER_OF_FILE)
	go WriteInterfaces(destinationFile, wg, portList, listOfDone, NUMBER_OF_FILE)
	wg.Wait()
	scheduler.Done()
}
