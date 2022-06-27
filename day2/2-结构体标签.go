package main

import (
	"encoding/json"
	"fmt"
)

type Teacher struct {
	Name    string `json:"-"`            //在使用json编码时，这个编码不参与
	Age     int    `json:"age,string"`   //在json编码时，会编码成age，并且类型转换成string，一定是两个字段中间逗号没有空格
	Subject string `json:"Subject_name"` //json编码时，这个字段会编码成Subject_name
	gender  string
	Address string `json:"address,omitempty"` //如果在json编码时，如果这个字段是空，那么就忽略掉，不参与编码
}

func main() {
	t1 := Teacher{
		Name:    "zjf",
		Age:     22,
		Subject: "Golang",
		gender:  "男",
		Address: "北京",
	}
	fmt.Println("t1:", t1)
	encodeInfo, _ := json.Marshal(&t1)
	fmt.Println("encodeInfo:", string(encodeInfo))
}
