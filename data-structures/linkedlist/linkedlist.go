package linkedlist

import "fmt"

// Node represents a linked list's node
type Node struct {
	Data int
	Next *Node
}

// NewNode creates a pointer to a new Node
func NewNode(data int) *Node {
	newNode := &Node{}
	newNode.Data = data

	return newNode
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

// AddToHead adds a new node into linked list head
func (ll *LinkedList) AddToHead(data int) {

	newNode := NewNode(data)

	newNode.Next = ll.Head
	ll.Head = newNode

}

// AddAfter adds a new node after the given prevNode
func (ll *LinkedList) AddAfter(prevNode *Node, data int) {
	if prevNode == nil {
		fmt.Println("The given previous node cannot be null")
	}

	newNode := NewNode(data)

	newNode.Next = prevNode.Next

	prevNode.Next = newNode
}

// AddToEnd adds a new node into do end of the linked list
func (ll *LinkedList) AddToEnd(data int) {
	newNode := NewNode(data)

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	last := ll.Head
	for last.Next != nil {
		last = last.Next
	}

	last.Next = newNode
}

// DeleteNode deletes the first occurrence of a key in linked list
func (ll *LinkedList) DeleteNode(key int) {
	temp, prev := ll.Head, &Node{}

	if temp != nil && temp.Data == key {
		ll.Head = temp.Next
		return
	}

	for temp != nil && temp.Data != key {
		prev = temp
		temp = temp.Next
	}

	if temp == nil {
		return
	}

	prev.Next = temp.Next

}

// DeleteAt deletes a node at specific position
func (ll *LinkedList) DeleteAt(position int) {
	if ll.Head == nil {
		return
	}

	node := ll.Head

	if position == 0 {
		ll.Head = node.Next
	}

	for i := 0; node != nil && i < position-1; i++ {
		node = node.Next
	}

	if node == nil || node.Next == nil {
		return
	}

	next := node.Next.Next

	node.Next = next
}

// DeleteList deletes the entire Linked List
func (ll *LinkedList) DeleteList() {
	ll.Head = nil
}
