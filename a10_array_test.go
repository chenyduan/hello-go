package main

import (
	"fmt"
	"testing"
)

// 数组

func TestArray(t *testing.T) {
	var a [3]int8
	a[0] = 1
	a[1] = 2
	a[2] = 3

	fmt.Println(a)

	// 遍历数组
	var i = 0
	for i < len(a) {
		fmt.Println(a[i])
		i++
	}

	// 初始化数组

	var b = [3]string{"你好", "地球", "人"}
	fmt.Println(b)
}
