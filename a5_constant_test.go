package main

import (
	"fmt"
	"testing"
)

// 常量

func TestConstant(t *testing.T) {

	const a string = "3.14"

	// a = 3.15  报错，常量不能再次赋值
	fmt.Println(a)

	// 常量作为枚举
	const (
		Unknown = 0
		Male    = 1
		Famale  = 2
	)

	println(Unknown)

}

// iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
func TestIota(t *testing.T) {

	const (
		a = iota
		b = iota
		c = iota
	)

	println(a)
	println(b)
	println(c)

}
