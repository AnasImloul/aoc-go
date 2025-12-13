package containers

import (
	"github.com/AnasImloul/aoc-go/pkg/utils/types"
	"github.com/emirpasic/gods/trees/redblacktree"
)

// OrderedMap is a generic ordered map implementation (like C++ map).
// It maintains keys in sorted order using a red-black tree.
//
// Time complexity (all operations are O(log n)):
//   - Set: O(log n)
//   - Delete: O(log n)
//   - Get/Contains: O(log n)
//   - Size/IsEmpty: O(1)
type OrderedMap[K types.Ordered, V any] struct {
	tree *redblacktree.Tree
}

// NewOrderedMap creates a new ordered map.
func NewOrderedMap[K types.Ordered, V any]() *OrderedMap[K, V] {
	return &OrderedMap[K, V]{
		tree: redblacktree.NewWith(orderedComparator[K]()),
	}
}

// Set sets the value for a key.
// Time complexity: O(log n)
func (om *OrderedMap[K, V]) Set(key K, value V) {
	om.tree.Put(key, value)
}

// Get returns the value for a key and whether it exists.
// Time complexity: O(log n)
func (om *OrderedMap[K, V]) Get(key K) (V, bool) {
	value, found := om.tree.Get(key)
	if !found {
		var zero V
		return zero, false
	}
	return value.(V), true
}

// Delete removes a key from the map.
// Returns true if the key was removed, false if it didn't exist.
// Time complexity: O(log n)
func (om *OrderedMap[K, V]) Delete(key K) bool {
	_, found := om.tree.Get(key)
	if !found {
		return false
	}
	om.tree.Remove(key)
	return true
}

// Contains checks if the map contains a key.
// Time complexity: O(log n)
func (om *OrderedMap[K, V]) Contains(key K) bool {
	_, found := om.tree.Get(key)
	return found
}

// Size returns the number of key-value pairs in the map.
// Time complexity: O(1)
func (om *OrderedMap[K, V]) Size() int {
	return om.tree.Size()
}

// IsEmpty returns true if the map is empty.
// Time complexity: O(1)
func (om *OrderedMap[K, V]) IsEmpty() bool {
	return om.tree.Empty()
}

// Clear removes all key-value pairs from the map.
// Time complexity: O(n)
func (om *OrderedMap[K, V]) Clear() {
	om.tree.Clear()
}

// Keys returns a copy of all keys in sorted order.
// Time complexity: O(n)
func (om *OrderedMap[K, V]) Keys() []K {
	result := make([]K, 0, om.tree.Size())
	it := om.tree.Iterator()
	for it.Next() {
		result = append(result, it.Key().(K))
	}
	return result
}
