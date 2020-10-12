package cache

import (
	"fmt"
)

// Node ...
type Node struct {
	Data int
	Next *Node
}

// NewNode ...
func NewNode(data int, next *Node) *Node {
	return &Node{
		Data: data,
		Next: next,
	}
}

// LinkedList ...
type LinkedList struct {
	Head *Node
	Size int
}

// NewLinkedList creates a new object of Linkedlist
func NewLinkedList(head *Node) *LinkedList {
	return &LinkedList{
		Size: 0,
		Head: head,
	}
}

// AddLast appends node to the end
func (ll *LinkedList) AddLast(node *Node) {
	if ll.Head != nil {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = node
	} else {
		ll.Head = node
	}

	ll.Size++
}

// AddFirst appends node to the start
func (ll *LinkedList) AddFirst(node *Node) {
	current := ll.Head
	ll.Head = node
	ll.Head.Next = current
	ll.Size++
}

// InsertAtIndex inserts new node at particular index
func (ll *LinkedList) InsertAtIndex(node *Node, index int) {
	if index < 0 || index > ll.Size {
		return
	}

	if index == 0 {
		ll.AddFirst(node)
		return
	}

	count := 0

	var previous *Node
	current := ll.Head

	for count < index {
		previous = current
		current = current.Next

		count++
	}

	node.Next = current
	previous.Next = node
	ll.Size++
}

// GetAtIndex gets the node at particular index
func (ll *LinkedList) GetAtIndex(index int) *Node {
	current := ll.Head
	count := 0
	for current != nil {
		if index == count {
			return current
		}

		count++
		current = current.Next
	}

	return nil
}

// RemoveAtIndex removes node at particular index
func (ll *LinkedList) RemoveAtIndex(index int) {
	if index < 0 || index > ll.Size {
		return
	}

	var previous *Node
	current := ll.Head

	count := 0

	for count < index {
		previous = current
		current = current.Next
		count++
	}

	if index == 0 {
		ll.Head = ll.Head.Next
	} else {
		previous.Next = current.Next
	}

}

// Print displays the current linkedlist
func (ll *LinkedList) Print() {
	current := ll.Head

	for current != nil {
		fmt.Println(current)
		current = current.Next
	}
}
