// Package doublylinkedlist provides the doubly-linked list data structure and associated methods
package doublylinkedlist

// DoublyLinkedListNode is a data structure that contains a value and pointers to the previous and
// next nodes in the list.
type DoublyLinkedListNode struct {
	Val  int
	Prev *DoublyLinkedListNode
	Next *DoublyLinkedListNode
}

// NewDoublyLinkedListNode instantiates a doubly-linked list node and returns a pointer to it.
func NewDoublyLinkedListNode(i int) *DoublyLinkedListNode {
	d := DoublyLinkedListNode{}
	d.Val = i
	return &d
}

// InsertAfter inserts a new node after the current one. If there is already a node there then the
// previous and next pointers of all three nodes will be updated as required.
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

// Remove removes the current node and updates the pointers of the previous and next nodes as
// required.
func (n *DoublyLinkedListNode) Remove() {
	prev := n.Prev
	next := n.Next
	if prev != nil {
		prev.Next = next
	}
	if next != nil {
		next.Prev = prev
	}
}
