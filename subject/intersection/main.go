package main

import (
	"errors"
	"fmt"
)

func main() {
	s1 := []int{1, 4, 5, 4, 7, 8, 6, 2}
	s2 := []int{2, 4, 6, 3, 4}
	result1, _ := intersect1(s1, s2)
	fmt.Println(result1)
	result2, _ := intersect2(s1, s2)
	fmt.Println(result2)
}

func intersect0(s1, s2 []int) ([]int, error) {
	if s1 == nil || s2 == nil {
		return nil, errors.New("nil silce")
	}
	result := []int{}
	if len(s1) == 0 || len(s2) == 0 {
		return result, nil
	}
	for _, v1 := range s1 {
		for _, v2 := range s2 {
			if v1 == v2 {
				result = append(result, v1)
				break
			}
		}
	}
	return result, nil
}

func intersect1(s1, s2 []int) ([]int, error) {
	if s1 == nil || s2 == nil {
		return nil, errors.New("nil silce")
	}
	result := []int{}
	if len(s1) == 0 || len(s2) == 0 {
		return result, nil
	}
	tempMap := make(map[int]bool)
	for _, v := range s1 {
		tempMap[v] = true
	}
	for _, v := range s2 {
		if c, ok := tempMap[v]; ok && c {
			result = append(result, v)
			tempMap[v] = false
		}
	}
	return result, nil
}

func intersect2(s1, s2 []int) ([]int, error) {
	if s1 == nil || s2 == nil {
		return nil, errors.New("nil silce")
	}
	result := []int{}
	if len(s1) == 0 || len(s2) == 0 {
		return result, nil
	}
	tempMap := make(map[int]int)
	for _, v := range s1 {
		tempMap[v]++
	}
	for _, v := range s2 {
		if c, ok := tempMap[v]; ok && c > 0 {
			result = append(result, v)
			tempMap[v]--
		}
	}
	return result, nil
}
