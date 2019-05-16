package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
	"time"
)

// URLInfo URL的信息
type URLInfo struct {
	Method    string
	Host      string
	ReadBytes []byte
}

func main() {
	port := *flag.String("p", "1080", "listing port")
	flag.Parse()

	address := ":" + port
	fmt.Println(address)
	lis, err := net.Listen("tcp", address)
	checkErr(err)

	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		conn.SetDeadline(time.Now().Add(1 * time.Minute))
		go handleAccept(conn)
	}
}

func handleAccept(conn net.Conn) {
	urlInfo, err := GetMethodAndHost(conn)
	if err != nil {
		log.Printf("获得请求方法和主机地址失败: %v\n", err)
		return
	}
	address, err := GetTrueAddress(urlInfo.Method, urlInfo.Host)
	if err != nil {
		log.Printf("获得真实地址失败: %v\n", err)
		return
	}

	//获得了请求的host和port，就开始拨号吧
	server, err := net.Dial("tcp", address)
	if err != nil {
		log.Printf("拨号失败: %v\n", err)
		return
	}

	if urlInfo.Method == "CONNECT" { // 网页
		fmt.Fprint(conn, "HTTP/1.1 200 Connection established\r\n\r\n")
	} else {
		server.Write(urlInfo.ReadBytes)
	}

	//进行转发
	go io.Copy(server, conn)
	io.Copy(conn, server)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// GetMethodAndHost 获得请求方法和URL主机地址
func GetMethodAndHost(conn net.Conn) (*URLInfo, error) {
	b := make([]byte, 1024)
	n, err := conn.Read(b) // 读取前1024个字节的数据
	if err != nil {
		return nil, err
	}

	// 从请求行中获得请求方法和请求的主机地址
	var method, host string
	linePosition := bytes.IndexByte(b, '\n')
	requestLine := string(b[:linePosition])         // GET http://www.baidu.com/ HTTP/1.1
	fmt.Sscanf(requestLine, "%s%s", &method, &host) // 读取
	return &URLInfo{method, host, b[:n]}, nil
}

// GetTrueAddress 获得真正的请求地址(HOST:PORT)
func GetTrueAddress(method, host string) (string, error) {
	forwardURL, err := url.Parse(host)
	if err != nil {
		return "", err
	}

	// 最终地址
	var address string
	if forwardURL.Opaque == "443" { // https访问
		address = forwardURL.Scheme + ":443"
	} else { // http访问
		address = forwardURL.Host
		if strings.Index(forwardURL.Host, ":") == -1 { // 如果host不带端口， 默认80
			address += ":80"
		}
	}
	return address, nil
}
