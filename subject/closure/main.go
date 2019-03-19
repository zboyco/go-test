package main

import "log"

func test1() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		funcs = append(funcs, func() {
			log.Println(i, &i) // 每次放入匿名函数中的i都是同一个变量
		})
	}
	return funcs
}

func test2() []func() {
	var funcs []func()
	for i := 0; i < 2; i++ {
		x := i // 新建临时变量保存i的值
		funcs = append(funcs, func() {
			log.Println(i, &i, x, &x)
		})
	}
	return funcs
}

// 返回的两个函数使用同一个x变量
func test3(x int) (func(), func()) {
	return func() {
			log.Println(x)
			x += 10
		}, func() {
			log.Println(x)
		}
}

func main() {
	funcs1 := test1()
	for _, f := range funcs1 {
		f()
	}
	funcs2 := test2()
	for _, f := range funcs2 {
		f()
	}
	f1, f2 := test3(100)
	f1() // 100 x: 110
	f2() // 110 x: 110
	f1() // 110 x: 120
}
