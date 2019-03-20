package main

import (
	"log"
	"time"

	"github.com/zboyco/go-test/queue/slice/queue"
)

func main() {
	q, err := queue.InitQueue(10)

	if err != nil {
		log.Panicln(err)
	}

	go read(q, 1)
	go read(q, 2)
	go read(q, 3)

	i := 0
	for {
		time.Sleep(50 * time.Millisecond)
		err := q.EnQueue(i)
		if err != nil {
			// log.Println(err)
			continue
		}
		i++
	}
}

func read(q *queue.Queue, id int) {
	for {
		item, err := q.DeQueue()
		if err != nil {
			//log.Println(err)
			//time.Sleep(1 * time.Second)
			continue
		}
		log.Println("id", id, "------", item.(int))
	}
}
