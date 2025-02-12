package mytimers

import (
	"fmt"
	"time"
)

func run500ms(i int) {
	fmt.Println("this code runs every 500 ms", i)
}

func run100ms(i int) {
	fmt.Println("this code runs every 100 ms", i)
}

func run1s(i int) {
	fmt.Println("THIS CODE RUNS EVERY 1 s", i)
}

func Scheduler(N int) {
	sched3ms := time.NewTicker(500 * time.Millisecond)
	sched100ms := time.NewTicker(100 * time.Millisecond)
	sched1s := time.NewTicker(time.Second)
	for i := 0; i < N; i++ {
		select {
		case <-sched3ms.C:
			run500ms(i)
		case <-sched100ms.C:
			run100ms(i)
		case <-sched1s.C:
			run1s(i)

		}
	}
}
