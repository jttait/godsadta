package doublylinkedlist

import (
	"testing"
)

func TestShouldHaveCorrectValForNewlyInstantiateDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedListNode[int](1)
	result := d.Val
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
