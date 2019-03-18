package main

import "log"

func main() {
	for i := 0; i < 10; i++ {
	loop:
		log.Println(i)
	}
	// goto loop jumps into block starting
}

// goto 不能跳转到其他函数或者内层代码块
