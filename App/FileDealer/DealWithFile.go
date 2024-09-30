package FileDealer

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"sync"
)

func ReadCheck(e error) {
	if e != nil {
		fmt.Println(e)
		panic(e)
	}
}

// func ReadInterfaces(filePath string) (rpList []string, ppList []string)
// func ReadInterfaces(filePath string, wg sync.WaitGroup, rpList chan []string, ppList chan []string) {
func ReadInterfaces(filePath string, wg *sync.WaitGroup, portList chan SWCInterfaces, currentInterface chan SWCInterfaces, endRead chan bool) {

	appList := ReadInFile(filePath)
	portList <- appList
	endRead <- true
	currentInterface <- appList
	wg.Done()
}

// func writeInterfaces(filePath string, wg sync.WaitGroup, rpList chan []string, ppList chan []string) {
func WriteInterfaces(filePath string, wg *sync.WaitGroup, portList chan SWCInterfaces, endRead chan bool, numberOfThread int) {
	// open output file
	fo, err := os.Create(filePath)
	ReadCheck(err)
	processClosed := 0
	for {

		WriteInFile(<-portList, fo)

		// if i recive an and i increment the internal counter
		select {
		case <-endRead:
			processClosed = processClosed + 1
		default:
		}
		// if the number of thread close is equal to the needed one close this thread
		if processClosed >= numberOfThread {
			goto end
		}
	}

end:
	// close fo on exit and check for its returned error
	if err := fo.Close(); err != nil {
		panic(err)
	}
	wg.Done()
}

func WriteInFile(InterfacesList SWCInterfaces, fo *os.File) {
	var err error
	len_Trat := 200
	fmt.Fprint(fo, "\nSWC : "+InterfacesList.fileName+" ")
	for j := 0; j < len_Trat-len("\nSWC : "+InterfacesList.fileName); j++ {
		fmt.Fprint(fo, "-")
	}
	fmt.Fprint(fo, "\n")
	//write into file required ports
	fmt.Fprint(fo, "	Required port:\n")
	ReadCheck(err)
	lenghtLine := 0
	if len(InterfacesList.rpList) > 0 {
		for _, val := range InterfacesList.rpList {
			lenghtLine = lenghtLine + len(val)
			if lenghtLine > len_Trat {
				_, err := fmt.Fprint(fo, "\n")
				lenghtLine = 0
				ReadCheck(err)
			}
			_, err = fmt.Fprint(fo, val+" ")
			ReadCheck(err)
		}
	} else {
		fmt.Fprint(fo, "n.a.")
	}
	//write into file provided ports
	fmt.Fprint(fo, "\n	Provided port:\n")
	ReadCheck(err)
	lenghtLine = 0
	if len(InterfacesList.ppList) > 0 {
		for _, val := range InterfacesList.ppList {
			lenghtLine = lenghtLine + len(val)
			if lenghtLine > len_Trat {
				_, err := fmt.Fprint(fo, "\n")
				lenghtLine = 0
				ReadCheck(err)
			}
			_, err = fmt.Fprint(fo, val+" ")
			ReadCheck(err)
		}
	} else {
		fmt.Fprint(fo, "n.a.")
	}

}

func ReadInFile(filePath string) (appList SWCInterfaces) {
	//open file and read it
	f, err := os.Open(filePath)
	ReadCheck(err)
	Scanner := bufio.NewScanner(f)
	Scanner.Split(bufio.ScanWords)
	// find file name and create keyword to look for
	fileNameApp := strings.Split(filePath, "/")
	fileName := fileNameApp[len(fileNameApp)-1]
	fileNameApp = strings.Split(fileName, ".")
	fileName = fileNameApp[0]
	rpKeyWord := "rp" + fileName
	ppKeyWord := "pp" + fileName
	var (
		app       string
		rpListApp []string
		ppListApp []string
	)
	//look for required and provided port
	required_Port := make(map[string]int)
	provided_Port := make(map[string]int)
	for Scanner.Scan() {
		app = (Scanner.Text())
		//look for required port
		if strings.Contains(Scanner.Text(), rpKeyWord) {
			required_Port[app] = required_Port[app] + 1
		}
		//look for provided port
		if strings.Contains(Scanner.Text(), ppKeyWord) {
			provided_Port[app] = provided_Port[app] + 1
		}
	}
	//close file
	f.Close()
	// Trim required port
	for i, _ := range required_Port {
		val := strings.Split(i, "_")
		if len(val) >= 4 {
			rpListApp = append(rpListApp, val[3])
		}
	}
	// Trim provided port
	for i, _ := range provided_Port {
		val := strings.Split(i, "_")
		if len(val) >= 3 {
			ppListApp = append(ppListApp, val[3])
		}
	}
	appList.fileName = fileName
	appList.ppList = ppListApp
	appList.rpList = rpListApp
	return appList
}

func FindAllASWFiles(repoPath string, destinationFilePath string) {
	var currentFilePath string
	fo, err := os.Create(destinationFilePath)
	ReadCheck(err)
	app, err := os.ReadDir(repoPath)
	if err != nil {
		panic(err)
	}
	for _, currentFolder := range app {
		if currentFolder.IsDir() {
			currentFilePath = repoPath + "/" + currentFolder.Name() + "/code/source/" + currentFolder.Name() + ".c"
			if _, err := os.Stat(currentFilePath); err == nil {
				//fmt.Println(currentFilePath)
				fmt.Fprintln(fo, currentFilePath)

			} else if errors.Is(err, os.ErrNotExist) {
			} else {
				// Schrodinger: file may or may not exist. See err for details.

				fmt.Println(err)

			}

		}
	}
	// close fo on exit and check for its returned error
	if err := fo.Close(); err != nil {
		panic(err)
	}

}
