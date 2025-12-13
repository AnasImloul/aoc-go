package containers

// Deque is a generic double-ended queue implementation using a circular buffer.
// All operations are O(1) amortized.
type Deque[T any] struct {
	items []T
	head  int
	tail  int
	size  int
}

// NewDeque creates a new deque.
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		items: make([]T, 1),
		head:  0,
		tail:  0,
		size:  0,
	}
}

// PushFront adds an item to the front of the deque.
// Time complexity: O(1) amortized.
func (d *Deque[T]) PushFront(item T) {
	if d.size == len(d.items) {
		d.resize()
	}
	d.head = (d.head - 1 + len(d.items)) % len(d.items)
	d.items[d.head] = item
	d.size++
}

// PushBack adds an item to the back of the deque.
// Time complexity: O(1) amortized.
func (d *Deque[T]) PushBack(item T) {
	if d.size == len(d.items) {
		d.resize()
	}
	d.items[d.tail] = item
	d.tail = (d.tail + 1) % len(d.items)
	d.size++
}

// PopFront removes and returns the front item from the deque.
// Returns (zero value, false) if the deque is empty.
// Time complexity: O(1) amortized.
func (d *Deque[T]) PopFront() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	item := d.items[d.head]
	var zero T
	d.items[d.head] = zero // Clear reference for GC
	d.head = (d.head + 1) % len(d.items)
	d.size--
	if d.size > 0 && d.size == len(d.items)/4 {
		d.shrink()
	}
	return item, true
}

// PopBack removes and returns the back item from the deque.
// Returns (zero value, false) if the deque is empty.
// Time complexity: O(1) amortized.
func (d *Deque[T]) PopBack() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	d.tail = (d.tail - 1 + len(d.items)) % len(d.items)
	item := d.items[d.tail]
	var zero T
	d.items[d.tail] = zero // Clear reference for GC
	d.size--
	if d.size > 0 && d.size == len(d.items)/4 {
		d.shrink()
	}
	return item, true
}

// Front returns the front item without removing it.
// Returns (zero value, false) if the deque is empty.
// Time complexity: O(1).
func (d *Deque[T]) Front() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	return d.items[d.head], true
}

// Back returns the back item without removing it.
// Returns (zero value, false) if the deque is empty.
// Time complexity: O(1).
func (d *Deque[T]) Back() (T, bool) {
	if d.IsEmpty() {
		var zero T
		return zero, false
	}
	idx := (d.tail - 1 + len(d.items)) % len(d.items)
	return d.items[idx], true
}

// Size returns the number of items in the deque.
func (d *Deque[T]) Size() int {
	return d.size
}

// IsEmpty returns true if the deque is empty.
func (d *Deque[T]) IsEmpty() bool {
	return d.size == 0
}

// Clear removes all items from the deque.
func (d *Deque[T]) Clear() {
	d.items = make([]T, 1)
	d.head = 0
	d.tail = 0
	d.size = 0
}

func (d *Deque[T]) resize() {
	newItems := make([]T, 2*len(d.items))
	if d.head < d.tail {
		copy(newItems, d.items[d.head:d.tail])
	} else {
		copy(newItems, d.items[d.head:])
		copy(newItems[len(d.items)-d.head:], d.items[:d.tail])
	}
	d.items = newItems
	d.head = 0
	d.tail = d.size
}

func (d *Deque[T]) shrink() {
	newItems := make([]T, len(d.items)/2)
	if d.head < d.tail {
		copy(newItems, d.items[d.head:d.tail])
	} else {
		copy(newItems, d.items[d.head:])
		copy(newItems[len(d.items)-d.head:], d.items[:d.tail])
	}
	d.items = newItems
	d.head = 0
	d.tail = d.size
}
