package main

import (
	"log"
	"sync"
)

type threadSafeSet struct {
	sync.RWMutex
	s []interface{}
}

func (set *threadSafeSet) Iter() <-chan interface{} {
	// ch := make(chan interface{}) // 无缓冲通道造成阻塞
	ch := make(chan interface{}, len(set.s))
	go func() {
		set.RLock()
		for elem, value := range set.s {
			ch <- value
			log.Println("Iter:", elem, value)
		}
		close(ch)
		set.RUnlock()
	}()
	return ch
}

func main() {
	th := threadSafeSet{
		s: []interface{}{"1", 2},
	}
	ch := th.Iter()
	v1 := <-ch
	log.Println("ch", v1)
	v2 := <-ch
	log.Println("ch", v2)
}
