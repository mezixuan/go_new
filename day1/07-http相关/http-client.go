package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	//http包
	client := http.Client{}
	//func(c *Client) Get(url, string) (resp *Response, err error)
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("Client.Get error, err:", err)
		return
	}
	//请求体
	body := resp.Body
	fmt.Println("body 111:", body)
	readBodyStr, err := ioutil.ReadAll(body)
	if err != nil {
		fmt.Println("read body err:", err)
		return
	}
	fmt.Println("body string:", string(readBodyStr))

	//获取请求头
	//beego, gin ==> web框架
	ct := resp.Header.Get("Content-Type")
	date := resp.Header.Get("Date")
	server := resp.Header.Get("Server")

	fmt.Println("Content-Type:", ct)
	fmt.Println("Date:", date)
	//BWS是Baidu Web Server，是百度开发的一个web服务器 大部分的百度的web应用程序使用的是BWS
	fmt.Println("Server:", server)

	url := resp.Request.URL
	code := resp.StatusCode
	status := resp.Status

	fmt.Println("URL:", url)
	fmt.Println("code:", code)
	fmt.Println("status:", status)
}
