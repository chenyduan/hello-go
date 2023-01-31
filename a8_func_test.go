package main

import (
	"fmt"
	"testing"
)

//函数

func TestFun(t *testing.T) {
	var result = max(12, 32)
	fmt.Println(result)
}

// 定义一个函数，返回较大值
func max(a int8, b int8) int8 {
	if a > b {
		return a
	} else {
		return b
	}
}
