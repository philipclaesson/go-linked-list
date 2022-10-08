package main

import (
	"fmt"
	"go-linked-list/linkedlist"
)

func main() {
	var list linkedlist.LinkedList = linkedlist.LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	list.Remove(1)
	fmt.Println(list)
}
