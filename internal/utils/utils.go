package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ToIntSlice(line string, sep string) []int {
	var nums []int
	for _, sNum := range strings.Split(line, sep) {
		num, err := strconv.Atoi(sNum)
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}

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

func Pow10(n int) int {
	res := 1
	for n > 0 {
		res *= 10
		n--
	}
	return res
}

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

func GridSize(grid [][]int) (int, int) {
	return len(grid), len(grid[0])
}

func IsOutOfBounds(i, j, n, m int) bool {
	return i < 0 || i >= n || j < 0 || j >= m
}
