package parser

import (
	"lession/crawler/engine"
	"lession/crawler/model"
	"log"
	"regexp"
)

var urlRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`)
var cityURLRe = regexp.MustCompile(`href="(http://www.zhenai.com/zhenghun/[^"]+)"`)

// ParseProfileList 解析用户列表
func ParseProfileList(contents []byte, _ string) engine.ParseResult {

	profileURLS := extractStrings(contents, urlRe)

	parseResult := engine.ParseResult{}

	for _, profileURL := range profileURLS {
		log.Printf("user url: %s, name: %s", profileURL.URL, profileURL.Name)
		parseResult.Requests = append(parseResult.Requests,
			engine.Request{
				URL:        profileURL.URL,
				ParserFunc: profileParser(profileURL.Name),
			})
	}

	// 当前页的数据
	allLinks := cityURLRe.FindAllSubmatch(contents, -1)
	for _, link := range allLinks {
		parseResult.Requests = append(parseResult.Requests, engine.Request{
			URL:        string(link[1]),
			ParserFunc: ParseProfileList,
		})
	}

	return parseResult
}

func extractStrings(contents []byte, urlRe *regexp.Regexp) []model.ProfileURL {
	match := urlRe.FindAllSubmatch(contents, -1)
	var urls []model.ProfileURL
	for _, m := range match {
		if len(m) >= 3 {
			urls = append(urls, model.ProfileURL{URL: string(m[1]), Name: string(m[2])})
		} else {
			urls = append(urls, model.ProfileURL{URL: string(m[1]), Name: "unknow"})
		}
	}
	return urls
}

/**
 func(bytes []byte) engine.ParseResult { // 包裹一次
	return ParseProfile(bytes, url.URL, name)
 }
*/
func profileParser(name string) engine.ParserFunc {
	return func(bytes []byte, url string) engine.ParseResult {
		return ParseProfile(bytes, url, name)
	}
}
