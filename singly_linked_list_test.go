package godsa

import (
	"testing"
)

func TestShouldBeCorrectValForNewlyInstantiatedSinglyLinkedListNode(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	want := 5
	result := n.Val
	if want != result {
		t.Fatalf("Want %v. Got %v.\n", result, want)
	}
}

func TestShouldBeNilNextForNewlyInstantiatedSinglyLinkedListNode(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	result := n.Next
	if result != nil {
		t.Fatalf("Want %v. Got %v.\n", nil, result)
	}
}

func TestShouldInsertAfterAtTail(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.InsertAfter(6)
	result := []int{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	want := []int{5, 6}
	if len(result) != len(want) {
		t.Fatalf("Linked lists are not equal.")
	}
	for i := 0; i < len(result); i++ {
		if result[i] != want[i] {
			t.Fatalf("Linked lists are not equal.")
		}
	}
}

func TestShouldInsertAfterInMiddleOfList(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.Next = NewSinglyLinkedListNode(7)
	n.InsertAfter(6)
	result := []int{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	want := []int{5, 6, 7}
	if len(result) != len(want) {
		t.Fatalf("Linked lists are not equal.")
	}
	for i := 0; i < len(result); i++ {
		if result[i] != want[i] {
			t.Fatalf("Linked lists are not equal.")
		}
	}
}

func TestShouldTraverseListUntilValFound(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.TraverseUntilValEquals(7)
	result := n.Val
	want := 7
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldTraverseListUntilNextValFound(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.TraverseUntilNextValEquals(8)
	result := n.Val
	want := 7
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}

}

func TestShouldTraverseToTailOfSinglyLinkedList(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.TraverseToTail()
	result := n.Val
	want := 8
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}

}

func TestShouldDeleteSinglyLinkedListWithOneNode(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.InsertAfter(6)
	n.DeleteNext()
	result := []int{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	want := []int{5}
	if len(result) != len(want) {
		t.Fatalf("Linked lists are not equal.")
	}
	for i := 0; i < len(result); i++ {
		if result[i] != want[i] {
			t.Fatalf("Linked lists are not equal.")
		}
	}

}
