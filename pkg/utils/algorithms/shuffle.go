package algorithms

import (
	"math/rand"
	"time"
)

// Shuffle shuffles the slice in-place using Fisher-Yates algorithm.
func Shuffle[T any](slice []T) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := len(slice) - 1; i > 0; i-- {
		j := r.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
