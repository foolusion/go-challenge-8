package main

type action struct {
	square string
	value  string
}

func (a action) do(s nstate) (nstate, error) {
	res := newStateFrom(s)
	err := res.assign(a.square, a.value)
	return res, err
}
