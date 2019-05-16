package main

import "fmt"
import "lession/queue/queue"

func main() {
	q := queue.Queue{1}
	q.Push("avc")
	q.Push(3.3)
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())
}
