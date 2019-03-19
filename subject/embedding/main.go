package main

import "log"

type People struct{}

func (p People) ShowA() {
	log.Println("ShowA")
	p.ShowB()
}

func (p People) ShowB() {
	log.Println("ShowB")
}

type Teacher struct {
	People
}

func (t Teacher) ShowB() {
	log.Println("Teacher Show B")
}

func main() {
	t := &Teacher{}
	t.ShowA() // 等价于 t.People.ShowA()
	t.People.ShowA()
	t.ShowB()
}

// Go没有继承，只有组合，t.ShowA() 等价于 t.People.ShowA()
