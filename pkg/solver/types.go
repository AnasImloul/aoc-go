package solver

// Numeric is a constraint for numeric types.
type Numeric interface {
	int | int32 | int64 | float32 | float64
}

// Point is a generic 2D point with numeric coordinates.
type Point[T Numeric] struct {
	X, Y T
}

// Add adds two points.
func (p Point[T]) Add(other Point[T]) Point[T] {
	return Point[T]{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub subtracts two points.
func (p Point[T]) Sub(other Point[T]) Point[T] {
	return Point[T]{X: p.X - other.X, Y: p.Y - other.Y}
}

// Point3D is a generic 3D point with numeric coordinates.
type Point3D[T Numeric] struct {
	X, Y, Z T
}

// Add adds two 3D points.
func (p Point3D[T]) Add(other Point3D[T]) Point3D[T] {
	return Point3D[T]{X: p.X + other.X, Y: p.Y + other.Y, Z: p.Z + other.Z}
}

// Sub subtracts two 3D points.
func (p Point3D[T]) Sub(other Point3D[T]) Point3D[T] {
	return Point3D[T]{X: p.X - other.X, Y: p.Y - other.Y, Z: p.Z - other.Z}
}
