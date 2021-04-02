package main

import "fmt"

func main() {
	test(nil)
}

func test(s []string) {
	for _, v := range s {
		fmt.Println(v)
	}
}
