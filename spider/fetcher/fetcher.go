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

func ClientFetch(url string)([]byte,error){
	client := &http.Client{}
	req,err := http.NewRequest("GET",url,nil)
	if err != nil {
		return nil, fmt.Errorf("error creating an request %v ",err)
	}
	req.Header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.110 Safari/537.36")
	resp,err := client.Do(req)
	if err != nil {
		return nil,fmt.Errorf("error do request:",err)
	}
	defer resp.Body.Close()
	if http.StatusOK != resp.StatusCode{
		return nil, fmt.Errorf("error resp status code: %v ",resp.StatusCode)
	}
	return ioutil.ReadAll(resp.Body)
}

