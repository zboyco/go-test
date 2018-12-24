package main

import (
	"fmt"
	"sync"
	"time"
)

type model struct {
	m     map[int]int
	mutex sync.Mutex
}

// main 测试互斥锁
func main() {
	fmt.Println("测试开始")

	obj := model{
		m: make(map[int]int),
	}

	// 开启goroutine循环放开和锁住model
	go lock(&obj)

	// 开启goroutine循环插入新数据
	go insert(&obj)

	// 循环输出map内容长度
	for {
		fmt.Println(len(obj.m))
		time.Sleep(time.Second)
	}
}

// lock 锁住5秒,放开5秒
func lock(obj *model) {
	for {
		obj.mutex.Lock()
		fmt.Println("锁住了...")
		time.Sleep(5 * time.Second)
		fmt.Println("解锁了...")
		obj.mutex.Unlock()
		time.Sleep(5 * time.Second)
	}
}

// insert 循环插入新数据
func insert(obj *model) {
	i := 0
	for {
		i++
		obj.mutex.Lock()
		obj.m[i] = i
		obj.mutex.Unlock()
		time.Sleep(time.Second)
	}
}
