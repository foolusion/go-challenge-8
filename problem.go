package main

import "bytes"

type problem []byte

func (p problem) initialState() state {
	return state(p)
}

func (p problem) actions(s state) []action {
	// TODO: get all valid actions for the current state.
	results := []action{}
	e := emptyCells(s)
	for _, v := range e {
		for i := 1; i <= 9; i++ {
			if s.rowHas(v.row, i) || s.colHas(v.col, i) || s.boxHas(v.box(), i) {
				continue
			}
			results = append(results, action{c: cell{row: v.row, col: v.col}, value: i})
		}
	}
	return results
}

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

func (p problem) Result(s state, a action) state {
	return a.do(s)
}

func (p problem) stepCost(s state, a action) int {
	return s.numPossibleActions(a.c)
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
