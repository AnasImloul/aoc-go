package types

// Ordered is a constraint for types that support ordering.
type Ordered interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | uintptr | float32 | float64 | string
}

// Numerical is a constraint for numerical types.
type Numerical interface {
	int | int32 | int64 | float32 | float64
}
