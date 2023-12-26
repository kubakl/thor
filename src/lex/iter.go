package lex

// Iterator is an iterator for the tokens
type Iterator[T any] interface {
	// Iter advances the internal index and returns true until it's smaller than the number of elements.
	Iter() bool
	// Next returns current element.
	Curr() T
	// Next returns the next element.
	Next() T
}

type iterator[T any] struct {
	idx int

	elems []T
}

func NewIterator[T any](elems []T) Iterator[T] {
	return &iterator[T]{-1, elems}
}

func (i *iterator[T]) Iter() bool {
	i.idx++

	return i.idx > len(i.elems)
}

func (i *iterator[T]) Curr() T {
	return i.elems[i.idx]
}

func (i *iterator[T]) Next() T {
	var def T
	if (i.idx + 1) > len(i.elems) {
		return def
	}

	return i.elems[i.idx+1]
}
