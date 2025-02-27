package multithread

import (
	"fmt"
	"sync"
)

type commonCounter struct {
	counter int32
}

func (c *commonCounter) increment() {
	c.counter++
}

func collectorCount(counter *commonCounter, wg *sync.WaitGroup, flagSum chan bool) {
	defer wg.Done()
	for {
		app := <-flagSum
		if app {
			counter.increment()
		} else {
			goto end
		}
	}
end:
}

func incrementCaller(numIncrement int, wg *sync.WaitGroup, flagSum chan bool) {
	defer wg.Done()
	i := 0
	for i < numIncrement {
		flagSum <- true
		i++
	}

}

func Execute_Concurrency(maxThread int) int32 {
	var wg, wgSum sync.WaitGroup
	counter := commonCounter{counter: 0}
	//semaphore := make(chan bool, maxThread)
	n := 0
	flagSum := make(chan bool)
	wgSum.Add(1)
	go collectorCount(&counter, &wgSum, flagSum)
	for n < maxThread {
		wg.Add(1)
		go incrementCaller(1000, &wg, flagSum)
		n++
	}
	wg.Wait()
	flagSum <- false
	close(flagSum)
	wgSum.Wait()

	return counter.counter
}

/*Titolo dell'esercizio: "Calcolo parallelo della somma"

Descrizione:

Scrivi un programma che calcoli la somma di tutti i numeri interi compresi tra 1 e
𝑁
N (dove
𝑁
N è un numero dato come input). La somma deve essere eseguita in parallelo, dividendo l'intervallo in blocchi e utilizzando più thread per calcolare parzialmente la somma in parallelo.

Requisiti:
L'utente fornirà il valore di
𝑁
N e il numero di thread da utilizzare (ad esempio, 4, 8, ecc.).
Suddividi l'intervallo
[
1
,
𝑁
]
[1,N] in blocchi uguali e calcola la somma parziale di ciascun blocco in un thread separato.
Una volta che tutti i thread hanno completato il loro lavoro, raccogli i risultati e somma le somme parziali.
Il programma deve garantire la sincronizzazione tra i thread per evitare problemi di concorrenza, se necessario.
Dettagli aggiuntivi:
Usa la libreria di threading del linguaggio scelto (in Python, ad esempio, threading).
Ogni thread dovrà calcolare la somma di una parte dell'intervallo.
Puoi usare una struttura come una queue o una lock per raccogliere i risultati dei thread in modo sicuro.
Il programma dovrebbe stampare la somma totale alla fine.*/

func singleSum(start int64, end int64, delta chan int64, wg *sync.WaitGroup) {
	var currentSum int64
	defer wg.Done()
	for i := start; i < end; i++ {
		currentSum = currentSum + i
	}
	delta <- currentSum
}

func collectorSum(recivedPart chan int64, wg *sync.WaitGroup, result *int64) {
	defer wg.Done()

	for {
		increment, app := <-recivedPart
		if app {
			*result = *result + increment
		} else {
			break
		}

	}
}

func BigSum(N int64, nthreadMAx int64) {
	var (
		delta         int64
		firstDeltaAdd int64
		wgSum         sync.WaitGroup
		wgCollector   sync.WaitGroup
		result        int64
	)
	partialSum := make(chan int64, 1)
	//semaphore :=
	if nthreadMAx < N {
		delta = int64(N / nthreadMAx)
		firstDeltaAdd = N % nthreadMAx
	} else {
		delta = int64(1)
	}
	var allocatedRun int64
	wgCollector.Add(1)
	go collectorSum(partialSum, &wgCollector, &result)
	for allocatedRun < nthreadMAx && firstDeltaAdd < N {
		wgSum.Add(1)
		go singleSum(firstDeltaAdd, firstDeltaAdd+delta, partialSum, &wgSum)
		firstDeltaAdd = firstDeltaAdd + delta
		allocatedRun++
	}
	wgSum.Wait()
	close(partialSum)
	wgCollector.Wait()

	fmt.Println(result)

}
