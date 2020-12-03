package linkedlist

import "fmt"

// Node represents a linked list's node
type Node struct {
	Data int
	Next *Node
}

// LinkedList is a linked list data structure
type LinkedList struct {
	Head *Node
}

// PrintList prints out all data into linked list
func (ll *LinkedList) PrintList() {
	n := ll.Head
	for n != nil {
		fmt.Printf("%d ", n.Data)
		n = n.Next
	}
}
