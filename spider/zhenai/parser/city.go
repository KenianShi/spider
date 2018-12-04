package parser

import (
	"shikenian/learn/spider/engine"
	"regexp"
)

const cityRege = `<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`

func ParseCity(content []byte) engine.ParseResult{
	//fmt.Print("网页")
	//fmt.Println(string(content))
	reg := regexp.MustCompile(cityRege)
	matches := reg.FindAllSubmatch(content, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, "User:"+string(m[2]))
		result.Request = append(result.Request, engine.Request{
			Url:       string(m[1]),
			ParseFunc: ParseProfile,
		})
	}
	return result
}