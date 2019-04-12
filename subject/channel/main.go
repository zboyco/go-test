package main

import (
	"fmt"
	"time"
)

func main() {
	// 创建一个缓冲通道
	ch := make(chan int, 1000)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}

	}()

	go func() {
		for {
			a, ok := <-ch

			if !ok {
				fmt.Println("close")
				close(ch)
				return
			}
			fmt.Println("a: ", a)
		}
	}()

	fmt.Println("ok")
	time.Sleep(10 * time.Second)
}
