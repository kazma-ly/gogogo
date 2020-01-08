package message

import (
	"bytes"
	"encoding/binary"
	"io"

	"github.com/golang/protobuf/proto"
)

func (sockObj *SockObj) Read(reader io.Reader) error {
	sizeBuf := make([]byte, 8)
	_, err := io.ReadFull(reader, sizeBuf)
	if err != nil {
		return err
	}

	size := binary.BigEndian.Uint64(sizeBuf)

	bs := make([]byte, size)
	n, err := io.ReadFull(reader, bs)
	if err != nil {
		return err
	}

	buf := &bytes.Buffer{}
	_, err = buf.Write(bs[:n])
	if err != nil {
		return err
	}

	err = proto.Unmarshal(buf.Bytes(), sockObj)
	if err != nil {
		return err
	}

	return nil
}

func (sockObj *SockObj) Write(write io.Writer) error {
	bs, err := proto.Marshal(sockObj)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(bs)

	sizeBuf := make([]byte, 8) // long; uint64
	binary.BigEndian.PutUint64(sizeBuf, uint64(len(bs)))
	_, _ = write.Write(sizeBuf)

	_, err = buf.WriteTo(write)
	if err != nil {
		return err
	}

	return nil
}
