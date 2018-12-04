package fetcher

import (
	"testing"
	"log"
)

func TestFetch(t *testing.T) {
	//contents, err := fetcher.Fetch(req.Url)
	content,err := Fetch("http://album.zhenai.com/u/107705060")
	if err != nil {
		log.Println("err: ",err)
	}
	log.Println(string(content))
}