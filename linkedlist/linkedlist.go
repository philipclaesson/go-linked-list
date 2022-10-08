package linkedlist

import (
	"strconv"
)

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

func (l *LinkedList) Remove(i int) bool {
	// a pointer to the next pointer of the previous node
	var prev **Node = &l.head
	for n := l.head; n != nil; n = n.next {
		if n.data == i {
			// remove n by changing the value of the next pointer of the previous node
			*prev = n.next
			return true
		}
		prev = &n.next
	}
	return false
}

func (l *LinkedList) add(new Node) {
	// tailPointer a pointer to the last known next pointer
	var tailPointer **Node

	// traverse through list until we find the pointer that is pointing to a nil-pointer
	for tailPointer = &l.head; *tailPointer != nil; tailPointer = &(*tailPointer).next {
	}

	// append to list by making the nil-pointer point at our new node
	*tailPointer = &new
}

func (l LinkedList) String() string {
	var s string = ""
	for n := l.head; n != nil; n = n.next {
		s += (strconv.Itoa(n.data) + " ")
	}
	return s
}
