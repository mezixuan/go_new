package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	ip := "127.0.0.1"
	port := 8848
	address := fmt.Sprintf("%s:%d", ip, port)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("net listen err:", err)
		return
	}

	fmt.Println("监听中")
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err:", err)
	}

	fmt.Println("连接建立成功!")

	//创建一个容器，用于接收读取到的数据
	buf := make([]byte, 1024) //使用make来创建切片

	//cnt:真正读取client发来的数据的长度
	cnt, err := conn.Read(buf)
	if err != nil {
		fmt.Println("conn.Read err:", err)
		return
	}
	fmt.Println("client===>server,长度:", cnt, "数据:", string(buf))

	//将数据转成大写“hello”==>HELLO
	upperdata := strings.ToUpper(string(buf))
	cnt, err = conn.Write([]byte(upperdata))
	fmt.Println("Client <=======  Server, 长度是:", cnt, "内容是", upperdata)
	conn.Close()
}
