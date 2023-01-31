package main

import (
	"fmt"
	"testing"
)

//map

func TestMap(t *testing.T) {
	var b map[string]string
	fmt.Println(b)

	a := make(map[string]string)
	a["name"] = "zhangsan"
	a["age"] = "18"

	fmt.Println(a["name"])

	//遍历map
	for k, v := range a {
		fmt.Println("key={},val={}", k, v)
	}
	// 删除
	delete(a, "name")
	for k, v := range a {
		fmt.Println("删除后key={},val={}", k, v)
	}
}
