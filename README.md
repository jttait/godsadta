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
| Map            | O(n)      | O(n)           |
| Filter         | O(n)      | O(n)           |

## Stack[T any]

| Implementation | LinkedListStack |
| -------------- | --------------- |
| Size           | O(n)            |
| Push           | O(1)            |
| Pop            | O(1)            |
| Peek           | O(1)            |

## Queue[T any]

| Implementation | QueueLinkedList |
| -------------- | --------------- |
| Size           | O(n)            |
| Insert         | O(1)            |
| Remove         | O(1)            |

## DoubleEndedQueue[T any]

| Implementation | DoubleEndedQueueLinkedList |
| -------------- | -------------------------- |
| Size           | O(n)                       |
| InsertFront    | O(1)                       |
| InsertLast     | O(1)                       |
| RemoveFront    | O(1)                       |
| RemoveLast     | O(1)                       |

- Graph
- Priority queue
