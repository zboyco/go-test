package main

import (
	"log"
)

type myInt int

func main() {
	var v1 myInt = 1
	v1.test1()
	v1.test2()
	test3(v1)
	// test4(v1) // error

	v2 := &v1
	v2.test1()
	v2.test2()
	// test3(v2) // error
	test4(v2)

	var v3 interface{} = &v1
	v3.(*myInt).test1()
	v3.(*myInt).test2()
	// test3(v3.(*myInt)) // error
	test4(v3.(*myInt))

	var v4 interface{} = v1
	v4.(myInt).test1()
	// v4.(myInt).test2() // error
	test3(v4.(myInt))
	// test4(v4.(myInt)) // error
}

// 值接收者可以接收值或者指针
func (v myInt) test1() {
	log.Println(v)
}

// 指针接受者只能接收指针
func (v *myInt) test2() {
	log.Println(*v)
}

func test3(v myInt) {
	log.Println(v)
}

func test4(v *myInt) {
	log.Println(*v)
}
