package main

import (
	"shikenian/learn/spider/engine"
	"shikenian/learn/spider/zhenai/parser"
)

const url = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(engine.Request{
		Url:url,
		ParseFunc:parser.ParseCityList,
	})
}


