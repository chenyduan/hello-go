package main

import (
	"fmt"
	"testing"
)

//切片

func TestSlice(t *testing.T) {
	// 声明切片
	var a []int8
	fmt.Println(a)
	// 声明切片方式2: 长度可选
	var b []int = make([]int, 3)
	fmt.Println(b)
	c := make([]int, 3)
	fmt.Println(c)

	// 初识长度为3 ，最大长度为 5
	d := make([]int, 3, 5)
	fmt.Println(d)

	//长度 len(),cap()
	fmt.Println(len(d))
	fmt.Println(cap(d))
}

// 切片初始化,截取
func TestSlice2(t *testing.T) {
	// 初始化切片
	a := []int{1, 2, 3}
	fmt.Println(a)
	//数组转切片
	arr := [3]string{"a", "b", "c"}
	b := arr[:]
	fmt.Println("b=", b)

	// 取arr 0到（2-1）个元素
	c := arr[0:2]
	fmt.Println(c)

	//取0到 （len-1） 个元素
	d := arr[0:]
	fmt.Println(d)

	//长度 len(),cap()
	fmt.Println(len(a))
	fmt.Println(cap(a))
}

//切片 copy append

func TestSlice3(t *testing.T) {
	a := make([]int, 0)
	b := append(a, 0, 1, 2)
	fmt.Println(a)
	fmt.Println(b)

}
