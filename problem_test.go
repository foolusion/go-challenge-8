package main

import (
	"math/rand"
	"testing"
)

func TestCheckSlice(t *testing.T) {
	tests := []struct {
		in  []int
		out bool
	}{
		{in: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}, out: true},     // good
		{in: []int{9, 8, 7, 6, 5, 4, 3, 2, 1}, out: true},     // good reverse
		{in: []int{5, 7, 9, 1, 3, 6, 4, 8, 2}, out: true},     // good randomish
		{in: []int{1, 1, 2, 3, 4, 5, 6, 7, 8}, out: false},    // bad repeated number
		{in: []int{0, 1, 2, 3, 4, 5, 6, 7, 8}, out: false},    // bad missing 9
		{in: []int{10, 1, 2, 3, 4, 5, 6, 7, 8}, out: false},   // bad missing 9
		{in: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, out: false}, // bad too long
		{in: []int{}, out: false},                             // bad too short
	}
	for _, test := range tests {
		out := checkslice(test.in)
		if test.out != out {
			t.Errorf("checkslice(%v): got %v expected %v", test.in, out, test.out)
		}
	}
}

func randomRow() []int {
	row := make([]int, 9)
	for i := 0; i < len(row); i++ {
		row[i] = rand.Intn(10)
	}
	return row
}

func BenchmarkCheckSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		row := randomRow()
		checkslice(row)
	}
}
