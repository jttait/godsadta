package multiset

type MapMultiset[T comparable] struct {
	m map[T]int
}

func NewMapMultiset[T comparable]() *MapMultiset[T] {
	multiset := MapMultiset[T]{}
	multiset.m = make(map[T]int)
	return &multiset
}

func (multiset *MapMultiset[T]) Add(i T) {
	value, ok := multiset.m[i]
	if ok {
		multiset.m[i] = value + 1
	} else {
		multiset.m[i] = 1
	}
}

func (multiset *MapMultiset[T]) Remove(i T) bool {
	value, ok := multiset.m[i]
	if !ok {
		return false
	}
	multiset.m[i] = value - 1
	return true
}

func (multiset *MapMultiset[T]) Size() int {
	result := 0
	for _, val := range multiset.m {
		result += val
	}
	return result
}

func (multiset *MapMultiset[T]) Count(i T) int {
	value, ok := multiset.m[i]
	if !ok {
		return 0
	}
	return value
}
