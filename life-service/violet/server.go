package violet

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type (
	// RouterX 路由中心
	RouterX struct {
		Routers map[string]Router
		static  map[string]string
	}

	// Controller 控制器
	Controller func(c *Context)
)

var (
	currentRouterX *RouterX
)

// Default 默认处理器
func Default() *RouterX {
	currentRouterX = &RouterX{
		Routers: make(map[string]Router),
		static:  make(map[string]string),
	}
	return currentRouterX
}

// GET handle GET
func (rx *RouterX) GET(pattern string, c Controller) {
	if !strings.HasPrefix(pattern, "/") {
		panic("url must has '/' prefix")
	}
	r := Router{URL: pattern, Method: http.MethodGet, Con: c, regex: regexp.MustCompile(makeRegex(pattern))}
	rx.Routers[pattern] = r
}

// POST handle POST
func (rx *RouterX) POST(pattern string, c Controller) {
	if !strings.HasPrefix(pattern, "/") {
		panic("url must has '/' prefix")
	}
	r := Router{URL: pattern, Method: http.MethodPost, Con: c, regex: regexp.MustCompile(makeRegex(pattern))}
	rx.Routers[pattern] = r
}

// Start 启动服务器
func (rx *RouterX) Start(addr string) {

	// USE JWT REPLACE
	// session.ProcessSession()

	server := http.Server{
		Handler: newHandle(),
		Addr:    addr,
	}

	log.Fatal(server.ListenAndServe())
}

// MakeStatic 静态文件资源处理
func (rx *RouterX) MakeStatic(urlpath, filepath string) {
	rx.static[urlpath] = filepath
}

func newHandle() *handler {
	h := &handler{}
	h.p.New = func() interface{} {
		return &Context{}
	}
	return h
}

func makeRegex(url string) string {
	urls := strings.Split(url, "/")
	var regex string
	for _, u := range urls {
		if strings.HasPrefix(u, ":") {
			regex += ".+/"
		} else {
			regex += u + "/"
		}
	}
	return regex[:len(regex)-1]
}
