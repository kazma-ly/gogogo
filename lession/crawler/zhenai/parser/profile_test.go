package parser

import (
	"fmt"
	"lession/crawler/fetcher"
	"testing"
)

func TestParseProfile(t *testing.T) {
	body, err := fetcher.Fetch("http://album.zhenai.com/u/110409917")
	if err != nil {
		panic(err)
	}

	parseResult := ParseProfile(body, "http://album.zhenai.com/u/110409917", "110409917")
	fmt.Printf("%+v", parseResult)

}
