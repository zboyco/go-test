package main

import "log"

func main() {

	s := []int{1, 2, 3}

	testSlice(s)
	log.Println(s)
	testSlicePrt(&s)
	log.Println(s)

	m := map[int]string{
		0: "0",
		1: "2",
	}
	testMap(m)
	log.Println(m)
	testMapPtr(&m)
	log.Println(m)
}

func testSlice(s []int) {
	s[0] = 999
}

func testSlicePrt(s *[]int) {
	(*s)[0] = 888
}

func testMap(m map[int]string) {
	m[0] = "999"
}

func testMapPtr(m *map[int]string) {
	(*m)[0] = "888"
}
