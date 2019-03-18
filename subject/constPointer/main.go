package main

import "log"

const a = 1

var b = 2

func main() {
	log.Println(a, &a) //常量在预处理阶段直接展开，不能获取常量地址
	log.Println(b, &b)
}
