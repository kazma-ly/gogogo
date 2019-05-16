package main

import (
	"fmt"
	"sync"
	"time"
)

type atomicInt struct {
	value int
	lock  sync.Mutex
}

// type atomicInt int

func (i *atomicInt) increment() {
	i.lock.Lock()
	defer i.lock.Unlock()
	i.value++
}

func (i *atomicInt) get() int {
	i.lock.Lock()
	defer i.lock.Unlock()
	return i.value
}

func main() {
	// atomic.AddInt32()

	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	fmt.Println(a.get())
	time.Sleep(time.Second)
}
