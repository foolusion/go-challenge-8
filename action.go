package main

type action func(state) state

func actioner(r, c, v int) action {
	return func(s state) state {
		res := NewStateFromState(s)
		res[r*rowLen+c] = itoc(v)
		return res
	}
}
