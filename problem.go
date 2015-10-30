package main

import "bytes"

type problem []byte

func (p problem) GoalTest(s state) bool {
	if bytes.Contains(s, []byte{'_'}) {
		return false
	}

	for i := 0; i < rowLen; i++ {
		row := s.row(i)
		if !checkslice(row) {
			return false
		}
		col := s.col(i)
		if !checkslice(col) {
			return false
		}
		box := s.box(i)
		if !checkslice(box) {
			return false
		}
	}
	return true
}

func checkslice(row []int) bool {
	// row is invalid because of wrong length
	if len(row) != rowLen {
		return false
	}

	index, value := 0, 1
	for value < 10 {
		// value not found in row
		if index == len(row) {
			return false
		}
		// value found in row
		if row[index] == value {
			index, value = 0, value+1
			continue
		}
		index++
	}
	return true
}
