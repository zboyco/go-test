package main

import "fmt"

func main() {
	var i int
	var a [10]int
	var s []int
	var p *int
	var m map[int]int
	var f func(int) int
	var c <-chan int
	fmt.Println(i, a, s, p, m, f, c)
}
