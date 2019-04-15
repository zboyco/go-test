package main

import (
	"errors"
	"fmt"
	"sync"
)

// SafeStack 线程安全栈
type SafeStack struct {
	data []interface{}
	sync.RWMutex
}

func (s *SafeStack) initSilce() {
	if s.data == nil {
		s.data = make([]interface{}, 0)
	}
}

// Push Push
func (s *SafeStack) Push(v interface{}) {
	s.Lock()
	defer s.Unlock()
	s.initSilce()
	s.data = append(s.data, v)
}

// Pop Pop
func (s *SafeStack) Pop() (interface{}, error) {
	s.Lock()
	defer s.Unlock()
	s.initSilce()
	lenght := len(s.data)
	if lenght == 0 {
		return nil, errors.New("empty")
	}
	result := s.data[lenght-1]
	s.data = s.data[:lenght-1]
	return result, nil
}

// Top Top
func (s *SafeStack) Top() (interface{}, error) {
	s.RLock()
	defer s.RUnlock()
	if s.data == nil || len(s.data) == 0 {
		return nil, errors.New("empty")
	}
	return s.data[len(s.data)-1], nil
}

// IsEmpty IsEmpty
func (s *SafeStack) IsEmpty() bool {
	s.RLock()
	defer s.RUnlock()
	return s.data == nil || len(s.data) == 0
}

func main() {
	s := &SafeStack{}

	fmt.Println(s.IsEmpty())
	_, err1 := s.Top()
	if err1 != nil {
		fmt.Println(err1)
	}
	_, err2 := s.Pop()
	if err2 != nil {
		fmt.Println(err2)
	}
	for i := 0; i < 10; i++ {
		s.Push(i)
	}
	fmt.Println(s.IsEmpty())
	v3, err3 := s.Top()
	if err3 != nil {
		fmt.Println(err3)
	}
	fmt.Println(v3)
	v4, err4 := s.Pop()
	if err4 != nil {
		fmt.Println(err4)
	}
	fmt.Println(v4)
	for {
		v5, err5 := s.Pop()
		if err5 != nil {
			fmt.Println(err5)
			break
		}
		fmt.Println(v5)
	}
}
