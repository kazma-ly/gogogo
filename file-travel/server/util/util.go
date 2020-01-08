package util

import (
	"crypto/md5"
	"encoding/hex"
	"fserver/entity"
	"io/ioutil"
	"os"

	"go.uber.org/zap"
)

var (
	ZapLogger *zap.SugaredLogger
)

func init() {
	zapLogger, err := zap.NewProduction()

	if err != nil {
		panic("log init failed")
	}

	ZapLogger = zapLogger.Sugar()
}

func GetMD5Str(file *os.File) string {
	defer file.Seek(0, 0)

	hash := md5.New()
	bs := make([]byte, 1024)

	for {
		n, err := file.Read(bs)
		if err != nil {
			break
		}
		hash.Write(bs[:n])
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func GetFileTree(path string) []entity.FileInfo {
	fileInfoList, err := ioutil.ReadDir(path)

	var fis = []entity.FileInfo{}
	if err != nil {
		ZapLogger.Errorf("read dir error %v", err)
		return fis
	}

	for _, fileInfo := range fileInfoList {
		fi := entity.FileInfo{
			Name:    fileInfo.Name(),
			Size:    fileInfo.Size(),
			IsDir:   fileInfo.IsDir(),
			ModTime: fileInfo.ModTime(),
		}

		fis = append(fis, fi)
	}

	return fis
}
