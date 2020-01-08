package main

import (
	"flag"
	"fserver/handler"
)

func main() {

	path := *flag.String("p", "D:/bdpan", "要传输的文件夹根目录")
	port := *flag.String("P", "1331", "监听端口")

	flag.Parse()

	controller := handler.New(handler.Input{
		RootPath: path,
		Port:     port,
	})

	controller.Start()

}
