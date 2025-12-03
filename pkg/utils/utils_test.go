package utils

import (
	"testing"
)

func TestToIntSlice(t *testing.T) {
	tests := []struct {
		name     string
		line     string
		sep      string
		expected []int
	}{
		{"comma separated", "1,2,3,4,5", ",", []int{1, 2, 3, 4, 5}},
		{"space separated", "10 20 30", " ", []int{10, 20, 30}},
		{"single element", "42", ",", []int{42}},
		{"negative numbers", "-1,-2,-3", ",", []int{-1, -2, -3}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ToIntSlice(tt.line, tt.sep)
			if len(result) != len(tt.expected) {
				t.Fatalf("ToIntSlice(%q, %q) = %v, want %v", tt.line, tt.sep, result, tt.expected)
			}
			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("ToIntSlice(%q, %q)[%d] = %d, want %d", tt.line, tt.sep, i, v, tt.expected[i])
				}
			}
		})
	}
}

func TestTryToIntSlice(t *testing.T) {
	// Valid input
	result, err := TryToIntSlice("1,2,3", ",")
	if err != nil {
		t.Errorf("TryToIntSlice(\"1,2,3\", \",\") returned error: %v", err)
	}
	if len(result) != 3 {
		t.Errorf("TryToIntSlice(\"1,2,3\", \",\") = %v, want [1, 2, 3]", result)
	}

	// Invalid input
	_, err = TryToIntSlice("1,abc,3", ",")
	if err == nil {
		t.Error("TryToIntSlice(\"1,abc,3\", \",\") expected error, got nil")
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"", ""},
		{"a", "a"},
		{"ab", "ba"},
		{"12345", "54321"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := ReverseString(tt.input)
			if result != tt.expected {
				t.Errorf("ReverseString(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{5, 5},
		{-5, 5},
		{0, 0},
		{-100, 100},
	}

	for _, tt := range tests {
		result := Abs(tt.input)
		if result != tt.expected {
			t.Errorf("Abs(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestSum(t *testing.T) {
	t.Run("int slice", func(t *testing.T) {
		result := Sum([]int{1, 2, 3, 4, 5})
		if result != 15 {
			t.Errorf("Sum([1,2,3,4,5]) = %d, want 15", result)
		}
	})

	t.Run("empty slice", func(t *testing.T) {
		result := Sum([]int{})
		if result != 0 {
			t.Errorf("Sum([]) = %d, want 0", result)
		}
	})

	t.Run("float64 slice", func(t *testing.T) {
		result := Sum([]float64{1.5, 2.5, 3.0})
		if result != 7.0 {
			t.Errorf("Sum([1.5,2.5,3.0]) = %f, want 7.0", result)
		}
	})
}

func TestParseInt(t *testing.T) {
	result := ParseInt("42")
	if result != 42 {
		t.Errorf("ParseInt(\"42\") = %d, want 42", result)
	}

	result = ParseInt("-123")
	if result != -123 {
		t.Errorf("ParseInt(\"-123\") = %d, want -123", result)
	}
}

func TestParseIntPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("ParseInt(\"abc\") should panic")
		}
	}()
	ParseInt("abc")
}

func TestMust(t *testing.T) {
	result := Must(42, nil)
	if result != 42 {
		t.Errorf("Must(42, nil) = %d, want 42", result)
	}
}

func TestMustPanic(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("Must with error should panic")
		}
	}()
	Must(0, &testError{})
}

type testError struct{}

func (e *testError) Error() string { return "test error" }

func TestNumberOfDigits(t *testing.T) {
	tests := []struct {
		input    int64
		expected int
	}{
		{0, 1},
		{1, 1},
		{9, 1},
		{10, 2},
		{99, 2},
		{100, 3},
		{12345, 5},
	}

	for _, tt := range tests {
		result := NumberOfDigits(tt.input)
		if result != tt.expected {
			t.Errorf("NumberOfDigits(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestPow10(t *testing.T) {
	tests := []struct {
		input    int
		expected int
	}{
		{0, 1},
		{1, 10},
		{2, 100},
		{3, 1000},
	}

	for _, tt := range tests {
		result := Pow10(tt.input)
		if result != tt.expected {
			t.Errorf("Pow10(%d) = %d, want %d", tt.input, result, tt.expected)
		}
	}
}

func TestMakeGrid(t *testing.T) {
	grid := MakeGrid(3, 4, 0)
	if len(grid) != 3 {
		t.Errorf("MakeGrid(3, 4, 0) has %d rows, want 3", len(grid))
	}
	for i, row := range grid {
		if len(row) != 4 {
			t.Errorf("MakeGrid(3, 4, 0) row %d has %d cols, want 4", i, len(row))
		}
	}
}

func TestIsInBounds(t *testing.T) {
	tests := []struct {
		i, j, n, m int
		expected   bool
	}{
		{0, 0, 3, 3, true},
		{2, 2, 3, 3, true},
		{-1, 0, 3, 3, false},
		{0, -1, 3, 3, false},
		{3, 0, 3, 3, false},
		{0, 3, 3, 3, false},
	}

	for _, tt := range tests {
		result := IsInBounds(tt.i, tt.j, tt.n, tt.m)
		if result != tt.expected {
			t.Errorf("IsInBounds(%d, %d, %d, %d) = %v, want %v",
				tt.i, tt.j, tt.n, tt.m, result, tt.expected)
		}
	}
}

func TestIsOutOfBounds(t *testing.T) {
	result := IsOutOfBounds(0, 0, 3, 3)
	if result != false {
		t.Errorf("IsOutOfBounds(0, 0, 3, 3) = %v, want false", result)
	}

	result = IsOutOfBounds(-1, 0, 3, 3)
	if result != true {
		t.Errorf("IsOutOfBounds(-1, 0, 3, 3) = %v, want true", result)
	}
}
