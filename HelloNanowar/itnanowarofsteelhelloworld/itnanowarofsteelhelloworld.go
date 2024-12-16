package itnanowarofsteelhelloworld

import "fmt"

type HelloWorldMainLauncerClass struct {
	message string
	foo     int
}

func HelloWorldMainLauncerClassInit(foo int) HelloWorldMainLauncerClass {
	var h HelloWorldMainLauncerClass
	h.foo = foo
	h.message = "Hello World!"
	return h
}

func (h *HelloWorldMainLauncerClass) SongRefrain() {

	for i := 0; i < h.foo; i++ {
		fmt.Println(h.message)
	}
	pippo := 0
	for pippo < h.foo {
		fmt.Println(h.message)
		pippo++
	}
}
