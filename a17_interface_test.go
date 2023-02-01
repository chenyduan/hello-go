package main

import (
	"fmt"
	"testing"
)

// 接口

// 接口内只能有方法

type Animal interface {
	//叫
	bark()
}

type Dog struct {
	name string
}

// 结构体实现接口
func (dog Dog) bark() {
	fmt.Println("dog bark ", dog.name)
}

func TestInteface(t *testing.T) {
	dog := Dog{name: "小灰"}
	dog.bark()
}
