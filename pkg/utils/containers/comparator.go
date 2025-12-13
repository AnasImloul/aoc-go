package containers

import (
	"github.com/AnasImloul/aoc-go/pkg/utils/types"
	godsutils "github.com/emirpasic/gods/utils"
)

// orderedComparator creates a comparator function for Ordered types.
func orderedComparator[T types.Ordered]() godsutils.Comparator {
	return func(a, b interface{}) int {
		av := a.(T)
		bv := b.(T)
		if av < bv {
			return -1
		}
		if av > bv {
			return 1
		}
		return 0
	}
}
