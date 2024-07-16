package main

import (
	"fmt"
)

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}
func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}
func (s *Stack) Peek() int {
	item := s.items[len(s.items)-1]
	return item
}
func (s *Stack) Reverse() []int {
	for i, j := 0, len(s.items)-1; i < j; i, j = i+1, j-1 {
		s.items[i], s.items[j] = s.items[j], s.items[i]
	}
	return s.items
}

func main() {
	stack := &Stack{}
	stack.Push(5)
	stack.Push(6)
	fmt.Println("Stack :", stack)
	//stack.Pop()
	fmt.Println("After popping one element:", stack)
	fmt.Println("On Peeking:", stack.Peek())
	fmt.Println("On Reversing the values present in the stack:", stack.Reverse())
}
