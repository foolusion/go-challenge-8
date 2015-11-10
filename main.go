package main

import (
	"fmt"
	"strings"
)

const rowLen = 9

var cellValues = []byte("123456789")

func ctoi(c byte) int {
	if c >= '1' && c <= '9' {
		return int(c - '0')
	}
	return 0
}

var itocDict = []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

func itoc(i int) byte {
	if i > len(itocDict) {
		return '0'
	} else if i < 0 {
		return '0'
	}
	return itocDict[i]
}

func printBoard(values nstate) {
	width := 0
	for _, s := range squares {
		if width < len(values[s]) {
			width = len(values[s])
		}
	}
	width += 1

	lineSlice := make([]string, 3)
	for i := 0; i < len(lineSlice); i++ {
		lineSlice[i] = strings.Repeat("-", width*3)
	}

	line := strings.Join(lineSlice, "+")

	for _, r := range rows {
		for _, c := range cols {
			str := values[r+c]
			space := strings.Repeat(" ", width-len(str))
			fmt.Print(str, space)
			if c == "3" || c == "6" {
				fmt.Print("|")
			}
		}
		fmt.Println()

		if r == "C" || r == "F" {
			fmt.Println(line)
		}
	}
}

func cross(a, b []string) []string {
	res := make([]string, len(a)*len(b))
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			res[i*len(b)+j] = a[i] + b[j]
		}
	}
	return res
}

var (
	digits   = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	rows     = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I"}
	cols     = digits
	squares  = cross(rows, cols)
	unitlist = makeUnitList()
	units    = makeUnits()
	peers    = makePeers()
)

func makeUnitList() [][]string {
	var unitlist = make([][]string, 0, 27)
	for c := range cols {
		unitlist = append(unitlist, cross(rows, cols[c:c+1]))
	}
	for r := range rows {
		unitlist = append(unitlist, cross(rows[r:r+1], cols))
	}
	for _, rs := range [][]string{rows[:3], rows[3:6], rows[6:9]} {
		for _, cs := range [][]string{digits[:3], digits[3:6], digits[6:9]} {
			unitlist = append(unitlist, cross(rs, cs))
		}
	}
	return unitlist
}

func makeUnits() map[string][][]string {
	units := make(map[string][][]string, len(squares)*10)

	for _, s := range squares {
		for _, unit := range unitlist {
			for _, usq := range unit {
				if s == usq {
					units[s] = append(units[s], unit)
					break
				}
			}
		}
	}
	return units
}

func makePeers() map[string][]string {
	peers := make(map[string][]string)

	for _, s := range squares {
		p := make(map[string]struct{})
		for _, unit := range units[s] {
			for _, usq := range unit {
				if s != usq {
					p[usq] = struct{}{}
				}
			}
		}

		peers[s] = make([]string, 0, 20)
		for k := range p {
			peers[s] = append(peers[s], k)
		}
	}
	return peers
}

func main() {
	grid1 := "003020600900305001001806400008102900700000008006708200002609500800203009005010300"
	values, err := parseGrid(grid1)
	if err != nil {
		fmt.Println(err)
	}
	printBoard(values)
}
