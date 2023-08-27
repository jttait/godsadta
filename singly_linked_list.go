package godsa

type SinglyLinkedListNode struct {
	Val  int
	Next *SinglyLinkedListNode
}

func NewSinglyLinkedListNode(i int) *SinglyLinkedListNode {
	n := SinglyLinkedListNode{}
	n.Val = i
	return &n
}

func (n *SinglyLinkedListNode) InsertAfter(i int) {
	if n.Next == nil {
		n.Next = NewSinglyLinkedListNode(i)
		return
	}
	temp := n.Next
	n.Next = NewSinglyLinkedListNode(i)
	n.Next.Next = temp
}

func (n *SinglyLinkedListNode) TraverseUntilValEquals(i int) *SinglyLinkedListNode {
	for n.Val != i {
		n = n.Next
	}
	return n
}

func (n *SinglyLinkedListNode) TraverseUntilNextValEquals(i int) *SinglyLinkedListNode {
	for n.Next.Val != i {
		n = n.Next
	}
	return n
}

func (n *SinglyLinkedListNode) TraverseToTail() *SinglyLinkedListNode {
	for n.Next != nil {
		n = n.Next
	}
	return n
}

func (n *SinglyLinkedListNode) DeleteNext() {
	n.Next = n.Next.Next
}
