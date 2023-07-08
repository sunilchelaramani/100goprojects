package main

import (
	"fmt"
	"testing"
)

// function compares two integers and returns the smaller one
func IntMin(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

// test function is a function that starts with the word Test and takes one argument of type *testing.T
// test function is run when the test is executed
// test function is run with a struct of type *testing.T
// following test function will fail because IntMin(2, -2) should return -2 but returns 2
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {

		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// following test function will fail because IntMin(1, 1) should return 1 but returns 0
func TestIntMinTableDriven(t *testing.T) {
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	for _, tt := range tests {

		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// benchmark function is a function that starts with the word Benchmark and takes one argument of type *testing.B
// benchmark function is run when the benchmark is executed
// benchmark function is run with a struct of type *testing.B
// following benchmark function will run IntMin b.N times
// b.N is a value that is set by the testing framework and is adjusted until the benchmark function lasts long enough to be timed reliably
func BenchmarkIntMin(b *testing.B) {

	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}
