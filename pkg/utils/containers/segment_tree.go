package containers

// SegmentTree is a generic segment tree for range queries.
type SegmentTree[T any] struct {
	n       int
	data    []T
	combine func(T, T) T
	zero    T
}

// NewSegmentTree creates a new segment tree with initial values and a combine function.
func NewSegmentTree[T any](values []T, combine func(T, T) T, zero T) *SegmentTree[T] {
	n := len(values)
	st := &SegmentTree[T]{
		n:       n,
		data:    make([]T, 2*n),
		combine: combine,
		zero:    zero,
	}
	copy(st.data[n:], values)
	for i := n - 1; i > 0; i-- {
		st.data[i] = combine(st.data[2*i], st.data[2*i+1])
	}
	return st
}

// Update updates the value at index i.
func (st *SegmentTree[T]) Update(i int, value T) {
	i += st.n
	st.data[i] = value
	for i > 1 {
		i /= 2
		st.data[i] = st.combine(st.data[2*i], st.data[2*i+1])
	}
}

// Query returns the result of the range query from l to r (exclusive).
func (st *SegmentTree[T]) Query(l, r int) T {
	l += st.n
	r += st.n
	result := st.zero
	for l < r {
		if l&1 == 1 {
			result = st.combine(result, st.data[l])
			l++
		}
		if r&1 == 1 {
			r--
			result = st.combine(result, st.data[r])
		}
		l /= 2
		r /= 2
	}
	return result
}
