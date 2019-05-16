package engine

import (
	"lession/crawler/fetcher"
	"log"
)

// 耗时较长 并发执行
func worker(r Request) (ParseResult, error) {
	log.Printf("fetching %s \n", r.URL)

	body, err := fetcher.Fetch(r.URL)
	if err != nil {
		return ParseResult{}, err
	}

	return r.ParserFunc(body, r.URL), nil
}
