package main

import "fmt"

// implement using slice
// for better performance use linked list -> dequeue operation
type queue struct {
	items []int
}

// Push will add a value at the end
func (s *queue) enqueue(i int) {
	s.items = append(s.items, i)
}

// Pop will remove a value at the end
// and returns the removed value
func (s *queue) dequeue() int {
	toRemove := s.items[0]
	s.items = s.items[1:]
	return toRemove
}

func main() {
	myQueue := queue{}
	myQueue.enqueue(1)
	myQueue.enqueue(2)
	myQueue.enqueue(3)
	fmt.Println(myQueue)
	myQueue.dequeue()
	fmt.Println(myQueue)
}
