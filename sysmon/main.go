package main

import (
	"fmt"
	"runtime"
	"time"
)

// Main Goroutine
func main() {
	// 模拟单核 CPU
	runtime.GOMAXPROCS(1)

	// 模拟 Goroutine 死循环
	go func() {
		for {
		}
	}()

	time.Sleep(time.Millisecond)
	fmt.Println("123")
}
