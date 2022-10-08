package linkedlist

import (
	"fmt"
	"strconv"
)

func PrintHello() {
	fmt.Println("Hello, Modules! This is linkedlist speaking!")
}

type Node struct {
	next *Node
	data int
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Add(i int) {
	n := Node{nil, i}
	l.add(n)
}

func (l *LinkedList) add(new Node) {
	// edge case: empty list
	if l.head == nil {
		l.head = &new
		return
	}

	// traverse through list
	var tail *Node
	for tail = l.head; tail.next != nil; tail = tail.next {
	}
	tail.next = &new
}

func (l LinkedList) String() string {
	var s string = ""
	for n := l.head; n != nil; n = n.next {
		s += (strconv.Itoa(n.data) + " ")
	}
	return s
}