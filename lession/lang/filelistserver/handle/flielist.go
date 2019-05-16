package handle

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const presfix = "/list/"

// 自定义异常 实现userError接口
type userError string

func (ue userError) Error() string {
	return ue.Message()
}

func (ue userError) Message() string {
	return string(ue)
}

// HandleList 处理文件
func FileList(w http.ResponseWriter, r *http.Request) error {
	if strings.Index(r.URL.Path, presfix) != 0 {
		return userError("path must start with: " + presfix)
	}
	path := r.URL.Path[len("/path/"):]
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	w.Write(all)
	return nil
}
