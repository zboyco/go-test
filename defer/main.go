package main

import "fmt"

const zero = 0.0

func main() {
	for i := 0; i < 10; i++ {
		defer fmt.Printf("%d ", i)
	}
	s := []int{1, 2, 3, 4, 5}

	for i := range s {
		defer fmt.Printf("%d ", s[i])
	}
}
