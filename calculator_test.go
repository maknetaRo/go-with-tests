package calculator_test

import (
	"calculator"
	"testing"
	"math"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2, 2}, want: 4},
		{inputs: []float64{2, 2, 2}, want : 6},
	}
	for _, tc := range testCases {
		got := calculator.Add(tc.inputs...)
		if tc.want != got {
			t.Errorf("Add(%v): want %f, got %f", tc.inputs, tc.want, got)
		}
	}
}


func TestSubtract(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want float64
	}
	testCases := []testCase{
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{2, 2, 2}, want: -2},
		{inputs: []float64{10, 5, 3, 2}, want: 0},
	
	}
	for _, tc := range testCases {
		got := calculator.Subtract(tc.inputs...)
		if tc.want != got {
			t.Errorf("Subtract(%v): want %f, got %f", tc.inputs, tc.want, got)
		}
	}
}

func TestMultiply(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want float64
	}
	testCases := []testCase{
		{inputs: []float64{},  want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{0.5, 2}, want: 1},
	}
	for _, tc := range testCases {
		got := calculator.Multiply(tc.inputs...)
		if tc.want != got {
			t.Errorf("Multiply(%v): want %f, got %f", tc.inputs, tc.want, got)
		}
	}
}

func TestDivideMany(t *testing.T) {
	t.Parallel()
	type testCase struct {
		inputs []float64
		want float64
	}
	testCases := []testCase {
		{inputs: []float64{}, want: 0},
		{inputs: []float64{2}, want: 2},
		{inputs: []float64{1, 1}, want: 1},
		{inputs: []float64{30, 5, 3}, want: 2},
		{inputs: []float64{-2, 1}, want: -2},
		{inputs: []float64{1, 3}, want: 0.3333},

	}
	for _, tc := range testCases{
		got, err := calculator.DivideMany(tc.inputs...)
		if err != nil {
			t.Fatalf("Divide(%v): want no error for valid input, got %v", tc.inputs, err)
		}
		if !closeEnough(tc.want, got, 0.001) {
			t.Errorf("Divide(%v): want %f, got %f", tc.inputs, tc.want, got)
		}
		// if tc.want != got {
		// 	t.Errorf("Divide(%v): want %f, got %f",tc.inputs, tc.want, got)
		// }
	}
	
	}
func TestDivideManyInvalid(t * testing.T) {
	t.Parallel() 
	_, err := calculator.DivideMany(10, 5, 0, 7) 
	if err == nil {
		t.Error("Divide(10, 5, 0, 7): want error for zero divisor, got nil")
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