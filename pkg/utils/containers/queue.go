package containers

// Queue is a generic queue implementation using a circular buffer.
// All operations are O(1) amortized.
type Queue[T any] struct {
	items []T
	head  int
	tail  int
	size  int
}

// NewQueue creates a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		items: make([]T, 1),
		head:  0,
		tail:  0,
		size:  0,
	}
}

// Enqueue adds an item to the back of the queue.
// Time complexity: O(1) amortized.
func (q *Queue[T]) Enqueue(item T) {
	if q.size == len(q.items) {
		q.resize()
	}
	q.items[q.tail] = item
	q.tail = (q.tail + 1) % len(q.items)
	q.size++
}

// Dequeue removes and returns the front item from the queue.
// Returns (zero value, false) if the queue is empty.
// Time complexity: O(1) amortized.
func (q *Queue[T]) Dequeue() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	item := q.items[q.head]
	var zero T
	q.items[q.head] = zero // Clear reference for GC
	q.head = (q.head + 1) % len(q.items)
	q.size--
	if q.size > 0 && q.size == len(q.items)/4 {
		q.shrink()
	}
	return item, true
}

// Front returns the front item without removing it.
// Returns (zero value, false) if the queue is empty.
// Time complexity: O(1).
func (q *Queue[T]) Front() (T, bool) {
	if q.IsEmpty() {
		var zero T
		return zero, false
	}
	return q.items[q.head], true
}

// Size returns the number of items in the queue.
func (q *Queue[T]) Size() int {
	return q.size
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.size == 0
}

// Clear removes all items from the queue.
func (q *Queue[T]) Clear() {
	q.items = make([]T, 1)
	q.head = 0
	q.tail = 0
	q.size = 0
}

func (q *Queue[T]) resize() {
	newItems := make([]T, 2*len(q.items))
	if q.head < q.tail {
		copy(newItems, q.items[q.head:q.tail])
	} else {
		copy(newItems, q.items[q.head:])
		copy(newItems[len(q.items)-q.head:], q.items[:q.tail])
	}
	q.items = newItems
	q.head = 0
	q.tail = q.size
}

func (q *Queue[T]) shrink() {
	newItems := make([]T, len(q.items)/2)
	if q.head < q.tail {
		copy(newItems, q.items[q.head:q.tail])
	} else {
		copy(newItems, q.items[q.head:])
		copy(newItems[len(q.items)-q.head:], q.items[:q.tail])
	}
	q.items = newItems
	q.head = 0
	q.tail = q.size
}
