Data structures and algorithms implemented in Go.

# Data structures

## LinkedList[T any]

Time:

| Implementation | SinglyLinkedList | DoublyLinkedList |
| -------------- | ---------------- | ---------------- |
| InsertFront    | O(1)             | O(1)             |
| InsertLast     | O(n)             | O(1)             |
| RemoveFront    | O(1)             | O(1)             |
| RemoveLast     | O(n)             | O(1)             |
| PeekFront      | O(1)             | O(1)             |
| PeekLast       | O(n)             | O(n)             |
| Size           | O(n)             | O(n)             |

Space:

## MinHeap[T cmp.Ordered]

## MaxHeap[T cmp.Ordered]

## Trie

- Binary tree

# Abstract data types

## Set[T comparable]

Time: 

| Implementation | MapSet |
| -------------- | ------ |
| Size           | O(1)   |
| Insert         | O(1)   |
| Remove         | O(1)   |
| Contains       | O(1)   |
| Equals         | O(n)   |
| Iter           | O(n)   |
| Union          | O(m.n) |
| Intersection   | O(n)   |
| Difference     | O(m.n) |

## List[T any]

Time:

| Implementation | ArrayList | LinkedListList |             
| -------------- | --------- | -------------- |
| Size           | O(1)      | O(n)           |
| Prepend        | O(n)      | O(1)           |
| Append         | O(1),O(n) | O(1)           |
| Remove         | O(n)      | O(n)           |
| Insert         | O(n)      | O(n)           |
| Get            | O(1)      | O(n)           |

- Double-ended queue
- Graph
- Priority queue
- Queue
- Stack
