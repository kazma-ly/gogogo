package logx

import (
	"log"
	"sync"
)

var (
	// Logs 默认的日志记录
	DefaultLog *Log
	// 锁
	lock sync.Locker
)

func GetDefaultLog() *Log {
	lock.Lock()
	defer lock.Unlock()

	if DefaultLog == nil {
		DefaultLog = new("[SERVER] ")
	} else { // 使用新的文件
		fileInfo, err := DefaultLog.logFile.Stat()
		if err != nil {
			log.Panicf("file info error: %s", err.Error())
		}
		if fileInfo.Size() > 1024*8*1024 {
			DefaultLog.logFile.Close()
			DefaultLog = new("[SERVER] ")
		}
	}
	return DefaultLog
}

// LogInfo 默认的日志记录
func LogInfo(message ...interface{}) {
	GetDefaultLog().Log(message)
}

// LogInfo 默认的日志记录
func LogF(format string, message ...interface{}) {
	GetDefaultLog().LogF(format, message)
}
