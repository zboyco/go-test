package main

import (
	"fmt"
)

func main() {
	s := []int{6, 3, 7, 2, 9, 8, 4, 0, 1, 5}

	Quick3Sort(s, 0, len(s)-1)
	fmt.Println(s)
}

func Quick3Sort(a []int, left int, right int) {
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if a[left] >= a[i] {
			//分割位定位++
			explodeIndex++
			a[i], a[explodeIndex] = a[explodeIndex], a[i]
			fmt.Println(a, " i =", i, "  e =", explodeIndex)
		}
	}
	//起始位和分割位
	a[left], a[explodeIndex] = a[explodeIndex], a[left]

	Quick3Sort(a, left, explodeIndex-1)
	Quick3Sort(a, explodeIndex+1, right)
}
