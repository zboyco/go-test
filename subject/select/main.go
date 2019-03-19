package main

import (
	"log"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case v := <-int_chan:
		log.Println(v)
	case v := <-string_chan:
		log.Panicln(v)
	}
}

// 可能出现恐慌，因为select具有随机性
