package main

import (
	"errors"
	"strings"
	"unicode"
)

var (
	ErrRemovedLastValue = errors.New("Removed last value")
	ErrNoPlacement      = errors.New("No placement for d") // that's what she said
)

type nstate map[string]string

func newStateFrom(s nstate) nstate {
	newState := make(nstate, 81*2)
	for k, v := range s {
		newState[k] = v
	}
	return newState
}

func parseGrid(grid problem) (nstate, error) {
	values := make(nstate, 81*2)
	for _, s := range squares {
		values[s] = strings.Join(digits, "")
	}
	for s, d := range gridValues(grid) {
		if unicode.IsDigit(d) && d != '0' {
			if err := values.assign(s, string(d)); err != nil {
				return values, err
			}
		}
	}
	return values, nil
}

func gridValues(grid problem) map[string]rune {
	results := make(map[string]rune, 81*2)
	for i, c := range grid {
		if unicode.IsDigit(c) {
			results[squares[i]] = c
		} else if c == '.' {
			results[squares[i]] = '0'
		}
	}
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

	for _, u := range units[s] {
		dplaces := make([]string, 0, 20*2)
		for _, s2 := range u {
			if n.has(s2, d) {
				dplaces = append(dplaces, s2)
			}
		}
		if len(dplaces) == 0 {
			return ErrNoPlacement
		} else if len(dplaces) == 1 {
			if err := n.assign(dplaces[0], d); err != nil {
				return err
			}
		}
	}
	return nil
}
