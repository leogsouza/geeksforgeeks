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
	fmt.Printf("Total elements into Linked List %d\n", ll.Size())
	fmt.Println("Linked List after added 4 items to head")
	ll.PrintList()
	fmt.Println()
	fmt.Println("Is 3 in the list? ", ll.Search(3))

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
	fmt.Println()
	fmt.Printf("Total elements into Linked List %d\n", ll.SizeRecursive())
	fmt.Println("Is 1 in the list? ", ll.SearchRecursive(ll.Head, 1))
}
