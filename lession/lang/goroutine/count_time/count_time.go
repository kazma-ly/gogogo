package main

import (
	"log"
	"time"
)

func main() {
	ch := make(chan string, 3)

	for i := 0; i < 20; i++ {
		go func() {
			ch <- "1"
			sayHello()
			<-ch
		}()
	}

	for {

		select {
		case v := <-ch:
			log.Println(v)
		default:

		}
	}
}

func sayHello() {
	log.Println("Hello World")
	time.Sleep(1 * time.Second)
}
