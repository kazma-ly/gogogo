package violet

import (
	"regexp"
	"strings"
)

type (
	// Router 一个路由
	Router struct {
		URL    string         // 原始URL
		regex  *regexp.Regexp // 正则匹配器
		Method string         // 请求方法
		Con    Controller     // 控制器
	}
)

func (r *Router) match(url string) bool {
	return r.regex.ReplaceAllString(url, "") == ""
}

// parsePathVar 解析url的变量
func (r *Router) parsePathVar(url string) map[string]string {
	// 字符串末尾是否有"/"
	if strings.HasSuffix(url, "/") {
		// 截取最后一个/前面得字符串
		url = strings.TrimSuffix(url, "/")
	}
	pathInfo := strings.Split(url, "/")

	origin := r.URL
	if strings.HasSuffix(origin, "/") {
		origin = strings.TrimSuffix(origin, "/")
	}

	origins := strings.Split(origin, "/")

	pathVar := make(map[string]string)
	for i, o := range origins {
		if strings.HasPrefix(o, ":") {
			pathVar[o[1:]] = pathInfo[i]
		}
	}
	return pathVar
}
