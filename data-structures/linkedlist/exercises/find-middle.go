package exercises

import "github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"

// FindMiddle returns the element in the middle of the linked list
// If there are even nodes, then there would be two middle nodes, should be printed only
// the second middle element
func FindMiddle(head *linkedlist.Node) int {

	slowNode, fastNode := head, head

	if head != nil {
		for fastNode != nil && fastNode.Next != nil {
			fastNode = fastNode.Next.Next
			slowNode = slowNode.Next
		}
	}
	return slowNode.Data
}
