package utils

import "github.com/AnasImloul/aoc-go/pkg/solver"

// AreAligned checks if three points are collinear.
func AreAligned[T solver.Numeric](point1, point2, point3 solver.Point[T]) bool {
	det := point1.X*(point2.Y-point3.Y) +
		point2.X*(point3.Y-point1.Y) +
		point3.X*(point1.Y-point2.Y)
	return det == 0
}

// DistanceSquared returns the squared distance between two points.
func DistanceSquared[T solver.Numeric](point1, point2 solver.Point[T]) T {
	dx, dy := point1.X-point2.X, point1.Y-point2.Y
	return dx*dx + dy*dy
}
