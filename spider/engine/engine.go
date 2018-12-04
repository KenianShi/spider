package engine

import (
	"log"
	"shikenian/learn/spider/fetcher"
)

func Run(seeds ...Request) {
	var requests []Request
	for _, v := range seeds {
		requests = append(requests, v)
	}
	for len(requests) > 0 {
		req := requests[0]
		requests = requests[1:]
		contents, err := fetcher.Fetch(req.Url)
		if err != nil {
			log.Println("err:", err)
			continue
		}
		parseresult := req.ParseFunc(contents)
		requests = append(requests, parseresult.Request...)
		for _, item := range parseresult.Items {
			log.Printf("%v", item)
		}

	}
}
