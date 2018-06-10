package parse

import (
	"regexp"
	"strconv"
	"com.buff/Crawler/crawler/crawler/model"
	"com.buff/Crawler/crawler/crawler/engine"
)


var ageRe = regexp.MustCompile(`<td><span class="label">年龄：</span>(\d+)岁</td>`)
var heightRe = regexp.MustCompile(`<td><span class="label">身高：</span>(\d+)CM</td>`)
var incomeRe = regexp.MustCompile(`<td><span class="label">月收入：</span>([^<]+)</td>`)
var marrigedRe = regexp.MustCompile(`<td><span class="label">婚况：</span>([^<]+)</td>`)
var addrRe = regexp.MustCompile(`<td><span class="label">工作地：</span>([^<]+)</td>`)
var workRe = regexp.MustCompile(`<td><span class="label">职业： </span>([^<]+)</td>`)
var hashChirdRe = regexp.MustCompile(`<td><span class="label">有无孩子：</span>([^<]+)</td>`)


func ParseProfile(contents []byte, name string) engine.ParseResult{
	profile := model.Profile{}
	profile.Name = name
	age, err := strconv.Atoi(extractString(contents, ageRe))
	if err != nil {
		profile.Age = age
	}
	height, err := strconv.Atoi(extractString(contents, heightRe))
	if err != nil {
		profile.Heigth = height
	}
	//profile.Age = extractString(contents, ageRe)
	//profile.Heigth = extractString(contents, heightRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Marriged = extractString(contents, marrigedRe)
	profile.Addr = extractString(contents, addrRe)
	profile.Work = extractString(contents, workRe)
	profile.HashChird = extractString(contents, hashChirdRe)

	result := engine.ParseResult{
		Items: []interface{}{profile},
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) string{
	matchs := re.FindSubmatch(contents)
	if len(matchs) >= 2 {
		return string(matchs[1])
	}else{
		return ""
	}
}