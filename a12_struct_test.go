package main

import (
	"fmt"
	"testing"
)

// 结构体

type User struct {
	name string
	age  int8
	id   int16
}

func TestStruct(t *testing.T) {
	// 声明结构体数据
	zhangsan := User{name: "张三", age: 18, id: 1}
	fmt.Println(zhangsan)
	lisi := User{"莉丝", 20, 2}
	fmt.Println(lisi)
}
