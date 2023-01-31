package main

import (
	"fmt"
	"testing"
)

// range 表示下标和值或者   map 的key value

func TestRange(t *testing.T) {
	a := [3]string{"a", "b", "c"}

	for index, val := range a {
		fmt.Println(index)
		fmt.Println(val)

	}

}
