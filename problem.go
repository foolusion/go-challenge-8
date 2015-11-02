package main

import (
	"bytes"
	"fmt"
)

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
			results = append(results, actioner(v.row, v.col, i))
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
			fmt.Println("bad row")
			return false
		}
		col := s.col(i)
		if !checkslice(col) {
			fmt.Println("bad col")
			return false
		}
		box := s.box(i)
		if !checkslice(box) {
			fmt.Println("bad box")
			return false
		}
	}
	return true
}

func (p problem) Result(s state, a action) state {
	return a(s)
}

func (p problem) stepCost(s state, a action) int {
	return 1
}

func checkslice(row []int) bool {
	// row is invalid because of wrong length
	if len(row) != rowLen {
		fmt.Println("wrong row length", row)
		return false
	}

	index, value := 0, 1
	for value < 10 {
		// value not found in row
		if index == len(row) {
			fmt.Printf("%v not found in %v\n", value, row)
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
