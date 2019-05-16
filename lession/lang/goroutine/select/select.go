package main

import (
	"fmt"
	"math/rand"
	"time"
)

func work(id int, c chan int) {
	for v := range c {
		time.Sleep(time.Second)
		fmt.Printf("id = %d, result = %d\n", id, v)
	}
}

func createWork(id int) chan int {
	c := make(chan int)
	go work(id, c)
	return c
}

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWork(0)

	var values []int                   // 队列
	tm := time.After(10 * time.Second) // 10秒后发送一个chan
	tt := time.Tick(time.Second)       // 每秒发送一个chan
	for {
		var activeWorker chan int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n := <-c1:
			values = append(values, n)
		case n := <-c2:
			values = append(values, n)
		case activeWorker <- activeValue:
			values = values[1:]
		case <-time.After(800 * time.Millisecond): // 超过800毫秒
			fmt.Println("timeout")
		case <-tt:
			fmt.Println("queue len:", len(values))
		case <-tm:
			fmt.Println("bye")
			return
		}
	}
}
