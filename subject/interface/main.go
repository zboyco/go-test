package main

import "log"

type People interface {
	Speak(string) string
}

type Student struct{}

// func (s *Student) Speak(think string) (talk string) { // 编译失败，值类型 Student{} 未实现接口People的方法，不能定义为 People 类型。
func (s Student) Speak(think string) (talk string) {
	if think == "bitch" {
		talk = "you are a good boy"
	} else {
		talk = "hi"
	}
	return
}

func test(p People) {
	think := "bitch"
	log.Println(p.Speak(think))
}

func main() {
	// 	var peo = &Student{}
	// 	test(peo)
	var peo People = Student{}
	think := "bitch"
	log.Println(peo.Speak(think))
}

// 编译不通过！ 做错了！？说明你对golang的方法集还有一些疑问。
// 使用指针接收者实现一个接口，只有指向这个类型的指针才能实现对应的接口
// 使用值接收者实现一个接口，那么这个类型的值和指针都能够实现对应的接口
// 参考《GO语言实战》-方法集
