package datastructures

import "slices"

type vertex struct {
	value            string
	adjacentVertexes []*vertex
}

func (v *vertex) AddDirectedVertex(newVertex *vertex) {
	v.adjacentVertexes = append(v.adjacentVertexes, newVertex)
}

func (v *vertex) AddUndirectedVertex(newVertex *vertex) {
	// Avoid infinite loop
	if slices.Contains(v.adjacentVertexes, newVertex) {
		return
	}
	v.adjacentVertexes = append(v.adjacentVertexes, newVertex)
	newVertex.AddUndirectedVertex(v)
}

func NewVertex(val string) *vertex {
	return &vertex{
		value:            val,
		adjacentVertexes: []*vertex{},
	}
}
