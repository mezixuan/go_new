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

	//需求:
	//server可以接收多个连接，===》主go程负责监听，子go程负责数据处理
	//每个连接可以接收处理多轮数据请求
	//Accept
	for {
		fmt.Println("监听中")

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
		}

		fmt.Println("连接建立成功!")
		go handleFunc(conn)
		//conn.Close()
	}
}

//处理具体业务的逻辑，需要将conn传递进来，每一个新的连接，conn是彼此独立的
func handleFunc(conn net.Conn) {
	for {
		//创建一个容器，用于接收读取到的数据
		buf := make([]byte, 1024) //使用make来创建切片

		//cnt:真正读取client发来的数据的长度
		fmt.Println("准备读取客户端发送的数据...")
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
	}
	_ = conn.Close()
}
