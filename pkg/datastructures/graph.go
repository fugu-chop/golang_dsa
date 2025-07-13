package datastructures

import (
	"slices"
)

/*
`graph` is a type used to contain all `vertex` types within a graph.
It enables the overall data structure to retain references to disconnected
vertices.
*/
type graph struct {
	vertices map[string]*vertex
}

/*
NewGraph creates a pointer to a new `graph` type.
*/
func Graph() *graph {
	return &graph{
		vertices: make(map[string]*vertex),
	}
}

/*
ListVertices returns a map of all the vertices contained within
a `graph` type, regardless of whether they are connected or not.
*/
func (g *graph) ListVertices() map[string]*vertex {
	return g.vertices
}

/*
Vertex checks to see if a pointer to a `vertex` with the value
of `val` already exists within the graph.

  - If not, it creates a pointer to a `vertex` type with the value of `val`.

  - If the value already exists, a pointer to the existing `vertex` type is returned
*/
func (g *graph) Vertex(val string) *vertex {
	v, ok := g.vertices[val]
	if ok {
		return v
	}

	newVertex := &vertex{
		value:            val,
		adjacentVertices: []*vertex{},
	}
	g.vertices[val] = newVertex
	return newVertex
}

/*
`vertex` is a type used to implement a larger graph data structure.
It contains a string `value` and a slice of pointers to `vertex` types
that represent edges between other vertices.
*/
type vertex struct {
	value            string
	adjacentVertices []*vertex
}

/*
AddDirectedVertex adds a single directed, unweighted edge between
pointers to two `vertex` types.
*/
func (v *vertex) AddDirectedVertex(newVertex *vertex) {
	v.adjacentVertices = append(v.adjacentVertices, newVertex)
}

/*
AddUndirectedVertex adds two undirected, unweighted edges between
pointers to two `vertex` types.
*/
func (v *vertex) AddUndirectedVertex(newVertex *vertex) {
	// Avoid infinite loop
	if slices.Contains(v.adjacentVertices, newVertex) {
		return
	}
	v.adjacentVertices = append(v.adjacentVertices, newVertex)
	newVertex.AddUndirectedVertex(v)
}

// Value returns the underlying string value of a vertex
func (v *vertex) Value() string {
	return v.value
}

/*
DFS performs depth-first search on a `searchVertex` against a provided `searchVal` string.
It returns a pointer to a `vertex` type if:

  - A vertex exists with a value of `searchVal`; and

  - That vertex has an edge to the `searchVertex`

Otherwise it returns nil.
*/
func DFS(searchVertex *vertex, searchVal string, visitedVertices map[string]bool) *vertex {
	visitedVertices[searchVertex.value] = true

	if searchVertex.value == searchVal {
		return searchVertex
	}

	for _, vertex := range searchVertex.adjacentVertices {
		if visitedVertices[vertex.value] {
			continue
		}

		if vertex.value == searchVal {
			return vertex
		}

		result := DFS(vertex, searchVal, visitedVertices)
		if result != nil && result.value == searchVal {
			return result
		}
	}

	return nil
}

/*
BFS performs breadth-first search on a `searchVertex` against a provided `searchVal` string.
It returns a pointer to a `vertex` type if:

  - A vertex exists with a value of `searchVal`; and

  - That vertex has an edge to the `searchVertex`

Otherwise it returns nil.
*/
func BFS(searchVertex *vertex, searchVal string) *vertex {
	// TODO: Replace queue with generic Queue type
	queue := []*vertex{}
	visitedVertices := map[string]bool{}

	queue = append(queue, searchVertex)
	visitedVertices[searchVertex.value] = true

	for len(queue) > 0 {
		searchVertex = queue[0]
		queue = queue[1:]

		if searchVertex.value == searchVal {
			return searchVertex
		}

		for _, vertex := range searchVertex.adjacentVertices {
			if visitedVertices[vertex.value] {
				continue
			}

			if vertex.value == searchVal {
				return vertex
			}

			visitedVertices[vertex.value] = true

			queue = append(queue, vertex)
		}
	}

	return nil
}
