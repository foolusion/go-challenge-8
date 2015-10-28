package main

import (
	"fmt"
	"log"
	"strconv"
)

const rowLen = 9

type state string

func (s state) row(num int) []int {
	row := make([]int, rowLen)
	start := num * rowLen
	for i := 0; i < rowLen; i++ {
		val, err := strconv.Atoi(string(s[start : start+i+1]))
		if err != nil {
			log.Fatal(err)
		}
		row[i] = val
	}
	return row
}

func main() {
	fmt.Println("Hello World")
}
