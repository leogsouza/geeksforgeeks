package exercises

import (
	"testing"

	"github.com/leogsouza/geeksforgeeks/data-structures/linkedlist"
)

func TestDetectLoop(t *testing.T) {

	ll := &linkedlist.LinkedList{}
	ll.AddToEnd(10)
	ll.AddToEnd(20)
	ll.AddToEnd(30)
	current := ll.Head
	for current != nil {
		if current.Next == nil {
			current.Next = ll.Head
			break
		}
		current = current.Next
	}
	expected := true
	got := DetectLoop(ll.Head)

	if got != expected {
		t.Fatalf("got %t expected %t", got, expected)
	}

}
