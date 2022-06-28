package main

import (
	"fmt"
	"net/rpc"
)

func main() {
	conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		fmt.Println("rpc.Dial err", err)
		return
	}
	defer conn.Close()
	var reply string
	err = conn.Call("hello.HelloWorld", "赵津樊", &reply)
	if err != nil {
		fmt.Println("conn.Close err", err)
		return
	}
	fmt.Println("reply:", reply)
}
