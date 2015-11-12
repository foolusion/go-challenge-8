package main

type node struct {
	state    nstate
	parent   *node
	pathCost int
}

func childNode(p problem, parent *node, a action) (*node, error) {
	state, err := p.Result(parent.state, a)
	if err != nil {
		return nil, err
	}
	return &node{
		state:    state,
		parent:   parent,
		pathCost: parent.pathCost + p.stepCost(parent.state, a),
	}, nil
}
