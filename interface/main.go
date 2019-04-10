package main

import "log"

type A interface {
	exec(i int) error
}
type SubA struct{}

func (v *SubA) exec(i int) error { // 接收者为指针则只有指针类型实现接口
	log.Println("exec ok")
	return nil
}

func main() {
	// var m1 A = SubA{}
	// m1.exec(0)
	var m2 A = &SubA{}
	m2.exec(0)
	var m3 A = new(SubA)
	m3.exec(0)
}
