package handler

import (
	"encoding/json"
	"fserver/constant"
	"fserver/message"
	"fserver/util"
	"net"
	"os"
	"strings"
)

type (
	Controller struct {
		Input
	}

	Input struct {
		RootPath string
		Port     string
	}

	client struct {
		Input

		fileChan  chan string
		closeChan chan interface{}

		conn net.Conn
	}
)

func New(i Input) *Controller {

	return &Controller{
		Input: i,
	}
}

func (c *Controller) Start() {
	server := "0.0.0.0:" + c.Port
	listener, err := net.Listen("tcp", server)
	if err != nil {
		panic(err)
	}

	util.ZapLogger.Infof("listen at %v", server)

	for {
		conn, err := listener.Accept()

		if err != nil {
			util.ZapLogger.Errorf("accept error %v", err)
			continue
		}

		sockClient := client{
			Input:     c.Input,
			fileChan:  make(chan string, 1),
			closeChan: make(chan interface{}, 1),
			conn:      conn,
		}

		sockClient.handleConn()
	}

}

func (c *client) handleConn() {

	util.ZapLogger.Infof("connect from: %v", c.conn.RemoteAddr())

	go c.handlerWrite()
	go c.handlerRead()

}

func (c *client) handlerRead() {
	for {
		sockObj := message.SockObj{}
		err := sockObj.Read(c.conn)
		if err != nil {
			util.ZapLogger.Errorf("read data failed: %v", err)
			c.closeChan <- message.SockObj_CLOSE
			break
		}

		objType := sockObj.GetObjType()

		switch objType {
		case message.SockObj_CLOSE: // close
			c.closeChan <- message.SockObj_CLOSE
			break

		case message.SockObj_STRING: // string message
			val := sockObj.GetStrObj().GetVal()

			// finish command
			if val == constant.OVER {
				c.closeChan <- message.SockObj_CLOSE
				break
			}

			// show dir command
			if strings.HasPrefix(val, constant.LS) {
				path := strings.TrimSpace(strings.Split(val, constant.LS)[1])
				fs := util.GetFileTree(c.RootPath + path)
				data, _ := json.MarshalIndent(fs, "", `    `)
				c.sendStrMessage(string(data))
				break
			}

			// otherwise is file path, preper to send
			util.ZapLogger.Info(val)
			c.fileChan <- val
			break

		case message.SockObj_FILE:
			break
		default:
			break
		}
	}
}

func (c *client) handlerWrite() {
	for {
		select {
		case path := <-c.fileChan:
			c.sendFile(path)
		case <-c.closeChan:
			return
		default:
			break
		}
	}
}

func (c *client) sendFile(path string) {
	util.ZapLogger.Infof("start write file: %v", path)

	file, err := os.OpenFile(c.RootPath+"/"+path, os.O_RDONLY, 0777)
	if err != nil {
		util.ZapLogger.Errorf("open file error: %v", err)
		return
	}

	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		util.ZapLogger.Errorf("get file status: %v", err)
		return
	}

	md5Str := util.GetMD5Str(file)

	bs := make([]byte, 1024)
	var total int64 = 0
	for {
		n, err := file.Read(bs)
		if err != nil {
			break
		}

		total += int64(n)

		fileObj := message.SockObj_FileObj{
			Content: bs[:n],
			Len:     info.Size(),
			Md5:     md5Str,
			Last:    total >= info.Size(),
			Name:    path,
		}
		err = c.sendFileMessage(&fileObj)
		if err != nil {
			util.ZapLogger.Errorf("send file error: %v", err)
			break
		}
	}
}

func (c *client) sendFileMessage(f *message.SockObj_FileObj) error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_FILE,
		FileObj: f,
	}

	return sockObj.Write(c.conn)
}

func (c *client) sendStrMessage(msg string) error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_STRING,
		StrObj:  &message.SockObj_StrObj{System: true, Val: msg},
	}

	return sockObj.Write(c.conn)
}
