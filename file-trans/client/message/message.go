package message

import (
	"bytes"
	"encoding/binary"
	"fclient/util"
	"fmt"
	"github.com/golang/protobuf/proto"
	"io"
)

// 读取
func (sockObj *SockObj) Read(reader io.Reader) error {
	util.LogDate("开始读取数据")

	sizeBuf := make([]byte, 8)
	_, err := io.ReadFull(reader, sizeBuf)
	if err != nil {
		return err
	}
	size := binary.BigEndian.Uint64(sizeBuf)
	util.LogDate(fmt.Sprintf("需要接收的数据大小: %v", size))

	// 读取需要的大小数据
	bs := make([]byte, size)
	n, err := io.ReadFull(reader, bs)
	if err != nil {
		return err
	}

	// 写入buf
	buf := &bytes.Buffer{}
	_, err = buf.Write(bs[:n])
	if err != nil {
		return err
	}

	// 解码
	err = proto.Unmarshal(buf.Bytes(), sockObj)
	if err != nil {
		return err
	}
	util.LogDate("解码数据完成")

	return nil
}

// WriteIt 写出到流
func (sockObj *SockObj) Write(write io.Writer) error {
	util.LogDate("开始写出数据")

	bs, err := proto.Marshal(sockObj)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(bs)

	// 输出 将要输出的大小
	sizeBuf := make([]byte, 8) // long; uint64
	binary.BigEndian.PutUint64(sizeBuf, uint64(len(bs)))
	_, _ = write.Write(sizeBuf)

	// 输出
	n, err := buf.WriteTo(write)
	if err != nil {
		return err
	}

	util.LogDate("写出完成, 大小: " + fmt.Sprint(n))

	return nil
}