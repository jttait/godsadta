package singlylinkedlist

import (
	"testing"
)

func TestShouldBeCorrectValForNewlyInstantiatedSinglyLinkedListNode(t *testing.T) {
	n := NewSinglyLinkedListNode[int](5)
	want := 5
	result := n.Val
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeNilNextForNewlyInstantiatedSinglyLinkedListNode(t *testing.T) {
	n := NewSinglyLinkedListNode[int](5)
	result := n.Next
	if result != nil {
		t.Fatalf("Want %v. Got %v.\n", nil, result)
	}
}
