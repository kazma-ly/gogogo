package main

import (
	"net/http"
	"onedrivetool/handler"

	"github.com/julienschmidt/httprouter"
)

// RouterX 路由扩展，增加middleware
type RouterX struct {
	router *httprouter.Router
}

func (rx RouterX) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	rx.router.ServeHTTP(w, r)
}

// NewHandle new
func NewHandle(r *httprouter.Router) http.Handler {
	rx := RouterX{}
	rx.router = r
	return rx
}

// NewRouter 初始化路由
func NewRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/ms/auth", handler.GoMSAuth)
	router.GET("/ms/callback", handler.MSCallBack)
	router.GET("/ms/file", handler.GetOneDriveFile)
	router.GET("/ms/upload", handler.UploadDriveFile)

	return router
}

func main() {
	r := NewRouter()
	h := NewHandle(r)

	http.ListenAndServe(":9001", h)
}
