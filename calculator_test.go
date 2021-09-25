package calculator_test

import (
	"calculator"
	"testing"
	"math"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 4},
		{a: 1, b: 1, want: 2},
		{a: 5, b: 0, want: 5},
		{a: -1, b: 1, want: 0},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Add(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: 2, b: 2, want: 0},
		{a: 2.5, b: 2, want: 0.5},
		{a: 5, b: 10, want: -5},
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Subtract(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase{
		{a: -1, b: 4, want: -4},
		{a: 0, b: 4, want: 0},
		{a: 0.5, b: 2, want: 1},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.a, tc.b)
		if tc.want != got {
			t.Errorf("Multiply(%f, %f): want %f, got %f", tc.a, tc.b, tc.want, got)
		}
	}
}

func TestDivide(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a, b float64
		want float64
	}
	testCases := []testCase {
		{a: 2, b: 2, want: 1},
		{a: 2, b: 4, want: 0.5},
		{a: -3, b: -1, want: 3},
		{a: 10, b: 2, want: 5},
		{a: 6, b: -2, want: -3},
		{a: 1, b: 3, want: 0.333333,},
	}
	for _, tc := range testCases {
		got, err := calculator.Divide(tc.a, tc.b)	
		if err != nil {
			t.Fatalf("Divide(%f, %f): want no error for valid input, got %v", tc.a, tc.b, err)
		}

		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%f, %f): want %f, got %f",tc.a, tc.b, tc.want, got)
		}
		// if tc.want != got {
		// 	t.Errorf("Divide(%f, %f): want %f, got %f",tc.a, tc.b, tc.want, got)
		// }
	}
}

func TestDivideInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Divide(1, 0)
	if err == nil {
		t.Error("Divde(1, 0): want error for valid input, got nil")
	}
}

func closeEnough(a, b, tolerance float64) bool {
	return math.Abs(a-b) <= tolerance
}

func TestSqrt(t *testing.T) {
	t.Parallel()
	type testCase struct {
		a float64
		want float64
	}
	testCases := []testCase{
		{a: 4, want: 2},
		{a: 0.25, want: 0.5},
		{a: 0.0025, want: 0.05},
		{a: 0.0001, want: 0.01},
		{a: 121, want: 11},
		{a: 8, want: 2.828427125},
	}
	for _, tc := range testCases {
		got, err := calculator.Sqrt(tc.a)
		if err != nil {
			t.Fatalf("Sqrt(%f): want no error for valid input, got %v", tc.a,  err)
		}
		// if tc.want != got {
		// 	t.Errorf("Sqrt(%f): want %f, got %f", tc.a, tc.want, got)
		// }
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Sqrt(%f): want %f, got %f",tc.a, tc.want, got)
		}
	}
}

func TestSqrtInvalid(t *testing.T) {
	t.Parallel()
	_, err := calculator.Sqrt(-1)
	if err == nil {
		t.Error("Sqrt(-1): want error for valid input, got nil")
	}
}