package main

import (
	"log"
	"time"
)

func main() {
	c := make(chan interface{}, 10)

	go read(c, 1)
	go read(c, 2)
	go read(c, 3)

	for i := 0; i < 15; i++ {
		time.Sleep(100 * time.Millisecond)
		c <- i
	}
	close(c)
	time.Sleep(5 * time.Second)
}

func read(c <-chan interface{}, id int) {
	for {
		item, ok := <-c
		if !ok {
			log.Println("channel close")
			break
		}
		log.Println("id", id, "------", item.(int))
	}
}
