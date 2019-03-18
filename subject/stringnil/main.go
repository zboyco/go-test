package main

import "log"

func GetValue(m map[int]string, id int) (string, bool) {
	if _, ok := m[id]; ok {
		return "OK", true
	}
	return nil, false // cannot use nil as type string in return argument,replace ""
}

func main() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}
	v, err := GetValue(intmap, 3)
	log.Println(v, err)
}
