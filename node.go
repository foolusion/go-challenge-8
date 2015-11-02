package main

type node struct {
	state
	parent   *node
	pathCost int
}

func childNode(p problem, parent *node, a action) *node {
	return &node{
		state:    p.Result(parent.state, a),
		parent:   parent,
		pathCost: parent.pathCost + p.stepCost(parent.state, a),
	}
}
