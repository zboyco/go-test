package main

import "fmt"

func main() {
	obj := testType{}
	obj.value = make(map[string][]string)

	obj.value["123"] = append(obj.value["123"], "1")
	obj.value["123"] = append(obj.value["123"], "2")
	obj.value["123"] = append(obj.value["123"], "3")
	fmt.Println(obj)
}

type testType struct {
	value map[string][]string
}
