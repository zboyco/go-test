package main

import "log"

func main() {
	i := GetValue()
	switch i.(type) {
	case int:
		log.Println("int")
	default:
		log.Println("Unknown")
	}
}

func GetValue() int {
	return 1
}

// 编译失败，因为type只能使用在interface

// func GetValue() interface{} {
// 	return 1
// }
