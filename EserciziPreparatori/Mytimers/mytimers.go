package mytimers

import (
	"fmt"
	"time"
)

func run500ms(i bool) {
	if i {
		fmt.Println("this code runs every 500 ms")
	}
}

func run100ms(i bool) {
	if i {
		fmt.Println("this code runs every 100 ms", i)
	}
}

func run1s(i bool) {
	if i {
		fmt.Println("THIS CODE RUNS EVERY 1 s", i)
	}
}

func Scheduler(N int, printThreadsResult bool) {
	sched3ms := time.NewTicker(500 * time.Millisecond)
	sched100ms := time.NewTicker(100 * time.Millisecond)
	sched1s := time.NewTicker(time.Second)
	for i := 0; i < N; i++ {
		select {
		case <-sched3ms.C:
			run500ms(printThreadsResult)
		case <-sched100ms.C:
			run100ms(printThreadsResult)
		case <-sched1s.C:
			run1s(printThreadsResult)

		}
	}
}
