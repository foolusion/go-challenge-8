package main

import "testing"

var s state = []byte("824450507216973687861636086555015399559551508775001582149390418905747023746521756")

func TestRow(t *testing.T) {
	var tests = []struct {
		in  int
		out []int
	}{
		{in: 0, out: []int{8, 2, 4, 4, 5, 0, 5, 0, 7}},
		{in: 1, out: []int{2, 1, 6, 9, 7, 3, 6, 8, 7}},
		{in: 2, out: []int{8, 6, 1, 6, 3, 6, 0, 8, 6}},
		{in: 3, out: []int{5, 5, 5, 0, 1, 5, 3, 9, 9}},
		{in: 4, out: []int{5, 5, 9, 5, 5, 1, 5, 0, 8}},
		{in: 5, out: []int{7, 7, 5, 0, 0, 1, 5, 8, 2}},
		{in: 6, out: []int{1, 4, 9, 3, 9, 0, 4, 1, 8}},
		{in: 7, out: []int{9, 0, 5, 7, 4, 7, 0, 2, 3}},
		{in: 8, out: []int{7, 4, 6, 5, 2, 1, 7, 5, 6}},
	}
	for _, v := range tests {
		out := s.row(v.in)
		if len(out) != len(v.out) {
			t.Errorf("s.row(%v): got %v expected %v", v.in, out, v.out)
		}
		for i, val := range out {
			if v.out[i] != val {
				t.Errorf("s.row(%v): got %v expected %v", v.in, out, v.out)
			}
		}
	}
}

func TestCol(t *testing.T) {
	var tests = []struct {
		in  int
		out []int
	}{
		{in: 0, out: []int{8, 2, 8, 5, 5, 7, 1, 9, 7}},
		{in: 1, out: []int{2, 1, 6, 5, 5, 7, 4, 0, 4}},
		{in: 2, out: []int{4, 6, 1, 5, 9, 5, 9, 5, 6}},
		{in: 8, out: []int{7, 7, 6, 9, 8, 2, 8, 3, 6}},
	}
	for _, v := range tests {
		out := s.col(v.in)
		if len(out) != len(v.out) {
			t.Errorf("s.col(%v): got %v expected %v", v.in, out, v.out)
		}
		for i, val := range out {
			if v.out[i] != val {
				t.Errorf("s.col(%v): got %v expected %v", v.in, out, v.out)
			}
		}
	}
}

func TestBox(t *testing.T) {
	var tests = []struct {
		in  int
		out []int
	}{
		{in: 0, out: []int{8, 2, 4, 2, 1, 6, 8, 6, 1}},
		{in: 1, out: []int{4, 5, 0, 9, 7, 3, 6, 3, 6}},
		{in: 2, out: []int{5, 0, 7, 6, 8, 7, 0, 8, 6}},
		{in: 8, out: []int{4, 1, 8, 0, 2, 3, 7, 5, 6}},
	}
	for _, v := range tests {
		out := s.box(v.in)
		if len(out) != len(v.out) {
			t.Errorf("s.box(%v): got %v expected %v", v.in, out, v.out)
		}
		for i, val := range out {
			if v.out[i] != val {
				t.Errorf("s.box(%v): got %v expected %v", v.in, out, v.out)
			}
		}
	}
}
