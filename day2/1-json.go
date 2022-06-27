package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Id     int
	Name   string
	Age    int
	gender string //注意,gender是小写的
}

func main() {
	lily := Student{
		Id:     1,
		Name:   "lily",
		Age:    22,
		gender: "女",
	}

	//编码
	encodeInfo, err := json.Marshal(&lily)
	if err != nil {
		fmt.Println("json.Marshal err:", err)
		return
	}
	fmt.Println("encodeInfo:", string(encodeInfo))

	//对端接收到数据
	//解码
	var lily2 Student
	if err := json.Unmarshal([]byte(encodeInfo), &lily2); err != nil {
		fmt.Println("json.Unmarshal err:", err)
		return
	}
	fmt.Println("Name:", lily2.Name)
	fmt.Println("gender:", lily2.gender)
	fmt.Println("Age:", lily2.Age)
	fmt.Println("Id:", lily2.Id)
}
