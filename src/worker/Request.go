package worker

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func ReadContent(url string) (content string, err error){
	resp, err1 := http.Get(url)
	if err1 != nil {
		fmt.Println("出错了。。 url",url, err1)
	}
	defer resp.Body.Close()
	data, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("出错了。。 url", err2)
		return "", err2
	}
	body := string(data)
	return body, nil
}