package main

import (
	"fmt"
	"testing"
)

// 类型转换

func TestChangeType(t *testing.T) {

	var a int8 = 16
	b := float32(a)
	fmt.Println(b)
}
