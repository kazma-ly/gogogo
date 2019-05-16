package main

import (
	"lession/lang/filelistserver/handle"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"
)

type appHandle func(w http.ResponseWriter, r *http.Request) error

func errWrapper(handle appHandle) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		// 处理未知异常
		defer func() {
			r := recover()
			if r != nil {
				log.Println(r)
				http.Error(writer,
					http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError,
				)
			}
		}()

		// 包装了函数 处理异常
		err := handle(writer, request)

		// 对异常判断处理
		if err != nil {
			log.Println(err)
			// 自定义异常
			if userErr, ok := err.(userError); ok {
				http.Error(writer, userErr.Message(), http.StatusInternalServerError)
				return
			}
			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

// 自定义异常接口
type userError interface {
	error
	Message() string
}

func main() {

	http.HandleFunc("/",
		errWrapper(handle.FileList))

	log.Println(http.ListenAndServe(":8888", nil))
}
