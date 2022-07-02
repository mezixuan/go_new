package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

type World struct {
}

func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + "你好!"
	return nil
}

func main() {
	//1. 注册rpc服务，绑定对象方法
	//err := rpc.RegisterName("hello", new(World))
	//if err != nil {
	//	fmt.Println("注册rpc任务失败，err:", err)
	//	return
	//}

	//registerService(new(World))
	//fmt.Println("注册rpc成功...")
	//2. 设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("net.listen err", err)
		return
	}
	defer listener.Close()
	fmt.Println("开始监听...")
	//3. 建立连接
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("listener.Accept err", err)
		return
	}
	defer conn.Close()
	fmt.Println("已经建立连接...")
	//4. 绑定服务
	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
	fmt.Println("绑定成功...")
}
