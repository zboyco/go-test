package main

import "fmt"

func main() {
	fmt.Println(sliceIsNil(nil))
	fmt.Println(sliceIsEmpty(nil))
}

func sliceIsNil(slice []interface{}) bool {
	return slice == nil
}

func sliceIsEmpty(slice []interface{}) bool {
	return len(slice) == 0
}
