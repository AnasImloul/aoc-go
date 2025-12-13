package containers

import (
	"github.com/AnasImloul/aoc-go/pkg/utils/types"
	"github.com/emirpasic/gods/trees/redblacktree"
)

// OrderedSet is a generic ordered set implementation (like C++ set).
// It maintains elements in sorted order using a red-black tree.
//
// Time complexity (all operations are O(log n)):
//   - Insert: O(log n)
//   - Erase: O(log n)
//   - Contains: O(log n)
//   - Size/IsEmpty: O(1)
type OrderedSet[T types.Ordered] struct {
	tree *redblacktree.Tree
}

// NewOrderedSet creates a new ordered set.
func NewOrderedSet[T types.Ordered]() *OrderedSet[T] {
	return &OrderedSet[T]{
		tree: redblacktree.NewWith(orderedComparator[T]()),
	}
}

// Insert adds an item to the set if it doesn't already exist.
// Returns true if the item was added, false if it already existed.
// Time complexity: O(log n)
func (os *OrderedSet[T]) Insert(item T) bool {
	_, found := os.tree.Get(item)
	if found {
		return false
	}
	os.tree.Put(item, struct{}{})
	return true
}

// Erase removes an item from the set.
// Returns true if the item was removed, false if it didn't exist.
// Time complexity: O(log n)
func (os *OrderedSet[T]) Erase(item T) bool {
	_, found := os.tree.Get(item)
	if !found {
		return false
	}
	os.tree.Remove(item)
	return true
}

// Contains checks if the set contains an item.
// Time complexity: O(log n)
func (os *OrderedSet[T]) Contains(item T) bool {
	_, found := os.tree.Get(item)
	return found
}

// Size returns the number of items in the set.
// Time complexity: O(1)
func (os *OrderedSet[T]) Size() int {
	return os.tree.Size()
}

// IsEmpty returns true if the set is empty.
// Time complexity: O(1)
func (os *OrderedSet[T]) IsEmpty() bool {
	return os.tree.Empty()
}

// Clear removes all items from the set.
// Time complexity: O(n)
func (os *OrderedSet[T]) Clear() {
	os.tree.Clear()
}

// Items returns a copy of all items in sorted order.
// Time complexity: O(n)
func (os *OrderedSet[T]) Items() []T {
	result := make([]T, 0, os.tree.Size())
	it := os.tree.Iterator()
	for it.Next() {
		result = append(result, it.Key().(T))
	}
	return result
}
