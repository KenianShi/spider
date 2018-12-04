package parser

import (
	"regexp"
	"shikenian/learn/spider/engine"
)

const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	reg := regexp.MustCompile(cityListRe)
	matches := reg.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	limit := 10
	for _, m := range matches {
		result.Items = append(result.Items, "City:"+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseCity,
		})
		limit--
		if limit == 0 {
			break
		}
	}
	return result
}
