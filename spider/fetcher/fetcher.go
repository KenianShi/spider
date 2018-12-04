package fetcher

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"log"
)

func Fetch(url string)([]byte,error){
	log.Println("Got URL:",url)
	resp,err := http.Get(url)
	if err != nil {
		return nil,err
	}
	defer resp.Body.Close()
	if http.StatusOK != resp.StatusCode {
		return nil, fmt.Errorf("error resp status code: %v ",resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}