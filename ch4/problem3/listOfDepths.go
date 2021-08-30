package problem3

import (
	"crackingthecode/ch4/problem2"
)

type linkedListNode struct {
	val  int
	next *linkedListNode
}

type linkedList struct {
	head *linkedListNode
	tail *linkedListNode
}

func NewLinkedList() *linkedList {
	return &linkedList{head: nil, tail: nil}
}

func (ll *linkedList) AddNode(val int) {
	n := &linkedListNode{val: val, next: nil}
	if ll.head == nil {
		ll.head = n
		ll.tail = ll.head
		return
	}

	// Method 1 to add a node:
	// var trailingPtr *linkedListNode
	// currentPtr := ll.head
	// for currentPtr != nil {
	// 	trailingPtr = currentPtr
	// 	currentPtr = currentPtr.next
	// }
	// trailingPtr.next = n
	// ll.tail = trailingPtr.next
	// Method 2 to add a node:
	ll.tail.next = n
	ll.tail = ll.tail.next
}

// Create an array of linked lists, with index i pointing to the level
// Use preorder traversal, makes sense, we are adding as we go down the tree
func ListOfDepths(root *problem2.Node, arr *[]*linkedList, level int) {
	if root == nil {
		return
	}
	var ll *linkedList
	if len((*arr)) == level {
		ll = NewLinkedList()
		(*arr) = append((*arr), ll)
	} else {
		ll = (*arr)[level]
	}
	ll.AddNode(root.Value)
	ListOfDepths(root.Left, arr, level+1)
	ListOfDepths(root.Right, arr, level+1)
}
