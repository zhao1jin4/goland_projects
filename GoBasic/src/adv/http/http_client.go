package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func printRes(resp *http.Response, err error) {
	if err != nil {
		fmt.Print("error:", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Printf("body=%s \n", body)
}
func main() {
	{
		resp, err := http.Get("http://localhost:8080/J_JavaEE/index.jsp")
		printRes(resp, err)
		resp, err = http.PostForm("http://localhost:8080/J_JavaEE/receiveForm",
			url.Values{"username": {"lisi"}, "password": {"123"}})
		printRes(resp, err)
	}
	{
		client := &http.Client{
			//CheckRedirect: redirectPolicyFunc,
		}
		resp, err := client.Get("http://localhost:8080/J_JavaEE/index.jsp")
		printRes(resp, err)

		req, err := http.NewRequest("POST", "http://localhost:8080/J_JavaEE/receiveForm", nil)
		req.Header.Add("If-None-Match", `W/"wyzzy"`) //请求头
		resp, err = client.Do(req)                   //发起请求
		printRes(resp, err)
	}
	//https

	//客户端文件 上传/下载 未做
}
