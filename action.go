package main

type action struct {
	c     cell
	value int
}

func (a action) do(s state) state {
	res := NewStateFromState(s)
	res[a.c.index()] = itoc(a.value)
	return res
}
