package main

import (
	"fmt"
	"time"
)

func main() {
	//生产者消费者模型  consumer--消费者， producer--生产者
	//C++:   数组+锁    thread1：写  thread2：读
	//Go: goroutine +channel
	//1. 在主函数中创建一个双向通道 numChan
	numChan1 := make(chan int, 5)
	//2. 将numChan传递给producer， 负责生产
	go producer(numChan1) //双向通道可以赋值给同类型的单向通道，单向不能转双向

	//3. 将numChan传递给consumer， 负责消费
	go consumer(numChan1)

	time.Sleep(1 * time.Second)
	fmt.Println("Over!")
}

//producer 生产者 ==》 提供一个只写的通道
func producer(out chan<- int) {
	for i := 0; i < 50; i++ {
		out <- i
		fmt.Println("向管道中写数据:", i)
	}
}

//consumer 消费者 ==》 提供一个只读的通道
func consumer(in <-chan int) {
	for v := range in {
		fmt.Println("从管道往出读取数据:", v)
	}
}
