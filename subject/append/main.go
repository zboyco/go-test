package main

import "log"

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2...)
	log.Println(s1)
}

// append 添加切片不能漏掉...
