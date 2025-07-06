package datastructures

import "slices"

/*
`vertex` is a type used to implement a larger graph data structure.
It contains a string `value` and a slice of pointers to `vertex` types
that represent edges between other vertexes.
*/
type vertex struct {
	value            string
	adjacentVertexes []*vertex
}

/*
AddDirectedVertex adds a single directed, unweighted edge between
pointers to two `vertex` types.
*/
func (v *vertex) AddDirectedVertex(newVertex *vertex) {
	v.adjacentVertexes = append(v.adjacentVertexes, newVertex)
}

/*
AddUndirectedVertex adds two undirected, unweighted edges between
pointers to two `vertex` types.
*/
func (v *vertex) AddUndirectedVertex(newVertex *vertex) {
	// Avoid infinite loop
	if slices.Contains(v.adjacentVertexes, newVertex) {
		return
	}
	v.adjacentVertexes = append(v.adjacentVertexes, newVertex)
	newVertex.AddUndirectedVertex(v)
}

// NewVertex creates a pointer to a `vertex` type with the value of `val`.
func NewVertex(val string) *vertex {
	return &vertex{
		value:            val,
		adjacentVertexes: []*vertex{},
	}
}
