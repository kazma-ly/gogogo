package violet

import (
	"encoding/json"
	"io"
	"life-service/httpfly/httpres"
	"life-service/session"
	"log"
	"net/http"
	"os"
)

type (
	// Context 上下文
	Context struct {
		Request       *http.Request       // 请求
		ResponseWrite http.ResponseWriter // 响应
		pathVar       map[string]string   // url参数
	}
)

func (c *Context) init(req *http.Request, respw http.ResponseWriter) {
	c.Request = req
	c.ResponseWrite = respw
}

// WriteData 输出数据
func (c *Context) WriteData(bs []byte, contentType string) {
	c.Request.Header.Set("Content-Type", contentType)
	c.ResponseWrite.Write(bs)
}

// WriteFile 写出文件
func (c *Context) WriteFile(name string, contentType string) {
	c.Request.Header.Set("Content-Type", contentType)
	file, err := os.OpenFile(name, os.O_RDONLY, 0755)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer file.Close()
	io.Copy(c.ResponseWrite, file)
}

// WriteJSON 写出Json数据
func (c *Context) WriteJSON(obj interface{}) {
	c.ResponseWrite.Header().Set("Content-Type", "application/json")
	bs, err := json.Marshal(obj)
	if err != nil { // JSON ERROR
		c.ResponseWrite.Write(httpres.Create(http.StatusInternalServerError, err.Error(), nil, false).Bytes())
		return
	}
	c.ResponseWrite.Write(bs)
}

// Session 获得session
func (c *Context) Session() *session.Manager {
	panic("It Drop")
	// return session.GetSessionManage(c.Request, c.ResponseWrite)
}

// AddCookie 添加cookie
func (c *Context) AddCookie(cookie *http.Cookie) {
	http.SetCookie(c.ResponseWrite, cookie)
}

// GetPathVar 获得url上的变量
func (c *Context) GetPathVar(name string) string {
	return c.pathVar[name]
}
