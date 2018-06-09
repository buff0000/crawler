package parse

import (
	"com.buff/learngo/crawler/engine"
	"regexp"
)

const cityRe = `<th><a href="(http://album.zhenai.com/u/[0-9]+)"[^>]+">([^<]+)</a></th>`
func ParseCity(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches{
		name := string(m[2])
		result.Items = append(result.Items, "User " + string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: string(m[1]),
			ParserFunc: func(c []byte) engine.ParseResult{
				return ParseProfile(c, name)
			},
		})
		//fmt.Printf("User: %s, URL: %s", m[2], m[1])
	}
	return result
}
