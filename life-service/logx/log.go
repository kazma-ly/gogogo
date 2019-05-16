package logx

import (
	"log"
	"os"
)

type (
	// Log 日志结构体
	Log struct {
		filelog    *log.Logger
		consolelog *log.Logger
		logFile    *os.File
	}
)

// New 初始化
func NewLog(pre string) *Log {
	logFile := checkPathAndGetFile()
	logInfo := log.New(logFile, pre, log.LstdFlags|log.Lshortfile)
	return &Log{
		filelog:    logInfo,
		consolelog: log.New(os.Stdout, pre, log.LstdFlags|log.Lshortfile),
		logFile:    logFile,
	}
}

// Log 日志记录
func (mylog *Log) Log(message ...interface{}) {
	mylog.filelog.Println(message)
	mylog.consolelog.Println(message)
}

// LogF 日志记录
func (mylog *Log) LogF(foramt string, message ...interface{}) {
	mylog.filelog.Printf(foramt, message)
	mylog.consolelog.Printf(foramt, message)
}
