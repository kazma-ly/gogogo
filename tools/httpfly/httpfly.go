package httpfly

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

var (
	client = &http.Client{Timeout: 60 * time.Second}
)

// Get GET请求
func Get(url string) (*http.Response, error) {
	return client.Get(url)
}

// PostKV Post键值对的数据
func PostKV(url string, contentType string, params url.Values) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(params.Encode()))
	return client.Post(url, contentType, body)
}

// Post POST请求
func Post(url, contentType string, body io.Reader) (*http.Response, error) {
	return client.Post(url, contentType, body)
}

// Request 自定义请求
func Request(method, url string, header map[string]string, cookies []*http.Cookie, params url.Values) (*http.Response, error) {
	body := ioutil.NopCloser(strings.NewReader(params.Encode()))
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		req.Header.Add(k, v)
	}
	for _, c := range cookies {
		req.AddCookie(c)
	}
	return RequestByReq(req)
}

// RequestByReq 自定义Request
func RequestByReq(req *http.Request) (*http.Response, error) {
	defer req.Body.Close()
	return client.Do(req)
}

// FileDownload 文件下载
func FileDownload(url string, filePath string) {
	res, err := Get(url)
	if err != nil {
		log.Println("请求失败: ", err.Error())
		return
	}
	defer res.Body.Close()
	f, err := os.Create(filePath)
	if err != nil {
		log.Println("创建文件失败: ", err.Error())
		return
	}
	bs := make([]byte, 1024)
	for {
		n, err := res.Body.Read(bs)
		if err != nil {
			break
		}
		f.Write(bs[:n])
	}
}
