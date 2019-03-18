package main

import (
	"fmt"
	"log"
	"reflect"
)

func main() {
	s1 := struct {
		age  int
		name string
	}{age: 1, name: "1"}
	s2 := struct {
		age  int
		name string
	}{age: 1, name: "1"}

	if s1 == s2 {
		log.Println("s1 == s2")
	}

	s3 := struct {
		name string
		age  int
	}{age: 1, name: "1"}

	if s1 == s3 { // 结构体中属性顺序不相同
		log.Println("s1 == s3")
	}

	if reflect.DeepEqual(s1, s3) {
		fmt.Println("s1==s3")
	} else {
		fmt.Println("s1!=s3")
	}

	s4 := struct {
		age int
		m   map[string]string
	}{age: 1, m: map[string]string{"a": "b"}}
	s5 := struct {
		age int
		m   map[string]string
	}{age: 1, m: map[string]string{"a": "b"}}
	if s4 == s5 { // 结构体中包含不可比较的类型
		log.Println("s4 == s5")
	}
}

// 结构体中属性顺序不相同不可比较
// 结构体中包含不可比较的类型 map slice
// 可以使用reflect.DeepEqual进行比较
