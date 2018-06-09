package parse

import (
	"com.buff/learngo/crawler/engine"
	"regexp"
	"fmt"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`
func ParseCityList(contents []byte) engine.ParseResult{
	re := regexp.MustCompile(cityListRe)
	matchs := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	index := 3
	for _, m := range matchs{

		result.Items = append(result.Items, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url: 		string(m[1]),
			ParserFunc: ParseCity,
		})
	//	fmt.Printf("City: %s, URL: %s", m[2], m[1])
		if index < 0 {
			break
		}
		index --
	}
	fmt.Printf("Match found: %d", len(matchs))
	return result
}
