package main

import (
	"fmt"
	"testing"
)

//全局变量

var a = 10
var str = "你好"

func TestGlobalVariate(t *testing.T) {

	var a = 20
	// 输出20 ，局部变量会覆盖全局变量
	fmt.Println(a)

}
