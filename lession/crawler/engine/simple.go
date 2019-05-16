package engine

import (
	"log"
)

type SimpleEngine struct{}

func (s SimpleEngine) Run(seeds ...Request) {
	var requests []Request
	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		parserResult, err := worker(r)
		if err != nil {
			log.Printf("fetcher: err fecthing url %s: %v", r.URL, err)
			continue
		}

		requests = append(requests, parserResult.Requests...)
		for _, item := range parserResult.Items {
			log.Printf("got item: %s\n", item)
		}
	}
}
