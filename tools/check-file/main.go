package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"

	"flag"
	"fmt"
	"os"
	"strconv"
)

var queueSize = 1
var queue chan int

func main() {

	// 异常拦截处理
	defer func() {
		if p := recover(); p != nil {
			fmt.Printf("发生了错误: %v\n", p)
		}
		bs := bufio.NewScanner(os.Stdin)
		fmt.Print("按任意字母退出")
		for bs.Scan() {
			break
		}
	}()

	// 输入参数
	path := flag.String("p", "", "指定要检查的文件路径")
	flag.Parse()

	fmt.Println("路径:", *path)

	if len(*path) <= 0 {
		panic("请输入文件路径以判断文件的信息")
	}

	queue = make(chan int, queueSize)

	go getKey(*path)

	for index := 0; index < queueSize; index++ {
		<-queue
	}
}

func getKey(path string) {
	orginFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	if err != nil {
		panic("打开文件出错: " + err.Error())
	}
	defer orginFile.Close()
	fileInfo, err := orginFile.Stat()

	var filesize int64
	if err != nil {
		panic("获取文件信息出错: " + err.Error())
	} else {
		filesize = fileInfo.Size()
	}

	written := 0 // 总读取大小
	bs := make([]byte, 10240)
	_md5 := md5.New()
	_sha1 := sha1.New()
	for {
		n, err := orginFile.Read(bs)
		if err != nil { // EOF or other
			break
		}
		bsslice := bs[:n]
		_md5.Write(bsslice)
		_sha1.Write(bsslice)

		written = written + len(bsslice)
		// \r 表示刷新当前位置
		fmt.Fprintf(os.Stdout, "读取中: %f%%\r", float64(written)/float64(filesize)*100)
	}

	m := float64(written) / float64(1024*1024) // M
	fmt.Println("读取文件大小: " + strconv.FormatFloat(m, 'f', 2, 64) + "M(step by 1024)")

	fmt.Printf("[ md5]: %X\n", _md5.Sum(nil))
	fmt.Printf("[sha1]: %X\n", _sha1.Sum(nil))

	queue <- 1
}
