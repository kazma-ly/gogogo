package scheduler

import "lession/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Run() {
	s.workerChan = make(chan engine.Request)
}

// 继续往里面送
func (s *SimpleScheduler) Submit(request engine.Request) {
	go func() { s.workerChan <- request }()
}
