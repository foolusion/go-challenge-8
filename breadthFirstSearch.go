package main

import (
	"fmt"
	"log"
)

func in(a []nstate, v nstate) bool {
	for _, state := range a {
		for _, sq := range squares {
			if state[sq] != v[sq] {
				return false
			}
		}
		return true
	}
	return true
}

func stateIn(a []*node, v nstate) bool {
	for _, n := range a {
		for _, sq := range squares {
			if n.state[sq] != v[sq] {
				return false
			}
		}
		return true
	}
	return true
}

func pop(a []*node) (*node, []*node) {
	if len(a) == 0 {
		return nil, a
	}
	return a[0], a[1:]
}

func breadthFirstSearch(p problem) (nstate, error) {
	initialState, err := p.initialState()
	if err != nil {
		log.Fatal(err)
	}
	n := &node{state: initialState, pathCost: 0}
	if p.GoalTest(n.state) {
		return n.state, nil
	}
	frontier := []*node{n}
	explored := make([]nstate, 0, 81)
	for {
		if len(frontier) == 0 {
			return nil, fmt.Errorf("No solution exists")
		}
		n, frontier = pop(frontier)
		explored = append(explored, n.state)
		for _, action := range p.actions(n.state) {
			child, err := childNode(p, n, action)
			if err != nil {
				continue
			}
			if !in(explored, child.state) && !stateIn(frontier, child.state) {
				if p.GoalTest(child.state) {
					return child.state, nil
				}
				frontier = append(frontier, child)
			}
		}
	}
}
