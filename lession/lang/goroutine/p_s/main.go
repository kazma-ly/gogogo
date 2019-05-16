package main

import (
	"sync"
	"time"
)

// publish发布者和subscriber订阅者

type (
	subscriber chan interface{}         // 订阅者 是一个管道
	topicFunc  func(v interface{}) bool // 主题 是一个过滤器
)

// 发布者对象
type Publisher struct {
	m           sync.RWMutex
	buffer      int
	timeout     time.Duration
	subscribers map[subscriber]topicFunc
}

// 构建一个发布者对象, 可以设置发布超时时间和缓存队列长度
func NewPublisher(publishTimeout time.Duration, buffer int) *Publisher {
	return &Publisher{
		buffer:      buffer,
		timeout:     publishTimeout,
		subscribers: make(map[subscriber]topicFunc),
	}
}

// 添加一个新的订阅者，订阅过滤器筛选后的主题
func (p *Publisher) SubscribeTopic(topic topicFunc) chan interface{} {
	ch := make(chan interface{}, p.buffer)
	p.m.Lock()
	p.subscribers[ch] = topic
	p.m.Unlock()
	return ch
}

// 添加一个新订阅者, 订阅全部频道
func (p *Publisher) Subscribe() {
	p.SubscribeTopic(nil)
}

func main() {

}
