package main

import "testing"

/*
方法报错了，因为go只允许有一个main函数

	func main() {
		var a bool = true
		fmt.Println(a)
	}
*/

// go test
//文件名以_test.go结尾
// 函数名以 TestXXX开头

func TestDataType(t *testing.T) {

	var a bool = true
	println("布尔类型 a的值为：", a)

	// 自动类型推断
	var b = false
	println("布尔类型 b的值为：", b)
}

// 整型
func TestDataType2(t *testing.T) {
	var a int = 1
	println(a)

	// int8 的范围为 -128 ~ 127
	var b int8 = -128
	println(b)

	// uint 无符号整形
	var c uint = 127
	println(c)
}

// 浮点数
func TestFloat(t *testing.T) {
	var a float32 = 3.14
	println(a)
}
