package parser

import (
	"lession/crawler/engine"
	"lession/crawler/model"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>([\d]+)岁</td>`)
var marryRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var eductionRe = regexp.MustCompile(`<td><span class="label">学历：</span>([^<]+)</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>([0-9]+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([0-9-]+)元</td>`)
var occupationRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var jobRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var xinzuoRe = regexp.MustCompile(`<td><span class="label">星座：</span>([^<]+)</td>`)
var weightRe = regexp.MustCompile(`<td><span class="label">体重：</span><span field="">([0-9]+)KG</span></td>`)
var genderRe = regexp.MustCompile(`<td><span class="label">性别：</span><span field="">([^<]+)</span></td>`)
var idRe = regexp.MustCompile(`http://album.zhenai.com/u/([\d]+)`)

//var nameRe = regexp.MustCompile(`<a class="name fs24">([^<]+)</a>`)

// ParseProfile 用户信息详情
func ParseProfile(contents []byte, url string, name string) engine.ParseResult {
	profile := model.Profile{}

	// 年龄
	ageStr := extractString(contents, ageRe)
	age, err := strconv.Atoi(ageStr)
	if err == nil {
		profile.Age = age
	}

	// 婚姻状况
	profile.Marriage = extractString(contents, marryRe)

	// 教育情况
	profile.Education = extractString(contents, eductionRe)

	// 身高
	profile.Height = extractInt(contents, heightRe)

	// 收入
	incomeStr := extractString(contents, incomeRe)
	incomeSplit := strings.Split(incomeStr, "-")
	profile.IncomeLow, _ = strconv.Atoi(incomeSplit[0])
	if len(incomeSplit) >= 2 {
		profile.IncomeUp, _ = strconv.Atoi(incomeSplit[1])
	}

	// 工作地
	profile.Occupation = extractString(contents, occupationRe)

	// 工作title
	profile.Job = extractString(contents, jobRe)

	// 星座
	profile.Xinzuo = extractString(contents, xinzuoRe)

	// 体重
	profile.Weight = extractInt(contents, weightRe)

	// 性别
	profile.Gender = extractString(contents, genderRe)

	// 名字
	profile.Name = name //extractString(contents, nameRe)

	id := extractString([]byte(url), idRe)

	//log.Println(profile)
	result := engine.ParseResult{
		Items: []engine.Item{
			{Url: url, Type: "zhenai", Id: id, Payload: profile},
		},
	}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}

func extractInt(contents []byte, re *regexp.Regexp) int {
	str := extractString(contents, re)
	i, err := strconv.Atoi(str)
	if err == nil {
		return i
	}
	log.Printf("extract int fail: %s", err.Error())
	return 0
}
