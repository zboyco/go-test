package main

import "log"

type MyInt1 int // 定义类型

type MyInt2 = int // 设置别名

type T1 struct{}

func (t T1) m1() {
	log.Println("T1.m1()")
}

type T2 = T1
type T struct {
	T1
	T2
}

func main() {
	var i int = 6

	var m1 MyInt1 = i // cannot use i (type int) as type MyInt1 in assignment
	var m2 MyInt2 = i

	log.Println(m1, m2)

	var t T = T{}
	t.m1() // ambiguous selector t.m1
}

// MyInt1 是基于int的一个新类型，不能直接赋值，可以使用强转 var m1 MyInt1 = MyInt1(i)
// T2 完全等价于 T1, 所以不能t.m1()不能知道调用哪个方法，使用 t.T1.m1() 或者 t.T2.m1()
