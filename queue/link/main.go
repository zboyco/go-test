package main

import (
	"log"
	"time"

	"github.com/zboyco/go-test/queue/link/queue"
)

func main() {
	q := &queue.Queue{}

	go read(q, 1)
	go read(q, 2)
	go read(q, 3)

	i := 0
	for {
		time.Sleep(50 * time.Millisecond)
		err := q.Enqueue(i)
		if err != nil {
			// log.Println(err)
			continue
		}
		i++
	}
}

func read(q *queue.Queue, id int) {
	for {
		item, err := q.Dequeue()
		if err != nil {
			// log.Println(err)
			// time.Sleep(100 * time.Millisecond)
			continue
		}
		log.Println("id", id, "------", item.(int))
	}
}
