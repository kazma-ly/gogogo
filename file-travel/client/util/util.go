package util

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"os"
	"path/filepath"

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

func GetMD5Str(reader io.Reader) string {
	hash := md5.New()

	bs := make([]byte, 1024)
	for {
		n, err := reader.Read(bs)
		if err != nil {
			break
		}

		hash.Write(bs[:n])
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func FilePathExits(filePath string) bool {
	path := filepath.Dir(filePath)
	_, err := os.Stat(path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}

func CreateDirIfNotExits(filePath string) error {
	if FilePathExits(filePath) {
		return nil
	}

	path := filepath.Dir(filePath)
	return os.Mkdir(path, os.ModePerm)
}

func FileExist(_path string) bool {
	_, err := os.Stat(_path)
	if err != nil && os.IsNotExist(err) {
		return false
	}

	return true
}
