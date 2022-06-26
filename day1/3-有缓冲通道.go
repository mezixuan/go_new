package main

import (
	"fmt"
	"time"
)

func main() {
	//numChan := make(chan int, 10)
	//1. 当缓冲区写满的时候，写阻塞，当读取后，再恢复写入
	//2. 当缓冲区读取完毕，读阻塞
	//3. 如果管道没有使用make分配空间，那么管道默认是nil，读取、写入都会阻塞
	//4. 对于一个管道，读与写的次数，必须对等，会出现死锁
	names := make(chan string, 10)
	go func() {
		fmt.Println("names:", <-names)
	}()
	names <- "hello"
	time.Sleep(1 * time.Second)
}
