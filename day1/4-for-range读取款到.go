package main

import "fmt"

func main() {
	numsChan2 := make(chan int, 10)

	go func() {
		for i := 0; i < 50; i++ {
			numsChan2 <- i
			fmt.Println("写入数据:", i)
		}
		fmt.Println("数据全部写完毕，准备关闭管道")
		close(numsChan2)
	}()

	//遍历管道的时候只返回一个值，就是对应的通道里面的值
	//问题是range不知道管道是否写完，所以会一直等待
	//在写入端将管道关闭，for range遍历一个关闭的管道的时候就会退出
	for v := range numsChan2 {
		fmt.Println("读取数据:", v)
	}
	fmt.Println("Over!")
}
