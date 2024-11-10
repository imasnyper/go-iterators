package cyclic_iterator

// CyclicIterator holds the slice and the current position.
type CyclicIterator[T any] struct {
	data []T
	pos  int
}

// NewCyclicIterator initializes the iterator with a slice.
func NewCyclicIterator[T any](data []T) *CyclicIterator[T] {
	return &CyclicIterator[T]{data: data, pos: -1}
}

// Next moves to the next element, wrapping around to the start if needed.
func (it *CyclicIterator[T]) Next() T {
	it.pos = (it.pos + 1) % len(it.data)
	return it.data[it.pos]
}

// Reset resets the iterator back to the start.
func (it *CyclicIterator[T]) Reset() {
	it.pos = -1
}
