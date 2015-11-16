package main

import "errors"

var (
	ErrCutoff     = errors.New("Cutoff")
	ErrNoSolution = errors.New("No Solution")
)

func depthLimitedSearch(p problem, limit int) (nstate, error) {
	initialState, err := p.initialState()
	if err != nil {
		return initialState, err
	}
	return recursiveDLS(&node{state: initialState, pathCost: 0}, p, limit)
}

func recursiveDLS(n *node, p problem, limit int) (nstate, error) {
	if p.GoalTest(n.state) {
		return n.state, nil
	} else if limit == 0 {
		return n.state, ErrCutoff
	} else {
		cutoffOccurred := false
		for _, action := range p.actions(n.state) {
			child, err := childNode(p, n, action)
			if err != nil {
				continue
			}
			result, err := recursiveDLS(child, p, limit-1)
			if err == ErrCutoff {
				cutoffOccurred = true
			} else if err != ErrNoSolution {
				return result, nil
			}
		}
		if cutoffOccurred {
			return n.state, ErrCutoff
		} else {
			return n.state, ErrNoSolution
		}
	}
}
