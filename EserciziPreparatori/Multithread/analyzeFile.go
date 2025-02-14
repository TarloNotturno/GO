package multithread

import (
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type fileData struct {
	NomeFile     string `json:"fileName"`
	NumeroParole int    `json:"numWord"`
}

type result struct {
	Results []fileData `json:"Results"`
}

func createListOfResult(root string, listOfFile []string, maxNumThread int) []fileData {
	listOfResults := []fileData{}
	var readerWG, collectorWG sync.WaitGroup
	currentResult := make(chan fileData)
	semaphore := make(chan bool, maxNumThread)

	collectorWG.Add(1)
	go collector(currentResult, &listOfResults, &collectorWG)

	//go func() {
	//	defer collectorWG.Done()
	//	for {
	//		// quando ho finito di leggere tutti i file chiudo questa funzione di aggiornamento dell'uscita
	//		i, ok := <-currentResult
	//		if !ok {
	//			goto end
	//		}
	//		listOfResults = append(listOfResults, i)
	//
	//	}
	//end:
	//}()

	for _, currentFile := range listOfFile {
		readerWG.Add(1)
		// scrivo sul semaforo ogni volta che apro un thread, quando il canale è saturo rimango in attesa che un thread della funzione read
		// abbia finito di leggere un file

		semaphore <- true
		go reader(root, currentFile, currentResult, semaphore, &readerWG)

	}

	readerWG.Wait()
	// questi non li chiudevamo mai! se non chiudi i canali la riga i, ok := <-currentResult non torna mai non ok e dunque non si chiude
	// mai la funzione di aggiornamento della lista
	close(semaphore)
	close(currentResult)
	collectorWG.Wait()
	return listOfResults
}

func collector(currentResult chan fileData, listOfResults *[]fileData, collectorWG *sync.WaitGroup) {
	defer collectorWG.Done()
	for {
		// quando ho finito di leggere tutti i file chiudo questa funzione di aggiornamento dell'uscita
		i, ok := <-currentResult
		if !ok {
			goto end
		}
		*listOfResults = append(*listOfResults, i)

	}
end:
}

func reader(root string, fileToCheck string, currentResult chan fileData, semaphore chan bool, wg *sync.WaitGroup) {

	defer func() {
		<-semaphore
		wg.Done()
	}()
	currentfile := filepath.Join(root, fileToCheck)
	if fp, err := os.ReadFile(currentfile); err != nil {
		panic(err)
	} else {
		currentResult <- fileData{NomeFile: currentfile, NumeroParole: len(strings.Split(string(fp), " "))}

	}

}

func MyRead(root string, maxNumThread int, printThreadsResult bool) []byte {

	listOfFile := []string{}
	fileSystem := os.DirFS(root)

	fs.WalkDir(fileSystem, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Fatal(err)
		}
		if strings.Contains(path, ".txt") && !d.IsDir() {
			listOfFile = append(listOfFile, path)
		}
		return nil

	})

	listOfResults := createListOfResult(root, listOfFile, maxNumThread)

	toExport := result{Results: listOfResults}
	var (
		str  []byte
		errM error
	)
	if str, errM = json.Marshal(toExport); errM != nil {
		panic(errM)
	}

	if printThreadsResult {
		for _, result := range listOfResults {
			fmt.Println("number of word in file ", result.NomeFile, ":", result.NumeroParole)
		}
	}

	return str
}
