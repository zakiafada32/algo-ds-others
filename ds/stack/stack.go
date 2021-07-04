package main

import "fmt"

// implement using slice
type stack struct {
	items []int
}

// Push will add a value at the end
func (s *stack) push(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and returns the removed value
func (s *stack) pop() int {
	l := len(s.items) - 1
	toRemove := s.items[l]
	s.items = s.items[:l]
	return toRemove
}

func main() {
	myStack := stack{}
	myStack.push(1)
	myStack.push(2)
	myStack.push(3)
	fmt.Println(myStack)
	myStack.pop()
	fmt.Println(myStack)
}
