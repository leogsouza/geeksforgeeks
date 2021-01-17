package exercises

import "github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"

// TotalElementRepeated returns the total of elements is repeated into a linked list
func TotalElementRepeated(head *linkedlist.Node, num int) int {

	current := head
	c := 0
	for current != nil {
		if current.Data == num {
			c++
		}
		current = current.Next
	}
	return c
}
