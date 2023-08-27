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
	result := convertLinkedListValsToSlice(n)
	want := []int{5, 6}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func convertLinkedListValsToSlice(n *SinglyLinkedListNode) []int {
	result := []int{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	return result
}

func TestShouldInsertAfterInMiddleOfList(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.Next = NewSinglyLinkedListNode(7)
	n.InsertAfter(6)
	result := convertLinkedListValsToSlice(n)
	want := []int{5, 6, 7}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldRemoveSinglyLinkedListWithOneNode(t *testing.T) {
	n := NewSinglyLinkedListNode(5)
	n.InsertAfter(6)
	n.RemoveNext()
	result := convertLinkedListValsToSlice(n)
	want := []int{5}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}

}
