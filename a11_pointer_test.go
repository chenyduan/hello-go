package main

import (
	"fmt"
	"testing"
)

func TestPointer(t *testing.T) {

	//声明一个指针
	var a *int
	fmt.Println("未分配时指针的值为nil,如：", a)
	var b = 123
	a = &b

	fmt.Println("变量值", b)
	fmt.Println("指针", a)
	fmt.Println("指针", &b)
	fmt.Println("指针获取变量值", *a)

}
