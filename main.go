package main

import (
	"fmt"

	"github.com/davecheney/profile"
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

func boardstring(s state, e error) string {
	return fmt.Sprintf("%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n%v\n", string(s[0:9]), string(s[9:18]), string(s[18:27]), string(s[27:36]), string(s[36:45]), string(s[45:54]), string(s[54:63]), string(s[63:72]), string(s[72:81]))
}

func main() {
	defer profile.Start(profile.CPUProfile).Stop()
	a := problem(
		"___3_4___" +
			"_125_834_" +
			"543671982" +
			"327156498" +
			"158942637" +
			"469783521" +
			"691835274" +
			"275419863" +
			"83426715_")
	fmt.Println(boardstring(depthLimitedSearch(a, 81)))
}
