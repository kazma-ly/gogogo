package main

import (
	"flag"
	"fserver/handler"
	"log"
	"net"
)

func main() {

	// 处理flag
	_rootPath := *flag.String("p", "/logs/", "要传输的文件夹根目录")
	listenerPort := flag.String("P", "1331", "监听端口")
	flag.Parse()

	log.Printf("rootPath is: %v \n", _rootPath)

	listenerAddress := "0.0.0.0:" + *listenerPort

	// 开始监听连接
	listener, err := net.Listen("tcp", listenerAddress)
	if err != nil {
		panic(err)
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Printf("accept error %v \n", err)
			continue
		}

		// 处理连接
		go handler.HandleConn(conn, _rootPath)
	}

}
