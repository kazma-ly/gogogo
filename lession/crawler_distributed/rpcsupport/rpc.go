package rpcsupport

import (
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// ServeRPC 启动一个RPC服务
func ServeRPC(host string, service interface{}) error {
	rpc.Register(service)

	listener, err := net.Listen("tcp", host)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v", err)
			continue
		}

		go jsonrpc.ServeConn(conn)
	}
}

// NewClient 新建一个RPC客户端
func NewClient(host string) (*rpc.Client, error) {
	conn, err := net.Dial("tcp", host)

	if err != nil {
		return nil, err
	}

	return jsonrpc.NewClient(conn), nil
}
