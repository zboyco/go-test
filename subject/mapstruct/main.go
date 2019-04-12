package main

import "fmt"

type Student struct {
	name string
}

func main() {
	m := map[string]Student{"people": {"liyuechun"}}
	// m["people"].name = "wuyanzu"
	var s Student = m["people"] //深拷贝
	s.name = "xietingfeng"
	fmt.Println(m)
	fmt.Println(s)
}

// 不能直接修改map值中结构的属性值
