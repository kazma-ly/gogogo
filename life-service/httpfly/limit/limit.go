package limit

import (
	"sync"
	"time"

	"golang.org/x/time/rate"
)

/**
 * 令牌限流
 */
type (
	// Vistor 浏览结构体
	Vistor struct {
		limit    *rate.Limiter // 限流器
		lastView time.Time     // 最后访问时间 、删除一些太久远的请求 省的浪费内存
	}

	// Limiter 限流器
	Limiter struct {
		vistors map[string]Vistor
		mtx     sync.Mutex
	}
)

// New 初始化一个限流器
func New() *Limiter {
	l := &Limiter{vistors: make(map[string]Vistor)}
	go l.clearWork()
	return l
}

// Allow 是否允许访问
func (l *Limiter) Allow(ip string) bool {
	l.mtx.Lock()
	defer l.mtx.Unlock()
	vistor, ok := l.vistors[ip]
	if !ok {
		l.vistors[ip] = Vistor{rate.NewLimiter(2, 5), time.Now()}
		return true
	}
	vistor.lastView = time.Now()
	return vistor.limit.Allow()
}

func (l *Limiter) clearWork() {
	for {
		time.Sleep(1 * time.Minute)
		l.mtx.Lock()
		vs := l.vistors
		for ip, vistor := range vs {
			if time.Now().Sub(vistor.lastView) > 3*time.Minute {
				delete(vs, ip)
			}
		}
		l.mtx.Unlock()
	}
}
