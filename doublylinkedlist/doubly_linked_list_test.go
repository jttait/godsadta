package doublylinkedlist

import (
	"testing"

	"github.com/jttait/godsa/slice"
)

func TestShouldHaveCorrectValForNewlyInstantiateDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedListNode[int](1)
	result := d.Val
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func convertDoublyLinkedListValsToSlice[T any](n *DoublyLinkedListNode[T]) []T {
	result := []T{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	return result
}

func TestShouldInsertAfterTail(t *testing.T) {
	d := NewDoublyLinkedListNode[int](5)
	d.InsertAfter(6)
	result := convertDoublyLinkedListValsToSlice(d)
	want := []int{5, 6}
	if !slice.AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldInsertAfterInMiddleOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode[int](5)
	n.Next = NewDoublyLinkedListNode[int](7)
	n.InsertAfter(6)
	result := convertDoublyLinkedListValsToSlice(n)
	want := []int{5, 6, 7}
	if !slice.AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldRemoveInMiddleOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode[int](5)
	head := n
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.Next
	n.Remove()
	result := convertDoublyLinkedListValsToSlice(head)
	want := []int{5, 7, 8}
	if !slice.AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldRemoveHeadOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode[int](5)
	n.InsertAfter(6)
	head := n.Next
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n.Remove()
	result := convertDoublyLinkedListValsToSlice(head)
	want := []int{6, 7, 8}
	if !slice.AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldRemoveTailOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode[int](5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n.Next.Next.Next.Remove()
	result := convertDoublyLinkedListValsToSlice(n)
	want := []int{5, 6, 7}
	if !slice.AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}
