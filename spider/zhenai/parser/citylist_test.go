package parser

import (
	"testing"
	"io/ioutil"
	"log"
)

func TestParseCityList(t *testing.T) {
	//第一次将珍爱网的网页拉取下来，保存为文件为以后测试需要
	//content,err := fetcher.Fetch("http://www.zhenai.com/zhenghun")
	//if err != nil {
	//	log.Println(content)
	//	return
	//}
	//err = ioutil.WriteFile("./zhenai.html",content,0644)
	//if err != nil {
	//	log.Println(err)
	//}
	content,err := ioutil.ReadFile("./zhenai.html")
	if err != nil {
		log.Println("err:",err)
		return
	}
	result := ParseCityList(content)
	const resultCount = 470
	expectedItems := []string{"City:阿坝","City:阿克苏","City:阿拉善盟"}
	expectedRequest := []string{"http://www.zhenai.com/zhenghun/aba","http://www.zhenai.com/zhenghun/akesu","http://www.zhenai.com/zhenghun/alashanmeng"}
	if resultCount != len(result.Items){
		t.Errorf("result should have %d request,but got %d,",resultCount,len(result.Items))
	}
	for k,v := range expectedItems{
		if v != result.Items[k]{
			t.Errorf("result.Items[%d] should have %v, but got %v ",k,v,result.Items[k])
		}
	}
	for k,v := range expectedRequest {
		if v != result.Request[k].Url{
			t.Errorf("result.Request[%d].Url should be %v, but got %v",k,v,result.Request[k].Url)
		}
	}
}
