package containers

import "container/heap"

// PriorityQueueItem represents an item in the priority queue.
type PriorityQueueItem[T any] struct {
	Value    T
	Priority int
	index    int
}

// PriorityQueue is a generic priority queue implementation using a min-heap by default.
type PriorityQueue[T any] struct {
	items    []*PriorityQueueItem[T]
	lessFunc func(a, b *PriorityQueueItem[T]) bool
}

// NewPriorityQueue creates a new priority queue with a default min-heap comparator.
func NewPriorityQueue[T any]() *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items: make([]*PriorityQueueItem[T], 0),
		lessFunc: func(a, b *PriorityQueueItem[T]) bool {
			return a.Priority < b.Priority
		},
	}
}

// NewPriorityQueueWithComparator creates a new priority queue with a custom comparator.
func NewPriorityQueueWithComparator[T any](less func(a, b *PriorityQueueItem[T]) bool) *PriorityQueue[T] {
	return &PriorityQueue[T]{
		items:    make([]*PriorityQueueItem[T], 0),
		lessFunc: less,
	}
}

// PushItem adds an item to the priority queue.
func (pq *PriorityQueue[T]) PushItem(value T, priority int) {
	item := &PriorityQueueItem[T]{
		Value:    value,
		Priority: priority,
		index:    len(pq.items),
	}
	heap.Push(pq, item)
}

// PopItem removes and returns the top item from the priority queue.
// Returns (zero value, false) if the queue is empty.
func (pq *PriorityQueue[T]) PopItem() (T, bool) {
	if pq.IsEmpty() {
		var zero T
		return zero, false
	}
	item := heap.Pop(pq).(*PriorityQueueItem[T])
	return item.Value, true
}

// Top returns the top item without removing it.
// Returns (zero value, false) if the queue is empty.
func (pq *PriorityQueue[T]) Top() (T, bool) {
	if pq.IsEmpty() {
		var zero T
		return zero, false
	}
	return pq.items[0].Value, true
}

// Size returns the number of items in the priority queue.
func (pq *PriorityQueue[T]) Size() int {
	return len(pq.items)
}

// IsEmpty returns true if the priority queue is empty.
func (pq *PriorityQueue[T]) IsEmpty() bool {
	return len(pq.items) == 0
}

// Clear removes all items from the priority queue.
func (pq *PriorityQueue[T]) Clear() {
	pq.items = pq.items[:0]
}

// heap.Interface implementation
func (pq *PriorityQueue[T]) Len() int { return len(pq.items) }

func (pq *PriorityQueue[T]) Less(i, j int) bool {
	return pq.lessFunc(pq.items[i], pq.items[j])
}

func (pq *PriorityQueue[T]) Swap(i, j int) {
	pq.items[i], pq.items[j] = pq.items[j], pq.items[i]
	pq.items[i].index = i
	pq.items[j].index = j
}

func (pq *PriorityQueue[T]) Push(x interface{}) {
	item := x.(*PriorityQueueItem[T])
	item.index = len(pq.items)
	pq.items = append(pq.items, item)
}

func (pq *PriorityQueue[T]) Pop() interface{} {
	old := pq.items
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	pq.items = old[0 : n-1]
	return item
}
