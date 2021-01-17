package exercises

import (
	"fmt"

	"github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"
)

// DetectLoop detects the loop into a linked list
func DetectLoop(head *linkedlist.Node) bool {

	current := head.Next

	for current != nil {
		fmt.Printf("%v - %v", current, head)
		if current == head {
			return true
		}
		current = current.Next
	}
	return false
}
