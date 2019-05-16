package main

import (
	"fmt"
	"time"
)

func work(id int, c chan int) {
	// for {
	// 	v, ok := <-c // 判断chan是否被close
	// 	if !ok {
	// 		break
	// 	}
	// 	fmt.Printf("id = %d, result = %c\n", id, v)
	// }

	for v := range c {
		fmt.Printf("id = %d, result = %c\n", id, v)
	}
}

func createWork(id int) chan int {
	c := make(chan int)
	go work(id, c)
	return c
}

func chanDemo() {
	var channel [10]chan int
	for i := 0; i < 10; i++ {
		channel[i] = createWork(i)
	}

	for i := 0; i < 10; i++ {
		channel[i] <- 'a' + i
	}

	time.Sleep(1 * time.Second)
}

func bufferChannel() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'A' + 1
	c <- 'A' + 2
	c <- 'A' + 3
	c <- 'A' + 4

	time.Sleep(1 * time.Second)
}

func channelClose() {
	c := make(chan int, 3)
	go work(0, c)
	c <- 'A' + 1
	c <- 'A' + 2
	c <- 'A' + 3
	c <- 'A' + 4
	close(c) // 关闭了也能继续接收数据 但收到的是0
	time.Sleep(1 * time.Second)
}

func main() {
	fmt.Println("channel as first-class citizen")
	chanDemo()

	fmt.Println("buffered chanel")
	bufferChannel()

	fmt.Println("close chanel")
	channelClose()
}
