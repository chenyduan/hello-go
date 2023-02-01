package main

import (
	"errors"
	"fmt"
	"testing"
)

// 自定义错误信息

func errorMsg(a int) (int, error) {
	if a < 2 {
		return 0, errors.New("数据小于2")
	}
	return a, nil
}
func TestErrorMs(t *testing.T) {
	var a, err = errorMsg(0)
	fmt.Println(a)
	fmt.Println(err)
}

// 错误处理
