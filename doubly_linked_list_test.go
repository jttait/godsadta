package godsa

import (
	"testing"
)

func TestShouldHaveCorrectValForNewlyInstantiateDoublyLinkedList(t *testing.T) {
	d := NewDoublyLinkedListNode(1)
	result := d.Val
	want := 1
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func convertDoublyLinkedListValsToSlice(n *DoublyLinkedListNode) []int {
	result := []int{}
	for n != nil {
		result = append(result, n.Val)
		n = n.Next
	}
	return result
}

func TestShouldInsertAfterTail(t *testing.T) {
	d := NewDoublyLinkedListNode(5)
	d.InsertAfter(6)
	result := convertDoublyLinkedListValsToSlice(d)
	want := []int{5, 6}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldInsertAfterInMiddleOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.Next = NewDoublyLinkedListNode(7)
	n.InsertAfter(6)
	result := convertDoublyLinkedListValsToSlice(n)
	want := []int{5, 6, 7}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldTraverseListForwardUntilValFound(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.TraverseForwardUntilValEquals(7)
	result := n.Val
	want := 7
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldTraverseListForwardUntilNextValFound(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.TraverseForwardUntilNextValEquals(8)
	result := n.Val
	want := 7
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}

}

func TestShouldTraverseForwardToTailOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.TraverseForwardToTail()
	result := n.Val
	want := 8
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldTraverseListBackwardUntilValFound(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.Next.Next.Next
	n = n.TraverseBackwardUntilValEquals(6)
	result := n.Val
	want := 6
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldTraverseListBackwardUntilPrevValFound(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.Next.Next.Next
	n = n.TraverseBackwardUntilPrevValEquals(5)
	result := n.Val
	want := 6
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldTraverseBackwardToHeadOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.Next.Next.Next
	n = n.TraverseBackwardToHead()
	result := n.Val
	want := 5
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldDeleteInMiddleOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	head := n
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n = n.Next
	n.Delete()
	result := convertDoublyLinkedListValsToSlice(head)
	want := []int{5, 7, 8}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldDeleteHeadOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	head := n.Next
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n.Delete()
	result := convertDoublyLinkedListValsToSlice(head)
	want := []int{6, 7, 8}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}

func TestShouldDeleteTailOfDoublyLinkedList(t *testing.T) {
	n := NewDoublyLinkedListNode(5)
	n.InsertAfter(6)
	n.Next.InsertAfter(7)
	n.Next.Next.InsertAfter(8)
	n.Next.Next.Next.Delete()
	result := convertDoublyLinkedListValsToSlice(n)
	want := []int{5, 6, 7}
	if !AreSlicesEqual(result, want) {
		t.Fatalf("Linked lists are not equal.")
	}
}
