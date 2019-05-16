package queue

// 先进先出 FIFO

type Queue []interface{}

func (q *Queue) Push(val interface{}) {
	*q = append(*q, val)
}

func (q *Queue) Pop() interface{} {
	val := (*q)[0]
	*q = (*q)[1:]
	return val
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
