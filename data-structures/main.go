package main

import (
	"fmt"

	"github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"
)

func main() {
	ll := &linkedlist.LinkedList{}
	ll.AddToHead(0)
	ll.AddToHead(1)
	ll.AddToHead(2)
	ll.AddToHead(3)
	fmt.Println("Linked List after added 4 items to head")
	ll.PrintList()
	fmt.Println()

	ll.AddToEnd(5)
	fmt.Println("Linked List after added a node to the end")
	ll.PrintList()
	fmt.Println()

	ll.DeleteNode(2)
	fmt.Println("Linked list after delete a node by data")
	ll.PrintList()
	fmt.Println()

	ll.DeleteAt(1)
	fmt.Println("Linked list after delete a node by position")
	ll.PrintList()

}
