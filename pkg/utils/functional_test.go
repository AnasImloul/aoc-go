package utils

import (
	"testing"
)

func TestMap(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Map(input, func(x int) int { return x * 2 })
	expected := []int{2, 4, 6, 8, 10}

	if len(result) != len(expected) {
		t.Fatalf("Map result length = %d, want %d", len(result), len(expected))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map result[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestMapTypeConversion(t *testing.T) {
	input := []int{1, 2, 3}
	result := Map(input, func(x int) string {
		return string(rune('a' + x - 1))
	})
	expected := []string{"a", "b", "c"}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Map result[%d] = %q, want %q", i, v, expected[i])
		}
	}
}

func TestReduce(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	result := Reduce(input, 0, func(acc, x int) int { return acc + x })
	if result != 15 {
		t.Errorf("Reduce sum = %d, want 15", result)
	}
}

func TestReduceProduct(t *testing.T) {
	input := []int{1, 2, 3, 4}
	result := Reduce(input, 1, func(acc, x int) int { return acc * x })
	if result != 24 {
		t.Errorf("Reduce product = %d, want 24", result)
	}
}

func TestFilter(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6}
	result := Filter(input, func(x int) bool { return x%2 == 0 })
	expected := []int{2, 4, 6}

	if len(result) != len(expected) {
		t.Fatalf("Filter result length = %d, want %d", len(result), len(expected))
	}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("Filter result[%d] = %d, want %d", i, v, expected[i])
		}
	}
}

func TestFilterEmpty(t *testing.T) {
	input := []int{1, 3, 5}
	result := Filter(input, func(x int) bool { return x%2 == 0 })
	if len(result) != 0 {
		t.Errorf("Filter result = %v, want empty slice", result)
	}
}

func TestMax(t *testing.T) {
	result := Max([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if result != 9 {
		t.Errorf("Max = %d, want 9", result)
	}
}

func TestMaxSingle(t *testing.T) {
	result := Max([]int{42})
	if result != 42 {
		t.Errorf("Max([42]) = %d, want 42", result)
	}
}

func TestMaxEmpty(t *testing.T) {
	result := Max([]int{})
	if result != 0 {
		t.Errorf("Max([]) = %d, want 0", result)
	}
}

func TestMin(t *testing.T) {
	result := Min([]int{3, 1, 4, 1, 5, 9, 2, 6})
	if result != 1 {
		t.Errorf("Min = %d, want 1", result)
	}
}

func TestMinNegative(t *testing.T) {
	result := Min([]int{-5, 0, 5, -10})
	if result != -10 {
		t.Errorf("Min = %d, want -10", result)
	}
}

func TestMapChan(t *testing.T) {
	input := make(chan int, 3)
	input <- 1
	input <- 2
	input <- 3
	close(input)

	result := MapChan(input, func(x int) int { return x * 2 })

	expected := []int{2, 4, 6}
	i := 0
	for v := range result {
		if i >= len(expected) {
			t.Fatalf("MapChan produced too many values")
		}
		if v != expected[i] {
			t.Errorf("MapChan value[%d] = %d, want %d", i, v, expected[i])
		}
		i++
	}
	if i != len(expected) {
		t.Errorf("MapChan produced %d values, want %d", i, len(expected))
	}
}

func TestReduceChan(t *testing.T) {
	input := make(chan int, 5)
	for _, v := range []int{1, 2, 3, 4, 5} {
		input <- v
	}
	close(input)

	result := ReduceChan(input, 0, func(acc, x int) int { return acc + x })
	if result != 15 {
		t.Errorf("ReduceChan sum = %d, want 15", result)
	}
}

func TestFilterChan(t *testing.T) {
	input := make(chan int, 6)
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		input <- v
	}
	close(input)

	result := FilterChan(input, func(x int) bool { return x%2 == 0 })

	expected := []int{2, 4, 6}
	i := 0
	for v := range result {
		if i >= len(expected) {
			t.Fatalf("FilterChan produced too many values")
		}
		if v != expected[i] {
			t.Errorf("FilterChan value[%d] = %d, want %d", i, v, expected[i])
		}
		i++
	}
}
