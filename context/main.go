package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx := context.Background()
	for i := 0; i < 10; i++ {
		ctx = context.WithValue(ctx, i, i)
	}
	nextCtx, cancelFunc := context.WithCancel(ctx)
	go func() {
		for {
			select {
			case <-nextCtx.Done():
				break
			default:
				fmt.Println("knock")
				time.Sleep(time.Second)
			}
		}
	}()
	time.Sleep(5 * time.Second)
	cancelFunc()
	time.Sleep(30 * time.Second)
}
