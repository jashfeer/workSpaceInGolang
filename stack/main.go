package main

import (
	"fmt"
)

type stack struct {
	list []int
}

func (s *stack) push(n int) {
	s.list = append(s.list, n)
}

func (s *stack) pop() int {
	l := len(s.list) - 1
	popValue := s.list[l]
	s.list = s.list[:l]
	return popValue
}

func main() {
	myStack := stack{}

	myStack.push(20)
	myStack.push(30)
	myStack.push(40)
	myStack.push(50)
	myStack.push(60)
	fmt.Println(myStack)
	n := myStack.pop()
	fmt.Println(n)
	fmt.Println(myStack)
}
