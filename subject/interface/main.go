package main

import "log"

type People interface {
	Speak(string) string
}

type Student struct{}

func (s *Student) Speak(think string) (talk string) {
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
// golang的方法集仅仅影响接口实现和方法表达式转化，与通过实例或者指针调用方法无关。
