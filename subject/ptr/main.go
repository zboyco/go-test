package main

import "log"

type Object struct {
	value string
}

// 值拷贝
func (o Object) ValuePrint() {
	o.value = "bbb"
	log.Println(o.value)
}

// 指针
func (o *Object) PtrPrint() {
	o.value = "ccc"
	log.Println(o.value)
}

func main() {
	obj := Object{"aaa"}
	log.Println(obj.value)

	obj.ValuePrint()
	log.Println(obj.value)

	obj.PtrPrint()
	log.Println(obj.value)
}
