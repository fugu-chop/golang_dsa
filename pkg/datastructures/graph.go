package datastructures

import "slices"

/*
`graph` is a type used to contain all `vertex` types within a graph.
It enables the overall data structure to retain references to disconnected
vertexes.
*/
type graph struct {
	vertexes map[string]*vertex
}

/*
NewGraph creates a pointer to a new `graph` type.
*/
func Graph() *graph {
	return &graph{
		vertexes: make(map[string]*vertex),
	}
}

/*
Vertex checks to see if a pointer to a `vertex` with the value
of `val` already exists.

  - If not, it creates a pointer to a `vertex` type with the value of `val`.

  - If the value already exists, a pointer to the existing `vertex` type is returned
*/
func (g *graph) Vertex(val string) *vertex {
	v, ok := g.vertexes[val]
	if ok {
		return v
	}

	newVertex := &vertex{
		value:            val,
		adjacentVertexes: []*vertex{},
	}
	g.vertexes[val] = newVertex
	return newVertex
}

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
