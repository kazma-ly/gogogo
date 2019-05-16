package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// 生产者消费者模型

// Producer 生产者
func Producer(factor int, out chan<- int) {
	for i := 0; ; i++ {
		out <- i * factor
	}
}

// Consumer 消费者
func Consumer(in <-chan int) {
	for v := range in {
		log.Println(v)
	}
}

func main() {

	ch := make(chan int, 10)

	go Producer(2, ch)
	go Producer(5, ch)
	go Consumer(ch) // 消费者

	// Ctrl+C 退出
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	log.Printf("quit (%v)\n", <-sig)
}
