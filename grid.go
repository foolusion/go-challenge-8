package main

import (
	"errors"
	"strings"
)

var (
	ErrRemovedLastValue = errors.New("Removed last value")
)

type nstate map[string]string

func newStateFrom(s nstate) nstate {
	newState := make(nstate, 81*2)
	for k, v := range s {
		newState[k] = v
	}
	return newState
}

func parseGrid(grid string) {
	values := make(map[string][]string, 81*2)
	for _, s := range squares {
		values[s] = digits
	}
	for s, d := range gridValues(grid) {
		// TODO: real implementation
		if s == d {
			return
		}
	}
}

func gridValues(grid string) nstate {
	results := make(nstate, 81*2)
	return results
}

func (n nstate) assign(s, d string) error {
	// get all the values != d from square s
	otherValues := strings.Replace(n[s], d, "", -1)
	// eliminate each other value from the state at square s
	for _, d2 := range otherValues {
		if err := n.eliminate(s, string(d2)); err != nil {
			return err
		}
	}
	return nil
}

func (n nstate) has(sq, val string) bool {
	return strings.Contains(n[sq], val)
}

func (n nstate) eliminate(s, d string) error {
	// d is already eliminated
	if !n.has(s, d) {
		return nil
	}

	n[s] = strings.Replace(n[s], string(d), "", -1)
	if len(n[s]) == 0 {
		return ErrRemovedLastValue
	} else if len(n[s]) == 1 {
		d2 := n[s]
		for _, p := range peers[s] {
			if err := n.eliminate(p, d2); err != nil {
				return err
			}
		}
	}
	// TODO: implement the rest
	return nil
}
