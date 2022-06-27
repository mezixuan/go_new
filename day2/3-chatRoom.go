package main

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type User struct {
	name string
	id   string
	msg  chan string
}

//创建一个全局的map结构，用户保存所有的用户
var allUsers = make(map[string]User)

//定义一个全局的message通道，用于接收任何人发送过来的消息
var message = make(chan string, 10)

func main() {
	//创建服务器
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("net.listen err:", err)
		return
	}
	//启动全局唯一的go程，负责监听message管道，写入到每个用户的msg中
	go broadCast()
	fmt.Println("服务器启动成功，监听中....")
	for {
		//监听
		fmt.Println("主go程监听中！")
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err", err)
			return
		}
		//建立连接
		fmt.Println("建立连接成功!")
		//启动业务的go程
		go handler(conn)
	}
}

//处理具体业务
func handler(conn net.Conn) {
	fmt.Println("启动业务!")

	//客户端与服务器建立连接的时候，会有ip和端口号==》当成user的id
	clientAddr := conn.RemoteAddr().String()
	fmt.Println("conn.RemoteAddr:", clientAddr)
	newUser := User{
		name: clientAddr,            //可以修改可以提供rename来修改
		id:   clientAddr,            //不可以修改，唯一标识符
		msg:  make(chan string, 10), //注意要分配make空间
	}
	//添加User到map结构
	allUsers[newUser.id] = newUser

	//定义一个退出信号，用于监听client退出
	var isQuit = make(chan bool)

	//创建一个用于重置计数器的管道，用于告知watch函数，当前用户正在输入
	var restTimer = make(chan bool)

	go watch(&newUser, conn, isQuit, restTimer)

	//启动go程，负责将msg信息返回给go程
	go writeBackToClient(&newUser, conn)

	//向message写入数据，当前用户上线的消息,广播用于通知所有人
	loginInfo := fmt.Sprintf("[%s]:[%s] ==> 上线了login!\n", newUser.id, newUser.name)
	message <- loginInfo

	for {
		//具体业务
		//读取客户端发过来的数据
		buf := make([]byte, 1024)
		cnt, err := conn.Read(buf)
		if cnt == 0 {
			fmt.Println("客户端主动关闭ctrl + c，准备退出!")
			//map删除用户，用户，conn close掉
			//服务器可以主动退出
			//在这里不进行真正的退出动作，而是发送一个退出信号，统一做退出处理，可以使用新的管道来做信号传递
			isQuit <- true
		}
		if err != nil {
			fmt.Println("conn.Read err:", err, "cnt", cnt)
		}
		fmt.Println("服务器接受客户端发过来的数据，数据为:", string(buf[:cnt-1]), "cnt:", cnt)
		//------------业务逻辑处理 开始---------
		//1. 查询当前所有的用户 who
		//	a.先判断接受的数据是不是who ==》 长度&&字符串
		//  b.遍历allUsers这个map(key:userid  value:user本身)，将id和name拼接成一个字符串，返回给客户端
		userInput := string(buf[:cnt-1])
		if string(userInput) == "\\who" && len(userInput) == 4 {
			fmt.Println("用户即将查询所有用户信息!")
			//这个切片包含所有的用户信息
			var userInfos []string
			for _, user := range allUsers {
				userInfo := fmt.Sprintf("userId:%s, username:%s", user.id, user.name) //这个不用加\n
				userInfos = append(userInfos, userInfo)
			}
			//最终写到管道中，一定是一个字符串
			r := strings.Join(userInfos, "\n")

			newUser.msg <- r
		} else if len(userInput) > 9 && userInput[:7] == "\\rename" {
			//arry := strings.Split(userInput, "|")
			//name := arry[1]
			newUser.name = strings.Split(userInput, "|")[1]
			allUsers[newUser.id] = newUser //更新map中的user
			newUser.msg <- "rename successfully!"
		} else {
			message <- userInput
		}
		restTimer <- true
		//------------业务逻辑处理 结束---------
	}
}

//向所有的用户广播消息，启动一个全局唯一的go程
func broadCast() {
	fmt.Println("广播go程启动成功...")
	defer fmt.Println("broadcast 程序退出!")

	for {
		//1. 从message中读取数据
		fmt.Println("broadCast监听message中...")
		info := <-message
		//2. 将数据写入到每一个用户的msg管道中
		//如果message是非缓冲的，那么msg会阻塞
		for _, User := range allUsers {
			User.msg <- info
		}
	}
}

func writeBackToClient(user *User, conn net.Conn) {
	fmt.Printf("user : %s的go程正在监听自己的msg管道:\n", user.name)
	for data := range user.msg {
		fmt.Printf("user : %s,写回给客户端数据为:%s\n", user.name, data)
		_, _ = conn.Write([]byte(data))
	}
}

//启动一个go程，负责监听退出信号，触发后，进行清零工作 :delete, map, close, conn都在这里处理
func watch(user *User, conn net.Conn, isQuit, restTimer <-chan bool) {
	fmt.Println("启动监听退出信号的go程...")
	defer fmt.Println("watch go程退出!")
	for {
		select {
		case <-isQuit:
			logoutInfo := fmt.Sprintf("%s exit already!\n", user.name)
			fmt.Println("删除当前用户:", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo
			conn.Close()
			return
		case <-time.After(120 * time.Second):
			logoutInfo := fmt.Sprintf("%s timeout exit already!\n", user.name)
			fmt.Println("删除当前用户:", user.name)
			delete(allUsers, user.id)
			message <- logoutInfo
			conn.Close()
			return
		case <-restTimer:
			fmt.Printf("连接%s 重置结束计数器!\n", user.name)
		}
	}
}
