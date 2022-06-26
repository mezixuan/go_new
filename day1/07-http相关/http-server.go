package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//注册路由
	//xxxx/user ===> func1
	//xxxx/name ===> func2
	//xxxx/id   ===> func3
	//https://127.0.0.1:8080/user, func是回调函数，用于路由的响应，这个回调的函数原型是固定的
	http.HandleFunc("/user", func(writer http.ResponseWriter, request *http.Request) {
		//request : ==>包含客户端发来的数据
		fmt.Println("用户请求详情:")
		fmt.Println("request:", request)
		//writer : ==>通过writer将数据返回给客户端
		_, _ = io.WriteString(writer, "这是/user返回的数据")
	})
	//https://127.0.0.1:8080/name
	http.HandleFunc("/name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/name返回的数据")
	})
	//https://127.0.0.1:8080/id
	http.HandleFunc("/id", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "这是/id返回的数据")
	})
	fmt.Println("http server start.....")

	//简便的写法
	if err := http.ListenAndServe("127.0.0.1:8080", nil); err != nil {
		fmt.Println("http start failed, err:", err)
		return
	}
}
