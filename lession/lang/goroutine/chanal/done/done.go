package main

import (
	"fmt"
)

func doWork(id int, w work) {
	for v := range w.in {
		fmt.Printf("id = %d, result = %c\n", id, v)
		w.done <- true
	}
}

type work struct {
	in   chan int
	done chan bool
}

func createWork(id int) work {
	w := work{in: make(chan int), done: make(chan bool)}
	go doWork(id, w)
	return w
}

func chanDemo() {
	var works [10]work
	for i := 0; i < 10; i++ {
		works[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		works[i].in <- 'a' + i
	}

	for i := 0; i < 10; i++ {
		<-works[i].done
	}

	// time.Sleep(1 * time.Second)
}

func main() {
	chanDemo()
}
