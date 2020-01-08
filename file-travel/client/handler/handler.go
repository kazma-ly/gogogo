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
		Input

		MsgCh   chan string // msg
		CloseCh chan string // close msg

		conn net.Conn

		Temp map[string]*os.File
	}

	Input struct {
		DownloadPath string // download local path
		Server       string // the file server
	}
)

func New(i Input) *Controller {
	return &Controller{
		Input:   i,
		MsgCh:   make(chan string, 1),
		CloseCh: make(chan string, 1),
		Temp:    make(map[string]*os.File),
	}
}

func (c *Controller) Handler() {
	conn, err := net.Dial("tcp", c.Server)
	if err != nil {
		panic(err) // throw it, server must be alive
	}

	util.ZapLogger.Infof("receive from: %v, save to: %v", c.Server, c.DownloadPath)

	c.conn = conn

	go c.handlerRead()
	go c.handlerWrite()
}

func (c *Controller) handlerRead() {
	for {
		sockObj := message.SockObj{}
		err := sockObj.Read(c.conn)
		if err != nil {
			util.ZapLogger.Errorf("read data failed: %v", err)
			break
		}

		objType := sockObj.GetObjType()

		switch objType {
		case message.SockObj_CLOSE: // done
			return
		case message.SockObj_STRING: // string msg
			fmt.Println(sockObj.GetStrObj().GetVal())
			break
		case message.SockObj_FILE: // file msg
			err := c.processFileMessage(c.conn.RemoteAddr().String(), sockObj.GetFileObj())
			if err != nil {
				util.ZapLogger.Errorf("read file data failed: %v", err)
			}
			break
		default:
			break
		}
	}
}

func (c *Controller) processFileMessage(addr string, obj *message.SockObj_FileObj) error {
	if obj == nil {
		util.ZapLogger.Error("file obj not exits")
		return errors.New("file obj not exits")
	}

	name := obj.GetName()
	key := addr + "_" + name
	fpath := c.DownloadPath + name

	err := util.CreateDirIfNotExits(fpath)
	if err != nil {
		util.ZapLogger.Errorf("创建文件出错: %v", err)
		return err
	}

	if c.Temp[key] == nil {
		if util.FileExist(fpath) {
			if os.Remove(fpath) != nil {
				util.ZapLogger.Errorf("移除文件出错: %v", err)
				return err
			}
		}

		file, err := os.OpenFile(fpath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0777)
		if err != nil {
			util.ZapLogger.Errorf("打开文件出错: %v", err)
			return err
		}

		c.Temp[key] = file
	}

	_file := c.Temp[key]

	_, err = _file.Write(obj.Content)

	if err != nil {
		util.ZapLogger.Errorf("写出文件出错: %v", err)
		_file.Close()
		return err
	}

	fileInfo, _ := _file.Stat()
	fmt.Printf("the %s file progress: %.2f%%\r", fileInfo.Name(), float64(fileInfo.Size())/float64(obj.Len)*100.0)

	if obj.Last {
		fmt.Println()
		defer _file.Close()
		delete(c.Temp, key)

		_file.Seek(0, 0)
		md5String := util.GetMD5Str(_file)
		if md5String != obj.Md5 {
			util.ZapLogger.Errorf("%v: md5 check error remote: %v, local %v", obj.Name, obj.Md5, md5String)
		}

	}

	return nil
}

func (c *Controller) handlerWrite() {
	for {
		select {
		case val := <-c.MsgCh:
			_ = c.sendStrMessage(val)
		case <-c.CloseCh:
			_ = c.sendCloseMessage()
			return
		default:
			break
		}
	}
}

func (c *Controller) sendCloseMessage() error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_CLOSE,
	}
	return sockObj.Write(c.conn)
}

func (c *Controller) sendStrMessage(msg string) error {
	sockObj := message.SockObj{
		ObjType: message.SockObj_STRING,
		StrObj:  &message.SockObj_StrObj{System: true, Val: msg},
	}

	return sockObj.Write(c.conn)
}
