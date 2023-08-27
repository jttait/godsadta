package godsa

type DoublyLinkedListNode struct {
	Val  int
	Prev *DoublyLinkedListNode
	Next *DoublyLinkedListNode
}

func NewDoublyLinkedListNode(i int) *DoublyLinkedListNode {
	d := DoublyLinkedListNode{}
	d.Val = i
	return &d
}

func (n *DoublyLinkedListNode) InsertAfter(i int) {
	if n.Next == nil {
		insertedNode := NewDoublyLinkedListNode(i)
		insertedNode.Prev = n
		n.Next = insertedNode
		return
	}
	temp := n.Next
	insertedNode := NewDoublyLinkedListNode(i)
	insertedNode.Prev = n
	insertedNode.Next = temp
	n.Next = insertedNode
}

func (n *DoublyLinkedListNode) TraverseForwardUntilValEquals(i int) *DoublyLinkedListNode {
	for n.Val != i {
		n = n.Next
	}
	return n
}

func (n *DoublyLinkedListNode) TraverseForwardUntilNextValEquals(i int) *DoublyLinkedListNode {
	for n.Next.Val != i {
		n = n.Next
	}
	return n
}

func (n *DoublyLinkedListNode) TraverseForwardToTail() *DoublyLinkedListNode {
	for n.Next != nil {
		n = n.Next
	}
	return n
}

func (n *DoublyLinkedListNode) TraverseBackwardUntilValEquals(i int) *DoublyLinkedListNode {
	for n.Val != i {
		n = n.Prev
	}
	return n
}

func (n *DoublyLinkedListNode) TraverseBackwardUntilPrevValEquals(i int) *DoublyLinkedListNode {
	for n.Prev.Val != i {
		n = n.Prev
	}
	return n
}

func (n *DoublyLinkedListNode) TraverseBackwardToHead() *DoublyLinkedListNode {
	for n.Prev != nil {
		n = n.Prev
	}
	return n
}

func (n *DoublyLinkedListNode) Delete() {
	prev := n.Prev
	next := n.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
}
