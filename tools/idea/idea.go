package main

import (
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type handle struct {
	reverseProxy string
}

func (h *handle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	remote, err := url.Parse(h.reverseProxy) // 要转向的url
	if err != nil {
		log.Fatalln(err)
	}
	proxy := httputil.NewSingleHostReverseProxy(remote) // 转向

	// 把请求的地址从本地转向代理地址
	r.Host = remote.Host
	proxy.ServeHTTP(w, r)
}

func main() {

	bind := flag.String("h", "0.0.0.0", "listen on ip")
	port := flag.String("p", "8888", "listen on port")
	remote := flag.String("r", "http://idea.iteblog.com/key.php", "reverse proxy addr") // http://idea.imsxm.com:80
	flag.Parse()
	addr := *bind + ":" + *port

	log.Printf("Listening on %s:%s, forwarding to %s", *bind, *port, *remote)
	log.Println("在idea中填入该地址: ", "http://"+addr)
	h := &handle{reverseProxy: *remote}

	log.Fatalln("服务启动失败: ", http.ListenAndServe(addr, h))
}
