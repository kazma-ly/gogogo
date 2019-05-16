package parser

import (
	"lession/crawler/engine"
	"regexp"
)

// ^> 非右括号
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[a-z0-9]+)"[^>]*>([^<]+)</a>`

// ParseCityList 城市用户列表
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	all := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}

	for _, a := range all {
		// result.Items = append(result.Items, a[2])
		result.Requests = append(result.Requests, engine.Request{
			URL:        string(a[1]),
			ParserFunc: ParseProfileList,
		})
	}
	return result
}
