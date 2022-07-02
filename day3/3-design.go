package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

//要求，服务端在注册rpc对象时，能让编译器检测出 注册对象是否合法

//创建接口，在接口中定义方法原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

//调用该方法时，需要给 i 传参，参数应该是实现了 HelloWorld 方法的类对象;
func registerService(i MyInterface) {
	rpc.RegisterName("hello", i)
}

//----------------
//像调用本地函数一样，调用远程函数

//定义一个类
type MyClient struct {
	c *rpc.Client
}

func InitClient(addr string) MyClient {
	conn, _ := jsonrpc.Dial("tcp", addr)
	return MyClient{c: conn}
}

//实现函数,原型式参照上面的interface来实现
func (this *MyClient) HelloWorld(a string, b *string) error {
	return this.c.Call("hello.HelloWorld", a, b)
}
