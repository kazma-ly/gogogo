package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (c *ConcurrentEngine) Run(seeds ...Request) {
	// in := make(chan Request)
	out := make(chan ParseResult)

	//c.Scheduler.ConfigureMasterWorkerChan(in)
	c.Scheduler.Run()

	for i := 0; i < c.WorkerCount; i++ {
		createWorker(c.Scheduler.WorkerChan(), out, c.Scheduler)
	}

	for _, r := range seeds {
		c.Scheduler.Submit(r)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			go func() { c.ItemChan <- item }() // 保存
		}
		for _, request := range result.Requests {
			c.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, n ReadyNotifier) {
	go func() {
		for {
			n.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
