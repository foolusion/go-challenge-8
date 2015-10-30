package main

import (
	"fmt"
	"log"
)

type state []byte

var dict = map[byte]int{'1': 1, '2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9}

func ctoi(c byte) int {
	v, ok := dict[c]
	if !ok {
		return 0
	}
	return v
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
