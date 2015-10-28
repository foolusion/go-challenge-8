package main

import (
	"fmt"
	"log"
	"strconv"
)

type state string

var dict = map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}

func ctoi(c string) int {
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
		val, err := strconv.Atoi(string(s[start+i]))
		if err != nil {
			log.Fatal(err)
		}
		row[i] = val
	}
	return row
}

func (s state) col(num int) []int {
	col := make([]int, rowLen)
	for i := 0; i < rowLen; i++ {
		val, err := strconv.Atoi(string(s[i*rowLen+num]))
		if err != nil {
			log.Fatal(err)
		}
		col[i] = val
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
		ctoi(string(s[0+start])),
		ctoi(string(s[1+start])),
		ctoi(string(s[2+start])),
		ctoi(string(s[9+start])),
		ctoi(string(s[10+start])),
		ctoi(string(s[11+start])),
		ctoi(string(s[18+start])),
		ctoi(string(s[19+start])),
		ctoi(string(s[20+start])),
	}
}
