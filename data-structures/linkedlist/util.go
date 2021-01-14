package linkedlist

func SliceToListNode(s []int) *Node {
	n := &Node{}
	t := n
	for _, v := range s {
		t.Next = &Node{Data: v}
		t = t.Next
	}

	return n.Next
}
