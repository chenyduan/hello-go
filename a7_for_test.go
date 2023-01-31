package main

import "testing"

// for循环

func TestFor(t *testing.T) {
	var i = 0
	for i < 5 {
		println(i)
		i++
	}
}

// continue   break 跳出循环
