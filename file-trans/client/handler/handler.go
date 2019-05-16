package handler

import (
	"errors"
	"fclient/message"
	"fclient/util"
	"fmt"
	"net"
	"os"
)

type (
	Controller struct {
		StrCh   chan string // 消息
		CloseCh chan string // 结束
	}
)

var (
	base string
	// fileMap 文件时分片传输的
	fileMap map[string]*os.File
)

func Handler(conn net.Conn, c *Controller, rootPath string) {

	base = rootPath
	fileMap = make(map[string]*os.File)

	go HandlerRead(conn, c)
	go HandlerWrite(conn, c)

}

// handlerRead 处理socket的读操作
func HandlerRead(conn net.Conn, c *Controller) {
	for {
		sockObj := message.SockObj{}
		err := sockObj.Read(conn)
		if err != nil {
			util.LogDate(fmt.Sprintf("读取消息失败: %v", err))
			break
		}

		objType := sockObj.GetObjType()

		switch objType {
		case message.SockObj_CLOSE: // 结束了
			return
		case message.SockObj_STRING: // 字符串消息, 直接打印出来
			util.LogDate(sockObj.GetStrObj().GetVal())
			break
		case message.SockObj_FILE: // 文件消息
			err := processFileMessage(conn.RemoteAddr().String(), sockObj.GetFileObj())
			if err != nil {
				util.LogDate(fmt.Sprintf("读取文件消息出错: %v", err))
				return
			}
			break
		default:
			break
		}
	}
}

// processFileMessage 处理socket的文件消息
func processFileMessage(addr string, obj *message.SockObj_FileObj) error {
	if obj == nil {
		util.LogDate("文件对象为空")
		return errors.New("file obj not exits")
	}
	name := obj.GetName()
	key := addr + "_" + name
	fpath := base + name
	file, err := os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		return err
	}
	defer file.Close()
	fileMap[key] = file

	fileInfo, _ := file.Stat()

	util.LogDate(fmt.Sprintf("%v 进度: %v %", fileInfo.Name(), (float64(fileInfo.Size())/float64(obj.Len)*100.0)))

	_, err = file.Write(obj.Content)

	if obj.Last {
		delete(fileMap, key)
		if checkFile(fpath, obj.Md5) != nil {
			util.LogDate(fmt.Sprintf("文件传输完成，但是文件校验失败，请注意: %v", obj.Name))
		}
		util.LogDate(fmt.Sprintf("%v 进度: 100%", fileInfo.Name()))
		return nil
	}

	return err
}

// checkFile 检查传输文件的正确性
func checkFile(path string, md5Str string) error {
	_md5Str, err := util.GetMD5Str(path)
	if err != nil {
		return err
	}

	if md5Str != _md5Str {
		return errors.New("md5校验失败")
	}

	return nil
}

// handlerWrite 处理socket的写操作
func HandlerWrite(conn net.Conn, c *Controller) {
	for {
		select {
		case val := <-c.StrCh:
			_ = sendStrMessage(val, conn)
		case <-c.CloseCh:
			_ = sendCloseMessage(conn)
			return
		default:
			break
		}
	}
}

func sendCloseMessage(conn net.Conn) error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_CLOSE,
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
