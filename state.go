package main

import (
	"fmt"
	"log"
)

type cell struct {
	row int
	col int
}

func (c cell) box() int {
	row := c.row / 3
	col := c.col / 3
	return row*3 + col
}

type state []byte

func NewStateFromState(s state) state {
	n := make(state, len(s))
	copy(n, s)
	return n
}

func emptyCells(s state) []cell {
	cells := make([]cell, 0, 81)
	for i, v := range s {
		if v == '_' {
			c := cell{
				row: i / rowLen,
				col: i % rowLen,
			}
			cells = append(cells, c)
		}
	}
	return cells
}

func (s state) row(num int) []int {
	row := make([]int, rowLen)
	start := num * rowLen
	for i := 0; i < rowLen; i++ {
		row[i] = ctoi(s[start+i])
	}
	return row
}

func (s state) col(num int) []int {
	col := make([]int, rowLen)
	for i := 0; i < rowLen; i++ {
		col[i] = ctoi(s[i*rowLen+num])
	}
	return col
}

func (s state) box(num int) []int {
	var start int
	switch num {
	case 0:
		start = 0
	case 1:
		start = 3
	case 2:
		start = 6
	case 3:
		start = 27
	case 4:
		start = 30
	case 5:
		start = 33
	case 6:
		start = 54
	case 7:
		start = 57
	case 8:
		start = 60
	default:
		log.Fatal(fmt.Errorf("box: %v is not a valid input. Expected 0-8"))
	}
	return []int{
		ctoi(s[0+start]),
		ctoi(s[1+start]),
		ctoi(s[2+start]),
		ctoi(s[9+start]),
		ctoi(s[10+start]),
		ctoi(s[11+start]),
		ctoi(s[18+start]),
		ctoi(s[19+start]),
		ctoi(s[20+start]),
	}
}

func (s state) rowHas(rowNum, value int) bool {
	return rowColBoxHas(s.row, rowNum, value)
}

func (s state) colHas(colNum, value int) bool {
	return rowColBoxHas(s.col, colNum, value)
}

func (s state) boxHas(boxNum, value int) bool {
	return rowColBoxHas(s.box, boxNum, value)
}

func rowColBoxHas(f func(int) []int, num, value int) bool {
	row := f(num)
	for _, v := range row {
		if v == value {
			return true
		}
	}
	return false
}
