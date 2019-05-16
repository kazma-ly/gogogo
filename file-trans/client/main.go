package main

import (
	"bufio"
	"fclient/handler"
	"fclient/util"
	"flag"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	c := &handler.Controller{
		StrCh:   make(chan string, 1),
		CloseCh: make(chan string, 1),
	}

	base := *flag.String("p", "/temp/", "保存文件的路径")
	addr := *flag.String("a", "127.0.0.1:1331", "ip端口")
	flag.Parse()

	util.LogDate(fmt.Sprintf("保存路径：%v", base))

	scan := bufio.NewScanner(os.Stdin)

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		panic(err)
	}

	handler.Handler(conn, c, base)

	for scan.Scan() {
		input := scan.Text()
		if strings.Compare(input, "over") == 0 {
			c.CloseCh <- "see you"
			return
		}
		c.StrCh <- input
	}
}
