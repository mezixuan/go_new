package main

import (
	"fmt"
	"time"
)

func main() {
	//创建管道，创建一个装数字的管道 ==》 channel
	numChan := make(chan int, 10) //使用管道时，一定要make，同map一样
	//strChan := make(chan string) //装字符串的管道

	go func() {
		for i := 0; i < 50; i++ {
			data := <-numChan
			fmt.Println("子go程1,读出数据====>data：", data)
		}
	}()

	go func() {
		for i := 0; i < 20; i++ {
			numChan <- i
			fmt.Println("子go程2,写入数据====>data：", i)
		}
	}()

	//创建两个go程，父亲写数据，儿子读数据
	for i := 0; i < 30; i++ {
		numChan <- i
		fmt.Println("这是主go程，写入数据：", i)
	}

	time.Sleep(5 * time.Second)
}
