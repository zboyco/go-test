package main

import (
	"errors"
	"log"
)

var ErrorDoNotWork error = errors.New("Do Not Work")

func TryTheThing() (string, error) {
	return "", ErrorDoNotWork
}

func DoTheThing(really bool) (err error) {
	if really {
		result, err := TryTheThing() // err 覆盖了 自身函数返回值err，这里的err是代码块内临时变量
		if err != nil || result != "It Worked" {
			err = ErrorDoNotWork
		}
	}
	return err
}

func main() {
	log.Println(DoTheThing(true))
	log.Println(DoTheThing(false))
}

// 注意DoTheThing 返回变量err作用域
