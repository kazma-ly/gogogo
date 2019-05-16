package violet

import (
	"life-service/httpfly/httpres"
	"life-service/logx"
	"net/http"
	"strings"
	"sync"
)

type (
	// 实现路由，系统的一个 实现 ServeHTTP(ResponseWriter, *Request) 接口
	handler struct {
		p sync.Pool // 临时对象池
	}
)

// 处理http请求
func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// Get 会自动执行p.New()方法 初始化得到一个Context
	c := h.p.Get().(*Context)
	defer h.p.Put(c)

	if h.makeStatic(w, r, currentRouterX.static) {
		return
	}

	c.init(r, w)

	path := r.URL.Path
	for _, rou := range currentRouterX.Routers {
		if rou.match(path) {
			logx.LogInfo(rou.Method, path)
			if rou.Method != r.Method {
				w.WriteHeader(http.StatusMethodNotAllowed)
				w.Write(httpres.Create(http.StatusBadRequest, "请求方法错误", nil, false).Bytes())
			} else {
				c.pathVar = rou.parsePathVar(path)
				rou.Con(c)
			}
			return
		}
	}
	http.NotFound(w, r)
}

// 处理用户的静态文件
func (h *handler) makeStatic(w http.ResponseWriter, r *http.Request, staticDir map[string]string) bool {
	for prefix, path := range staticDir {
		if strings.HasPrefix(r.URL.Path, prefix) {
			file := path + r.URL.Path[len(prefix):]
			http.ServeFile(w, r, file)
			return true
		}
	}
	return false
}
