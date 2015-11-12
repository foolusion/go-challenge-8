package main

type problem string

func (p problem) initialState() (nstate, error) {
	return parseGrid(p)
}

func (p problem) actions(s nstate) []action {
	actions := make([]action, 0, 20)
	for sq, vals := range s {
		if len(vals) > 1 {
			for _, v := range vals {
				actions = append(actions, action{square: sq, value: string(v)})
			}
		}
	}
	return actions
}

func (p problem) GoalTest(s nstate) bool {
	for _, vals := range s {
		if len(vals) > 1 {
			return false
		}
	}
	return true
}

func (p problem) Result(s nstate, a action) (nstate, error) {
	return a.do(s)
}

func (p problem) stepCost(s nstate, a action) int {
	return len(s[a.square])
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
