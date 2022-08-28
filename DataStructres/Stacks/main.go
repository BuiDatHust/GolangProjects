package main

import "fmt"

type Stack struct {
	data []interface{}
	top  int
}

func (s *Stack) Push(element interface{}) {
	s.top++
	s.data = append(s.data, element)
}

func (s *Stack) Pop() interface{} {
	if len(s.data) > 0 {
		s.top--
		last := s.data[s.top]
		s.data = s.data[:s.top]

		return last
	}
	return nil
}

func (s *Stack) Peek() interface{} {
	if len(s.data) > 0 {
		return s.data[s.top-1]
	}
	return nil
}

func (s *Stack) Clear() {
	s.top = 0
	s.data = []interface{}{}
}

func (s *Stack) Length() int {
	return s.top
}

func main() {
	s := Stack{}
	for i := 0; i < 5; i++ {
		s.Push(i)
	}
	fmt.Println(s.Peek())
}
