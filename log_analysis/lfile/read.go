package lfile

import (
	"bufio"
	"io"
	"log"
	"os"
	"sync"
)

type (
	// ReadFile 读取文件的struct
	ReadFile struct {
		lock  sync.Mutex    // 锁
		path  string        // 文件路径
		f     *os.File      // 文件
		bf    *bufio.Reader // 文件缓冲
		count int           // 用于计算重试次数
	}
)

// New create new ReadFile struct
func New(path string) *ReadFile {
	stat, e := os.Stat(path)
	if e != nil {
		log.Println("warning: file stat check has error: ", e)
	}
	if stat.IsDir() {
		log.Println("error: you must check the path is file, not dir")
		return nil
	}
	return &ReadFile{
		path: path,
	}
}

// GetLine read one line by log file
func (rf *ReadFile) GetLine() string {
	if rf.f == nil || rf.bf == nil {
		rf.init()
	}
	line, _, err := rf.bf.ReadLine()
	if err != nil {
		if rf.count < 3 { // 只重试3次
			e := rf.init()
			if e != nil {
				rf.count++
			} else {
				rf.count = 0
			}
		} else {
			panic("the file can't read: " + err.Error())
		}
		if err != io.EOF {
			log.Println(err)
		}
		return ""
	}
	if len(line) > 0 {
		return string(line)
	}
	return ""
}

func (rf *ReadFile) init() error {
	rf.lock.Lock()
	defer rf.lock.Unlock()
	if rf.f == nil || rf.bf == nil {
		file, e := os.OpenFile(rf.path, os.O_RDONLY, 0755)
		if e != nil {
			log.Println("read file error: ", e)
			return e
		}
		rf.f = file
		rf.f.Seek(0, 2) // read from last
		rf.bf = bufio.NewReader(rf.f)
	}
	return nil
}

// Destroy 销毁
func (rf *ReadFile) Destroy() {
	if rf.f != nil {
		rf.f.Close()
	}
}
