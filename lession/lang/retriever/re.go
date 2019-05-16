package main

import (
	"fmt"
	"lession/retriever/mock"
	"lession/retriever/real"
	"time"
)

type (
	Retriever interface {
		Get(url string) string
	}
	Poster interface {
		Post(url string, form map[string]string) string
	}
	RetrieverPoster interface {
		Retriever
		Poster
	}
)

const url = "http://www.imooc.com"

func download(r Retriever) string {
	return r.Get("https://www.baidu.com")
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "zly",
		"course": "golang",
	})
}

func session(s RetrieverPoster) string {
	// s.Get()
	str := s.Post(url, map[string]string{
		"contents": "fack imooc.com",
	})
	return str
}

func main() {
	var r Retriever
	r = mock.Retriever{"good"}
	fmt.Printf("%T %v\n", r, r)
	r = &real.Retriever{"Mozilla/5.0", 30 * time.Second}
	fmt.Printf("%T %v\n", r, r)
	fmt.Println(download(r))

	rp := mock.Retriever{}
	fmt.Println(session(rp))
}
