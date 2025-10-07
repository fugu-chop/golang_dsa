package datastructures

import (
	"math"
	"slices"
)

/*
`weightedGraph` is a type used to contain all `weightedVertex` types within a graph.
It enables the overall data structure to retain references to disconnected
vertices.
*/
type weightedGraph struct {
	vertices map[string]*weightedVertex
}

/*
WeightedGraph creates a pointer to a new `weightedGraph` type.
*/
func WeightedGraph() *weightedGraph {
	return &weightedGraph{
		vertices: make(map[string]*weightedVertex),
	}
}

/*
ListVertices returns a map of all the vertices contained within
a `weightedGraph` type, regardless of whether they are connected or not.
*/
func (g *weightedGraph) ListVertices() map[string]*weightedVertex {
	return g.vertices
}

/*
Vertex checks to see if a pointer to a `weightedVertex` with the value
of `val` already exists within the graph.

  - If not, it creates a pointer to a `weightedVertex` type with the value of `val`.

  - If the value already exists, a pointer to the existing `weightedVertex` type is returned
*/
func (g *weightedGraph) Vertex(val string) *weightedVertex {
	v, ok := g.vertices[val]
	if ok {
		return v
	}

	newVertex := &weightedVertex{
		value:            val,
		adjacentVertices: make(map[*weightedVertex]uint),
	}
	g.vertices[val] = newVertex
	return newVertex
}

/*
`weightedVertex` is a type used to implement a weighted graph data structure.
It contains a string `value` and a map of pointers to `weightedVertex` types
that represent edges between other vertices and their weight.
*/
type weightedVertex struct {
	value            string
	adjacentVertices map[*weightedVertex]uint
}

/*
AddDirectedVertex adds a single directed, weighted edge between
pointers to two `weightedVertex` types.
*/
func (v *weightedVertex) AddDirectedVertex(newVertex *weightedVertex, weight uint) {
	v.adjacentVertices[newVertex] = weight
}

/*
AddUndirectedVertex adds two undirected, weighted edges between
pointers to two `weightedVertex` types.
*/
func (v *weightedVertex) AddUndirectedVertex(newVertex *weightedVertex, weight uint) {
	// Avoid infinite loop
	if _, ok := v.adjacentVertices[newVertex]; ok {
		return
	}

	v.adjacentVertices[newVertex] = weight
	newVertex.AddUndirectedVertex(v, weight)
}

// Value returns the underlying string value of a weightedVertex
func (v *weightedVertex) Value() string {
	return v.value
}

/*
Weight returns the underlying weight of an edge between two `weightedVertex`
types. If an edge does not exist between the two `weightedVertex`,
#Weight returns 0.
*/
func (v *weightedVertex) Weight(vertex *weightedVertex) uint {
	weight, ok := v.adjacentVertices[vertex]
	if !ok {
		// Assumption that all weights are > 0
		return 0
	}
	return weight
}

/*
DFSWeighted performs depth-first search on a `searchVertex` against a provided `searchVal` string.
It returns a pointer to a `weightedVertex` type if:

  - A vertex exists with a value of `searchVal`; and

  - That vertex has an edge to the `searchVertex`

Otherwise it returns nil.
*/
func DFSWeighted(searchVertex *weightedVertex, searchVal string, visitedVertices map[string]bool) *weightedVertex {
	if searchVertex.value == searchVal {
		return searchVertex
	}

	visitedVertices[searchVertex.value] = true

	for vertex := range searchVertex.adjacentVertices {
		if visitedVertices[vertex.Value()] {
			continue
		}

		if vertex.value == searchVal {
			return vertex
		}

		if result := DFSWeighted(vertex, searchVal, visitedVertices); result != nil {
			return result
		}
	}

	return nil
}

/*
BFSWeighted performs breadth-first search on a `searchVertex` against a provided `searchVal` string.
It returns a pointer to a `weightedVertex` type if:

  - A vertex exists with a value of `searchVal`; and

  - That vertex has an edge to the `searchVertex`

Otherwise it returns nil.
*/
func BFSWeighted(searchVertex *weightedVertex, searchVal string) *weightedVertex {
	queue := Queue(searchVertex)
	visitedVertices := map[string]bool{
		searchVertex.value: true,
	}

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

		for vertex := range searchVertex.adjacentVertices {
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

/*
DijkstraShortestPath implements Dijkstra's algorithm between two `weightedVertex`
types.

It returns a slice of `weightedVertex` in the order of traversal to achieve
the lowest possible total value of edges between `start` and `end`.

If there is no sequence of edges between `start` and `end`, DijkstraShortestPath
returns nil.
*/
func DijkstraShortestPath(start, end *weightedVertex) []*weightedVertex {
	lowestCost := map[*weightedVertex]uint{start: 0}
	lowestPreviousVertex := make(map[*weightedVertex]*weightedVertex)
	unvisitedVertices := make(map[*weightedVertex]bool)
	visitedVertices := make(map[*weightedVertex]bool)
	currentVertex := start

	for currentVertex != nil {
		visitedVertices[currentVertex] = true
		delete(unvisitedVertices, currentVertex)

		for nextVertex, weight := range currentVertex.adjacentVertices {
			if _, ok := visitedVertices[nextVertex]; !ok {
				unvisitedVertices[nextVertex] = true
			}

			weightThroughVertex := lowestCost[currentVertex] + weight

			if _, ok := lowestCost[nextVertex]; !ok ||
				weightThroughVertex < lowestCost[nextVertex] {
				lowestCost[nextVertex] = weightThroughVertex
				lowestPreviousVertex[nextVertex] = currentVertex
			}
		}

		//	Find the next unvisited vertex to visit based on lowest
		//	weight from `start`
		var lowestWeight uint = math.MaxUint
		for vertex := range unvisitedVertices {
			weight, ok := lowestCost[vertex]
			if ok && weight < lowestWeight {
				lowestWeight = weight
				currentVertex = vertex
			}
		}
		// Terminate the loop if unvisitedVertices is empty
		if lowestWeight == math.MaxUint {
			currentVertex = nil
		}
	}

	// Construct the sequence of vertices based on lowestPreviousVertex
	shortestPath := []*weightedVertex{}
	currentVertex = end
	for currentVertex != start {
		shortestPath = append(shortestPath, currentVertex)
		// Check if there is no way to get to end from start
		result, ok := lowestPreviousVertex[currentVertex]
		if !ok {
			return nil
		}
		currentVertex = result
	}

	shortestPath = append(shortestPath, start)
	slices.Reverse(shortestPath)

	return shortestPath
}
