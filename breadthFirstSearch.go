package main

import (
	"bytes"
	"fmt"
)

func in(a []state, v state) bool {
	for _, s := range a {
		if bytes.Equal(s, v) {
			return true
		}
	}
	return false
}

func stateIn(a []*node, v state) bool {
	for _, n := range a {
		if bytes.Equal(n.state, v) {
			return true
		}
	}
	return false
}

func pop(a []*node) (*node, []*node) {
	if len(a) == 0 {
		return nil, a
	}
	return a[0], a[1:]
}

func breadthFirstSearch(p problem) (state, error) {
	n := &node{state: p.initialState(), pathCost: 0}
	if p.GoalTest(n.state) {
		return n.state, nil
	}
	frontier := []*node{n}
	explored := make([]state, 0, 81)
	for {
		if len(frontier) == 0 {
			return nil, fmt.Errorf("No solution exists")
		}
		n, frontier = pop(frontier)
		explored = append(explored, n.state)
		for _, action := range p.actions(n.state) {
			child := childNode(p, n, action)
			if !in(explored, child.state) && !stateIn(frontier, child.state) {
				if p.GoalTest(child.state) {
					return child.state, nil
				}
				frontier = append(frontier, child)
			}
		}
	}
}
