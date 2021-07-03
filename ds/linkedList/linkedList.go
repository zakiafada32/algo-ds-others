package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length += 1
}

func (l linkedList) printListData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length -= 1
	}
	fmt.Println("")
}

func (l *linkedList) deleteWithValue(value int) {
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		previousToDelete = previousToDelete.next
	}

	previousToDelete.next = previousToDelete.next.next
	l.length -= 1
}

func main() {
	myList := linkedList{}
	node1 := &node{data: 4}
	node2 := &node{data: 10}
	myList.prepend(node1)
	myList.prepend(node2)
	myList.printListData()
	myList.deleteWithValue(4)
	myList.printListData()
}
