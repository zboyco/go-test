package main

import "fmt"

type People struct {
	Name string
}

func (p *People) String() string {
	return fmt.Sprintf("print: %v", p)
}

func main() {
	p := &People{}
	p.String()
}

// 上面的代码出现了栈溢出，原因是因为%v格式化字符串是本身会调用String()方法
// 上面的栈溢出是因为无限递归所致。
