package multithread

import (
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
