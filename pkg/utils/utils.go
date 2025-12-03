package utils

import (
	"fmt"
	"strconv"
	"strings"
)

// Numerical is a constraint for numerical types.
type Numerical interface {
	int | int32 | int64 | float32 | float64
}

// ToIntSlice converts a string with a separator to a slice of integers.
// Panics on parse error - use TryToIntSlice for error handling.
func ToIntSlice(line string, sep string) []int {
	nums, err := TryToIntSlice(line, sep)
	if err != nil {
		panic(err)
	}
	return nums
}

// TryToIntSlice converts a string with a separator to a slice of integers.
// Returns an error if any element cannot be parsed as an integer.
func TryToIntSlice(line string, sep string) ([]int, error) {
	var nums []int
	for _, sNum := range strings.Split(line, sep) {
		num, err := strconv.Atoi(sNum)
		if err != nil {
			return nil, fmt.Errorf("failed to parse %q as integer: %w", sNum, err)
		}
		nums = append(nums, num)
	}
	return nums, nil
}

// ReverseString reverses a string.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// Rotate rotates a list of strings (matrix) by the specified degrees (90, 180, 270).
func Rotate(matrix []string, degrees int) ([]string, error) {
	if len(matrix) == 0 {
		return nil, fmt.Errorf("matrix cannot be empty")
	}

	rows := len(matrix)
	cols := len(matrix[0])

	// Ensure all rows have the same length
	for _, row := range matrix {
		if len(row) != cols {
			return nil, fmt.Errorf("all rows must have the same length")
		}
	}

	switch degrees {
	case 90:
		return rotate90(matrix, rows, cols), nil
	case 180:
		return rotate180(matrix, rows, cols), nil
	case 270:
		return rotate270(matrix, rows, cols), nil
	default:
		return nil, fmt.Errorf("invalid degrees, must be one of 90, 180, 270")
	}
}

func rotate90(matrix []string, rows, cols int) []string {
	result := make([]string, cols)
	for i := 0; i < cols; i++ {
		var newRow string
		for j := rows - 1; j >= 0; j-- {
			newRow += string(matrix[j][i])
		}
		result[i] = newRow
	}
	return result
}

func rotate180(matrix []string, rows, cols int) []string {
	result := make([]string, rows)
	for i := 0; i < rows; i++ {
		var newRow string
		for j := cols - 1; j >= 0; j-- {
			newRow += string(matrix[rows-1-i][j])
		}
		result[i] = newRow
	}
	return result
}

func rotate270(matrix []string, rows, cols int) []string {
	result := make([]string, cols)
	for i := 0; i < cols; i++ {
		var newRow string
		for j := 0; j < rows; j++ {
			newRow += string(matrix[j][cols-1-i])
		}
		result[i] = newRow
	}
	return result
}

// NumberOfDigits returns the number of digits in an integer.
func NumberOfDigits(n int64) int {
	if n == 0 {
		return 1
	}

	count := 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}

// Pow10 returns 10 raised to the power of n.
func Pow10(n int) int {
	res := 1
	for n > 0 {
		res *= 10
		n--
	}
	return res
}

// MakeGrid creates a 2D grid with the given dimensions and initial value.
func MakeGrid[T any](n, m int, initialValue T) [][]T {
	res := make([][]T, n)
	for i := 0; i < n; i++ {
		res[i] = make([]T, m)
		for j := 0; j < m; j++ {
			res[i][j] = initialValue
		}
	}
	return res
}

// GridSize returns the dimensions of a grid.
func GridSize(grid [][]int) (int, int) {
	return len(grid), len(grid[0])
}

// IsOutOfBounds checks if coordinates are outside grid bounds.
func IsOutOfBounds(i, j, n, m int) bool {
	return i < 0 || i >= n || j < 0 || j >= m
}

// IsInBounds checks if coordinates are inside grid bounds.
func IsInBounds(i, j, n, m int) bool {
	return !IsOutOfBounds(i, j, n, m)
}

// Abs returns the absolute value of an integer.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// ParseInt parses a string to an integer, panicking on error.
// Use strconv.Atoi directly for error handling.
func ParseInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		panic(fmt.Sprintf("failed to parse int %q: %v", s, err))
	}
	return num
}

// Sum returns the sum of a slice of numbers.
func Sum[T Numerical](numbers []T) T {
	var sum T = 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// Must panics if err is not nil, otherwise returns t.
func Must[T any](t T, err error) T {
	if err != nil {
		panic(err)
	}
	return t
}
