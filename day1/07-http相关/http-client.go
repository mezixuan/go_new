package main

import "net/http"

func main() {
	//http包
	client := http.Client{}
	client.Get("www.baidu.com")
}
