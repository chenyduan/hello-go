package main

import (
	"fmt"
	"testing"
	"time"
)

// goroutine

func TestGoroutine(t *testing.T) {
	go say("你好")
	say("hello")
}

func say(a string) {
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println(a)
	}
}
