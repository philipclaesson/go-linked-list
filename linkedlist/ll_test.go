package linkedlist

import (
	"testing"
)

func TestAddElements(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Add(1)
	if list.head.data != 1 {
		t.Log("did not add an object with value 1")
		t.Fail()
	}

	if list.head.next != nil {
		t.Log("added one object, but there are more than one objects")
		t.Fail()
	}
}

func TestLength(t *testing.T) {
	var list LinkedList = LinkedList{}
	if list.Length() != 0 {
		t.Log("expected length 0, got length", list.Length())
		t.Fail()
	}

	list.Add(1)
	if list.Length() != 1 {
		t.Log("expected length 1, got length", list.Length())
		t.Fail()
	}

	list.Remove(1)
	if list.Length() != 0 {
		t.Log("expected length 0, got length", list.Length())
		t.Fail()
	}
}

func TestRemoveElement(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Add(1)
	list.Add(2)
	if !list.Remove(2) {
		t.Log("failed removing element that is in list")
		t.Fail()
	}
	if list.Length() != 1 {
		t.Log("Remove() succeeded but element is still in list")
		t.Fail()
	}
}

func TestToString(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Add(1)
	list.Add(2)
	list.Add(3)
	if list.String() != "1 2 3 " {
		t.Log("expected String to return `1 2 3 `, got ", list.String())
		t.Fail()
	}
}

func TestStringData(t *testing.T) {
	var list LinkedList = LinkedList{}
	list.Add("Hello")
	if list.String() != "Hello " {
		t.Log("Could not add a string to list")
		t.Fail()
	}
}
