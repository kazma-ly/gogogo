package util

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"time"
)

// LogDate 输出日志
func LogDate(msg string) {
	log.Printf("%v: "+msg+"\n", time.Now().Format("2006-01-02 15:04:05.999"))
}

// GetMD5Str 获得文件的md5
func GetMD5Str(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := md5.New()
	bs := make([]byte, 1024)

	for {
		n, err := file.Read(bs)
		if err != nil {
			break
		}
		hash.Write(bs[:n])
	}

	return hex.EncodeToString(hash.Sum(nil)), nil
}
