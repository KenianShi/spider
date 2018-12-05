package parser

import (
	"testing"
	"io/ioutil"
	"fmt"
)

func TestParseProfile(t *testing.T) {
	//首次获取页面下来以备以后的测试
	//content,err := fetcher.ClientFetch("http://album.zhenai.com/u/1046180939")
	//if err != nil {
	//	t.Errorf("err:",err)
	//}
	//err = ioutil.WriteFile("profile.html",content,0644)
	//if err != nil {
	//	t.Errorf("create file err:",err)
	//}
	content,err := ioutil.ReadFile("./profile.html")
	if err != nil {
		t.Errorf("error reading file ",err)
	}
	result := ParseProfile(content)
	fmt.Print("result")
	fmt.Println(result.Items)

}
