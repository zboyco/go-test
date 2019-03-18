package main

import "log"

func Foo(x interface{}) {
	if x == nil {
		log.Println("emptyinterface")
	} else {
		log.Println("non-empryinterface")
	}
}

func main() {
	var x *int = nil
	Foo(x)
}

// interface内部结构
