package handler

import (
	"fmt"
	"fserver/message"
	"fserver/util"
	"net"
	"os"
)

type (
	Controller struct {
		fileChan  chan string
		closeChan chan interface{}
	}
)

var (
	rootPath string
)

// HandleConn 处理连接的客户端 63938 63942
func HandleConn(conn net.Conn, rPath string) {
	rootPath = rPath

	c := &Controller{
		fileChan:  make(chan string, 10),
		closeChan: make(chan interface{}),
	}

	util.LogDate(fmt.Sprintf("connect from: %v", conn.RemoteAddr()))

	go handlerWrite(conn, c)
	go handlerRead(conn, c)
}

// handlerRead 处理socket的读操作
func handlerRead(conn net.Conn, c *Controller) {
	for {
		sockObj := message.SockObj{}
		err := sockObj.Read(conn)
		if err != nil {
			util.LogDate(fmt.Sprintf("读取消息失败: %v", err))
			c.closeChan <- message.SockObj_CLOSE
			break
		}

		objType := sockObj.GetObjType()

		switch objType {
		case message.SockObj_CLOSE: // 结束了
			c.closeChan <- message.SockObj_CLOSE
			return
		case message.SockObj_STRING: // 字符串消息, 直接打印出来
			val := sockObj.GetStrObj().GetVal()
			if val == "over" {
				c.closeChan <- message.SockObj_CLOSE
			} else {
				util.LogDate(val)
				c.fileChan <- rootPath + val
			}
			break
		case message.SockObj_FILE: // 文件消息
			// 忽略
			break
		default:
			break
		}
	}
}

// handlerWrite 处理socket的写操作
func handlerWrite(conn net.Conn, c *Controller) {
	for {
		select {
		case path := <-c.fileChan:
			err := processSendFile(path, conn)
			if err != nil {
				util.LogDate(fmt.Sprintf("处理写文件发生错误: %v\n", err))
				return
			}
		case <-c.closeChan:
			return
		default:
			break
		}
	}
}

func processSendFile(path string, conn net.Conn) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0777)
	if err != nil {
		return err
	}

	// 文件信息
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// 文件md5
	md5Str, err := util.GetMD5Str(path)
	if err != nil {
		return err
	}

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
			Name:    file.Name(),
		}
		err = sendFileMessage(conn, &fileObj)
		if err != nil {
			util.LogDate(fmt.Sprintf("发送文件失败: %v\n", err))
		}
	}
	return nil
}

func sendFileMessage(conn net.Conn, f *message.SockObj_FileObj) error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_FILE,
		FileObj: f,
	}
	return sockObj.Write(conn)
}

// sendStrMessage 发送socket的文字消息
func sendStrMessage(msg string, conn net.Conn) error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_STRING,
		StrObj:  &message.SockObj_StrObj{System: true, Val: msg},
	}
	return sockObj.Write(conn)
}
