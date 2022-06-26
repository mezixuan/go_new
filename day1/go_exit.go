package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		func() {
			fmt.Println("这是子go程内部的函数")
			//os.Exit(-1)
			runtime.Goexit()
		}()
		fmt.Println("子go程结束！")
	}()
	fmt.Println("这里是主go程")
	time.Sleep(5)
	fmt.Println("OVER")
}
