package main

import "log"

func main() {
	list := new([]int) // make([]int,0)
	list = append(list, 1)
	log.Println(list)
}

// 使用make([]int,0)创建长度为0的切片
