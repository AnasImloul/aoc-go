package containers

// FenwickTree is a Binary Indexed Tree for range sum queries.
type FenwickTree struct {
	tree []int
	n    int
}

// NewFenwickTree creates a new FenwickTree with n elements.
func NewFenwickTree(n int) *FenwickTree {
	return &FenwickTree{
		tree: make([]int, n+1),
		n:    n,
	}
}

// Update updates the value at index i by adding delta.
func (ft *FenwickTree) Update(i int, delta int) {
	i++
	for i <= ft.n {
		ft.tree[i] += delta
		i += i & -i
	}
}

// Query returns the prefix sum from 0 to i (inclusive).
func (ft *FenwickTree) Query(i int) int {
	i++
	sum := 0
	for i > 0 {
		sum += ft.tree[i]
		i -= i & -i
	}
	return sum
}

// QueryRange returns the sum from l to r (inclusive).
func (ft *FenwickTree) QueryRange(l, r int) int {
	return ft.Query(r) - ft.Query(l-1)
}
