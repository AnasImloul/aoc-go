package utils

// Point2D represents a 2D point with int coordinates for pathfinding algorithms.
// For generic numeric coordinates, see solver.Point[T].
type Point2D struct {
	X, Y int
}

// Point3D represents a 3D point with int coordinates.
// For generic numeric coordinates, see solver.Point3D[T].
type Point3D struct {
	X, Y, Z int
}

// Add adds two points
func (p Point2D) Add(other Point2D) Point2D {
	return Point2D{X: p.X + other.X, Y: p.Y + other.Y}
}

// Sub subtracts two points
func (p Point2D) Sub(other Point2D) Point2D {
	return Point2D{X: p.X - other.X, Y: p.Y - other.Y}
}

// ManhattanDist calculates Manhattan distance to another point
func (p Point2D) ManhattanDist(other Point2D) int {
	return Abs(p.X-other.X) + Abs(p.Y-other.Y)
}

// Neighbors4 returns 4-directional neighbors
func (p Point2D) Neighbors4() []Point2D {
	return []Point2D{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
	}
}

// Neighbors8 returns 8-directional neighbors (including diagonals)
func (p Point2D) Neighbors8() []Point2D {
	return []Point2D{
		{X: p.X + 1, Y: p.Y},
		{X: p.X - 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y - 1},
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y - 1},
	}
}

// AreAligned checks if three points are collinear.
func AreAligned(point1, point2, point3 Point2D) bool {
	det := point1.X*(point2.Y-point3.Y) +
		point2.X*(point3.Y-point1.Y) +
		point3.X*(point1.Y-point2.Y)
	return det == 0
}

// DistanceSquared returns the squared distance between two points.
func DistanceSquared(point1, point2 Point2D) int {
	dx, dy := point1.X-point2.X, point1.Y-point2.Y
	return dx*dx + dy*dy
}
