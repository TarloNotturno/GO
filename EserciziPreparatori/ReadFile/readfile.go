package readfile

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

type sentData struct {
	lineReaded    string
	wordReaded    []string
	orderOfThread int64
}

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
		*keepRead = <-input.fileNotFinished
		if *keepRead {
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
		} else {
			goto end
		}
	}
end:
	{
		//copy string to the output
		for i := int64(0); i < lastThread; i++ {
			*fullTextCopy = *fullTextCopy + unorderedString[i]
		}
		fmt.Println(*fullTextCopy)
	}
}

func (input fileReadInput) readLine(file *os.File) {
	defer input.wg.Done()
	//fileToBeRead := true
	text := make([]byte, input.dimByte)
	var endOfFile error
	_, endOfFile = file.ReadAt(text, int64(input.deltaRead*BYTESLICE))

	fileToBeRead := endOfFile == nil && int64(input.deltaRead*BYTESLICE) > 0
	readedLine := string(text)

	input.fileNotFinished <- fileToBeRead
	if fileToBeRead {
		textSplit := strings.Split(readedLine, " ")
		input.currentLine <- sentData{
			lineReaded:    readedLine,
			wordReaded:    textSplit,
			orderOfThread: input.deltaRead}
	}

	<-input.chanSem

}

func OpenAndDivideFile(maxNumbThreads int64) { /* max number of thread input
	is equal to 4 as per exercise request */

	//open file
	file, err := os.Open("ManzoniShort.txt")
	if err != nil {
		log.Fatal(err)
	}
	app, _ := file.Stat()
	dimensionFile := app.Size()

	nThreads := myMin(myMax(5, int64(dimensionFile/BYTESLICE)), maxNumbThreads)

	//textDiv := make([]string, nThreads)

	scanner := bufio.NewScanner(file)

	occurrences := make(map[string]int, nThreads)

	var wgRead sync.WaitGroup
	var wgMerge sync.WaitGroup
	//chanRead := make(chan string, nThreads)
	chanSem := make(chan bool, nThreads)
	fileNotFinished := make(chan bool)
	fullText := ""
	keepRead := true

	// obtain slice of file
	currentLine := make(chan sentData)

	i := int64(0)

	inputMerge := fileReadInput{
		scanner:         scanner,
		fileNotFinished: fileNotFinished,
		currentLine:     currentLine,
		wg:              &wgMerge}

	wgMerge.Add(1)
	go inputMerge.mergeInfo(&keepRead, occurrences, &fullText)

	for keepRead {
		dimByte := myMin(myMax(dimensionFile-i*BYTESLICE, 0), BYTESLICE)
		//if dimByte > 0 {
		wgRead.Add(1)
		chanSem <- true

		inputRead := fileReadInput{
			scanner:         scanner,
			chanSem:         chanSem,
			fileNotFinished: fileNotFinished,
			currentLine:     currentLine,
			wg:              &wgRead,
			dimByte:         dimByte,
			deltaRead:       i}

		go inputRead.readLine(file)
		//}
		i++
	}
	wgRead.Wait()
	close(currentLine)
	close(chanSem)
	wgMerge.Wait()

	file.Close()

	fmt.Println(fullText)

}
