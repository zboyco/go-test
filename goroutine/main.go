package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		wg.Done()
		go func() {
			for {
				fmt.Println("test")
				time.Sleep(1)
			}
		}()

	}()

	wg.Wait()

	time.Sleep(10)
}
