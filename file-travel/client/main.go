package main

import (
	"bufio"
	"fclient/handler"
	"flag"
	"os"
	"strings"
)

func main() {

	localPath := *flag.String("p", "D:/temp2019/", "保存文件的路径")
	server := *flag.String("a", "127.0.0.1:1331", "ip端口")

	flag.Parse()

	controller := handler.New(handler.Input{
		DownloadPath: localPath,
		Server:       server,
	})

	controller.Handler()

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {

		input := scan.Text()

		if strings.Compare(input, "over") == 0 {
			controller.CloseCh <- "see you"
			return
		}

		controller.MsgCh <- input
	}
}
