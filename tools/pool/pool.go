package pool

import (
	"runtime"
	"sync"
)

type (
	// Task 任务
	Task func() error
	// CallBack 完成任务的回调
	CallBack func()

	// GoroutinePool goroutine池
	GoroutinePool struct {
		Queue     chan Task // 维持住的任务队列
		Total     int
		finshback CallBack
		wg        sync.WaitGroup // 阻塞
		result    chan string
	}
)

// New 初始化线程池
func New(total int, callBack CallBack) GoroutinePool {
	if total <= 0 {
		total = runtime.NumCPU() // cpu核心数
	}
	return GoroutinePool{
		Queue:     make(chan Task, total),
		Total:     total,
		finshback: callBack,
	}
}

// Execute 执行(添加)任务
func (gp *GoroutinePool) Execute(task Task) {
	gp.Queue <- task
}

// Callback 设置回调
func (gp *GoroutinePool) Callback(cb CallBack) {
	gp.finshback = cb
}

// Start 开始任务
func (gp *GoroutinePool) Start() {
	gp.wg.Add(1)
	go func() {
		defer gp.wg.Done()
		for index := 0; index < gp.Total; index++ {
			task := <-gp.Queue
			go task()
		}
		if gp.finshback != nil {
			gp.finshback()
		}
		close(gp.Queue)
	}()

}
