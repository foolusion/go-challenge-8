package main

import "errors"

var (
	ErrCutoff     = errors.New("Cutoff")
	ErrNoSolution = errors.New("No Solution")
)

func depthLimitedSearch(p problem, limit int) (state, error) {
	return recursiveDLS(&node{state: state(p), pathCost: 0}, p, limit)
}

func recursiveDLS(n *node, p problem, limit int) (state, error) {
	if p.GoalTest(n.state) {
		return n.state, nil
	} else if limit == 0 {
		return n.state, ErrCutoff
	} else {
		cutoffOccurred := false
		for _, action := range p.actions(n.state) {
			child := childNode(p, n, action)
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
