package main

import "fmt"

type stack struct {
	size int
	top  int
	data []int
}

func main() {
	s := stack{}
	s.size = 5
	s.push(1)
	s.push(2)
	s.push(3)
	s.push(4)
	s.push(5)
	s.pop(2)

	fmt.Println(s.data)
	fmt.Println(s.top)

}

func (s *stack) push(x int) {
	if !s.isFull() {
		s.data = append(s.data, x)
		s.top++
	} else {
		fmt.Println("overflow")
	}
}

func (s *stack) isFull() bool {
	if s.top == s.size {
		return true
	} else {
		return false
	}
}

func (s *stack) pop(x int) {
	if !s.isEmpty() {
		s.data = append(s.data[:s.top-1], s.data[s.top-1+1:]...)
		s.top--
	} else {
		fmt.Println("underflow")
	}
}

func (s *stack) isEmpty() bool {
	if s.top == 0 {
		return true
	} else {
		return false
	}
}
