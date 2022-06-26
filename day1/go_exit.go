package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		go func() {
			func() {
				fmt.Println("这是子go程内部的函数")
				//os.Exit(-1)
				runtime.Goexit()
			}()
			fmt.Println("子go程结束！")
			fmt.Println("go 2222222")
		}()
		time.Sleep(2 * time.Second)
		fmt.Println("go 1111111")
	}()
	fmt.Println("这是主go程")
	time.Sleep(3 * time.Second)
	fmt.Println("Over!")
}
