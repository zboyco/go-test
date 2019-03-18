package main

import "log"

func main() {
	log.Println(deferFunc1(1))
	log.Println(deferFunc2(1))
	log.Println(deferFunc3(1))
}

func deferFunc1(i int) (t int) {
	t = i
	defer func() {
		t += 3
	}()
	return t // 设置了 t = t
}

func deferFunc2(i int) int {
	t := i
	defer func() {
		t += 3
	}()
	return t
}

func deferFunc3(i int) (t int) {
	defer func() {
		log.Println("defer t:", t)
		t += i
	}()
	return 2 // 设置了 t = 2
}

// 需要明确一点是defer需要在函数结束前执行。
// 函数返回值名字会在函数起始处被初始化为对应类型的零值并且作用域为整个函数
// DeferFunc1有函数返回值t作用域为整个函数，在return之前defer会被执行，所以t会被修改，返回4;
// DeferFunc2函数中t的作用域为函数，返回1;
// DeferFunc3返回3
