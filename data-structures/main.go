package main

import (
	"github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"
)

func main() {
	ll := &linkedlist.LinkedList{}

	ll.Head = &linkedlist.Node{Data: 1, Next: nil}
	second := &linkedlist.Node{Data: 2, Next: nil}
	third := &linkedlist.Node{Data: 3, Next: nil}

	ll.Head.Next = second

	second.Next = third

	ll.PrintList()
}
