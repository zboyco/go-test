package main

import "log"

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println(err)
		} else {
			log.Println("error")
		}
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("panic")
}

// panic 只有最后一个可以被recover捕捉到，defer panic 覆盖了先触发的 panic
