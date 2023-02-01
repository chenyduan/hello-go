package main

import (
	"fmt"
	"testing"
)

// 通道 channel
func TestChannel(t *testing.T) {

	// 声明通道
	ch1 := make(chan int)
	// 默认情况通道不带缓冲
	go chanMet(ch1)

	// 从通道取出值
	a := <-ch1
	fmt.Println(a)
}

func chanMet(c chan int) {
	//把 2 发送到通道 c
	c <- 2

}

// 带缓冲区的通道

func TestBufferChannel(t *testing.T) {

	chBuff := make(chan int, 100)
	go chanMet(chBuff)

	a := <-chBuff
	fmt.Println(a)

}