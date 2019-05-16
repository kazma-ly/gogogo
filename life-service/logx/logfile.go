package logx

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"time"
)

const (
	LOG_PRE = "/logs"
	LOG_SUR = ".log"
)

func makeFilePath() string {
	return LOG_PRE + time.Now().Format("2006_01_02-15_04_05") + LOG_SUR
}

func checkPathAndGetFile() *os.File {
	saveFilePath := filepath.ToSlash(makeFilePath())
	err := os.MkdirAll(path.Dir(saveFilePath), os.ModePerm)
	if err != nil {
		log.Panicf("Error create log file: %s", saveFilePath)
	}
	f, err := os.OpenFile(saveFilePath, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	if err != nil {
		log.Panicf("Error open log file: %s", saveFilePath)
	}
	return f
}
