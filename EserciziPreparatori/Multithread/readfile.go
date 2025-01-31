package multithread

/*
Supponiamo di avere un file di testo di grandi dimensioni che contiene migliaia o milioni di righe. Il tuo compito è scrivere un programma che legga il contenuto di questo file in parallelo utilizzando il multithreading per migliorare le prestazioni.

-Obiettivi:
Divisione del lavoro: Scrivi un programma che suddivida il file in più porzioni e le legga contemporaneamente utilizzando più thread.
Gestione delle risorse: Assicurati che il programma gestisca correttamente le risorse, come la lettura del file e la sincronizzazione dei thread.
Risultati: Una volta completata la lettura, il programma deve restituire il contenuto del file unito, rispettando l'ordine delle righe, ma senza duplicati (se un thread legge la stessa riga in modo parallelo, non devono esserci duplicazioni).
Prestazioni: Sfrutta il multithreading per ottenere miglioramenti nelle prestazioni rispetto alla lettura sequenziale del file.

-Requisiti:
Utilizza un massimo di 4 thread per leggere il file in parallelo.
Ogni thread deve leggere una porzione diversa del file (ad esempio, il file potrebbe essere suddiviso in 4 blocchi).
Una volta che tutti i thread hanno terminato la lettura, unisci i risultati senza duplicati.
Puoi utilizzare qualsiasi linguaggio di programmazione che supporta il multithreading (ad esempio Python, Java, C++, etc.).
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const BYTESLICE = 1000

// CREATE A STRUCT FOR INFO RELATED TO A READED LINE, WILL BE SENT THROUGH CH
type sentData struct {
	lineReaded    string
	wordReaded    []string
	orderOfThread int64
}

// INPUT STRUCT FOR THE THREAD FUNCTIONS
type fileReadInput struct {
	scanner         *bufio.Scanner
	chanSem         chan bool
	fileNotFinished chan bool
	currentLine     chan sentData
	wg              *sync.WaitGroup
	dimByte         int64
	deltaRead       int64
}

func myMax(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func myMin(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func (input fileReadInput) mergeInfo(keepRead *bool, occurrences map[string]int, fullTextCopy *string) {
	defer input.wg.Done()
	unorderedString := make(map[int64]string)
	lastThread := int64(0)
	for {
		// check if the file is totaly read
		*keepRead = <-input.fileNotFinished
		// obtain the data related to the last string read
		text := <-input.currentLine

		unorderedString[text.orderOfThread] = text.lineReaded
		lastThread = myMax(lastThread, text.orderOfThread)
		for _, currentWord := range text.wordReaded {
			if value, ok := occurrences[currentWord]; ok {
				occurrences[currentWord] = value + 1
			} else {
				occurrences[currentWord] = 1
			}
		}
		if !*keepRead {
			goto end
		}
	}
end:
	{
		//copy string to the output
		for i := int64(0); i <= lastThread; i++ {
			*fullTextCopy = *fullTextCopy + unorderedString[i]
		}
		fmt.Println(*fullTextCopy)
	}
}

func (input fileReadInput) readLine(file *os.File) {
	defer input.wg.Done()
	// text is the variable reading BYTESLICE byte
	text := make([]byte, BYTESLICE)
	// check if the file is finished reading and obtain the line
	var endOfFile error
	_, endOfFile = file.ReadAt(text, int64(input.deltaRead*BYTESLICE))

	readedLine := string(text)
	fileToBeRead := endOfFile == nil && readedLine != ""

	// SEND IF THE FILE IS ENDED TO mergeInfo THREAD ###############################################
	input.fileNotFinished <- fileToBeRead
	textSplit := strings.Split(readedLine, " ")
	// SEND THE LINE INFO TO mergeInfo THREAD ######################################################
	input.currentLine <- sentData{
		lineReaded:    readedLine,
		wordReaded:    textSplit,
		orderOfThread: input.deltaRead}
	// EMPTY ONE SEMAPHORE SLOT CAUSE THE THREAD IS DONE ###########################################
	<-input.chanSem

}

func OpenAndDivideFile(maxNumbThreads int64) { /* max number of thread input
	is equal to 4 as per exercise request */

	//open file
	file, err := os.Open("Manzoni.txt")
	if err != nil {
		log.Fatal(err)
	}
	app, _ := file.Stat()
	// find the dimension of file in order to find later when the thread executed is the last one needed,
	// its execution will finish the read of the file
	dimensionFile := app.Size()

	// select how many read thread execute maximum
	nThreads := myMin(myMax(5, int64(dimensionFile/BYTESLICE)), maxNumbThreads)
	// create a buffer to read the file
	scanner := bufio.NewScanner(file)

	// variable reporting the occurences of words inside the file read
	occurrences := make(map[string]int, nThreads)

	// two wait group, one for read one for parse the file
	var wgRead sync.WaitGroup
	var wgMerge sync.WaitGroup

	// chanSem IS THE CHANNEL TO STOP THE EXECUTION OF MORE THREAD THAN REQUESTED ########################
	chanSem := make(chan bool, nThreads)
	// channel variable reporting if we finished read the file
	fileNotFinished := make(chan bool)
	fullText := ""
	// variable reporting if the exchange of fileNotFinished report an end of file
	keepRead := true

	// line read by the read thread
	currentLine := make(chan sentData)

	i := int64(0)

	inputMerge := fileReadInput{
		scanner:         scanner,
		fileNotFinished: fileNotFinished,
		currentLine:     currentLine,
		wg:              &wgMerge}

	wgMerge.Add(1)
	// MERGE INFO READ THE INFO SENT FROM readLine threads and once the file is full read merge all lines
	// in a variable reporting the full text "fullText", on occurrences finds all the word occurrencies in
	// the text file
	go inputMerge.mergeInfo(&keepRead, occurrences, &fullText)
	var inputRead fileReadInput
	for keepRead {
		dimByte := myMin(myMax(dimensionFile-i*BYTESLICE, 0), BYTESLICE)
		// IF THE FILE IS FULLY READ DO NOT ADD NEW THREAD ##########################################
		// else add a read thread once we had an ok from semaphore max number of thread
		if dimensionFile-i*BYTESLICE > 0 {
			wgRead.Add(1)
			chanSem <- true

			inputRead = fileReadInput{
				scanner:         scanner,
				chanSem:         chanSem,
				fileNotFinished: fileNotFinished,
				currentLine:     currentLine,
				wg:              &wgRead,
				dimByte:         dimByte,
				deltaRead:       i}

			i++
			go inputRead.readLine(file)
		}
	}
	wgRead.Wait()
	close(currentLine)
	close(chanSem)
	wgMerge.Wait()

	file.Close()

	fmt.Println(fullText)
	fmt.Println()
	fmt.Println()
	fmt.Println(occurrences)

}
