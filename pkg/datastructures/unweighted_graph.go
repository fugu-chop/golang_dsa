package datastructures

import (
	"slices"
)

/*
`unweightedGraph` is a type used to contain all `vertex` types within a graph.
It enables the overall data structure to retain references to disconnected
vertices.
*/
type unweightedGraph struct {
	vertices map[string]*vertex
}

/*
UnweightedGraph creates a pointer to a new `unweightedGraph` type.
*/
func UnweightedGraph() *unweightedGraph {
	return &unweightedGraph{
		vertices: make(map[string]*vertex),
	}
}

/*
ListVertices returns a map of all the vertices contained within
an `unweightedGraph` type, regardless of whether they are connected or not.
*/
func (g *unweightedGraph) ListVertices() map[string]*vertex {
	return g.vertices
}

/*
Vertex checks to see if a pointer to a `vertex` with the value
of `val` already exists within the graph.

  - If not, it creates a pointer to a `vertex` type with the value of `val`.

  - If the value already exists, a pointer to the existing `vertex` type is returned
*/
func (g *unweightedGraph) Vertex(val string) *vertex {
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
`vertex` is a type used to implement an unweighted graph data structure.
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
func DFSUnweighted(searchVertex *vertex, searchVal string, visitedVertices map[string]bool) *vertex {
	visitedVertices[searchVertex.Value()] = true

	if searchVertex.value == searchVal {
		return searchVertex
	}

	for _, vertex := range searchVertex.adjacentVertices {
		if visitedVertices[vertex.Value()] {
			continue
		}

		if vertex.Value() == searchVal {
			return vertex
		}

		result := DFSUnweighted(vertex, searchVal, visitedVertices)
		if result != nil && result.Value() == searchVal {
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
func BFSUnweighted(searchVertex *vertex, searchVal string) *vertex {
	queue := Queue(searchVertex)
	visitedVertices := map[string]bool{}

	_ = queue.Enqueue(searchVertex)
	visitedVertices[searchVertex.Value()] = true

	for {
		if _, err := queue.Read(); err != nil {
			return nil
		}

		searchVertex, err := queue.Dequeue()
		if err != nil {
			return nil
		}

		if searchVertex.Value() == searchVal {
			return searchVertex
		}

		for _, vertex := range searchVertex.adjacentVertices {
			if visitedVertices[vertex.Value()] {
				continue
			}

			if vertex.Value() == searchVal {
				return vertex
			}

			visitedVertices[vertex.Value()] = true

			_ = queue.Enqueue(vertex)
		}
	}
}
