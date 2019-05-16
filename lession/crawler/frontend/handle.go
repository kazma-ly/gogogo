package frontend

import (
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

var (
	bs   []byte
	lock sync.Once
)

func RenderPage(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Add("Content-Type", "text/html; charset=UTF-8")
	resp.Write(getFile())
}

func getFile() []byte {
	lock.Do(func() {
		_f, err := os.OpenFile("crawler/frontend/view/index.html", os.O_RDONLY, 0755)
		if err != nil {
			panic(err)
		}
		_bs, err := ioutil.ReadAll(_f)
		defer _f.Close()
		if err != nil {
			panic(err)
		}
		bs = _bs
	})
	return bs
}
