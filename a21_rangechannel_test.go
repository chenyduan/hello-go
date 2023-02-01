package main

import (
	"fmt"
	"testing"
)

//遍历通道

func TestRangeChannel(t *testing.T) {

	ch1 := make(chan int)
	go sum(ch1)
	// 通道这里只有value
	for v := range ch1 {
		fmt.Println(v)
	}
}

func sum(a chan int) {
	for i := 0; i < 5; i++ {
		a <- i
	}
	close(a)
}
