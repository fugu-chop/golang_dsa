package datastructures_test

import (
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestVertex(t *testing.T) {
	t.Parallel()

	t.Run("creates a vertex within a graph", func(t *testing.T) {
		t.Parallel()

		g := datastructures.UnweightedGraph()

		_ = g.Vertex("home")
		_ = g.Vertex("a")

		got := len(g.ListVertices())
		want := 2
		if got != want {
			t.Fatalf("expected graph to have %d vertices, got: %d", want, got)
		}
	})

	t.Run("does not duplicate vertices when value is same", func(t *testing.T) {
		g := datastructures.UnweightedGraph()

		_ = g.Vertex("home")
		_ = g.Vertex("home")

		got := len(g.ListVertices())
		want := 1
		if got != 1 {
			t.Fatalf("expected graph to have %d vertices, got: %d", want, got)
		}
	})
}

func TestValue(t *testing.T) {
	t.Parallel()

	g := datastructures.UnweightedGraph()

	vertexValue := "home"
	v := g.Vertex(vertexValue)

	if v.Value() != vertexValue {
		t.Fatalf("expected vertext to have value of %s, got: %s", vertexValue, v.Value())
	}
}

func TestAddDirectedVertex(t *testing.T) {
	t.Parallel()

	g := datastructures.UnweightedGraph()

	v := g.Vertex("home")
	a := g.Vertex("a")
	b := g.Vertex("b")
	v.AddDirectedVertex(a)
	v.AddDirectedVertex(b)

	connectedVertexA := datastructures.DFS(v, "a", map[string]bool{})
	if connectedVertexA == nil {
		t.Fatal("expected home vertex to have edge to a")
	}

	connectedVertexB := datastructures.DFS(v, "b", map[string]bool{})
	if connectedVertexB == nil {
		t.Fatal("expected home vertex to have edge to b")
	}

	unconnectedVertex := datastructures.DFS(a, "home", map[string]bool{})
	if unconnectedVertex != nil {
		t.Fatalf("expected vertex to have no edge, got: %v", unconnectedVertex)
	}

	unconnectedVertex = datastructures.DFS(b, "home", map[string]bool{})
	if unconnectedVertex != nil {
		t.Fatalf("expected vertex to have no edge, got: %v", unconnectedVertex)
	}
}

func TestAddUndirectedVertex(t *testing.T) {
	t.Parallel()

	g := datastructures.UnweightedGraph()

	v := g.Vertex("home")
	a := g.Vertex("a")
	b := g.Vertex("b")
	v.AddUndirectedVertex(a)
	v.AddUndirectedVertex(b)

	connectedVertexA := datastructures.DFS(v, "a", map[string]bool{})
	if connectedVertexA == nil {
		t.Fatal("expected home vertex to have edge to a")
	}

	connectedVertexB := datastructures.DFS(v, "b", map[string]bool{})
	if connectedVertexB == nil {
		t.Fatal("expected home vertex to have edge to b")
	}

	connectedVertex := datastructures.DFS(a, "home", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex a to have edge to home")
	}

	connectedVertex = datastructures.DFS(b, "home", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex b to have edge to home")
	}

	connectedVertex = datastructures.DFS(a, "b", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex a to have edge to b")
	}

	connectedVertex = datastructures.DFS(b, "a", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex b to have edge to a")
	}
}

func TestDFS(t *testing.T) {
	t.Parallel()

	graph := datastructures.UnweightedGraph()

	a := graph.Vertex("a")
	b := graph.Vertex("b")
	c := graph.Vertex("c")
	d := graph.Vertex("d")
	e := graph.Vertex("e")
	f := graph.Vertex("f")
	g := graph.Vertex("g")

	/*
		graph should look like:
			   a
			/ | \ \
		   b  d  e-g
		  /   |  | /
		 c-----  f/
	*/

	a.AddDirectedVertex(b)
	a.AddDirectedVertex(d)
	b.AddDirectedVertex(c)
	c.AddDirectedVertex(d)
	a.AddDirectedVertex(e)
	a.AddDirectedVertex(g)
	e.AddDirectedVertex(g)
	e.AddDirectedVertex(f)
	f.AddDirectedVertex(g)

	got := datastructures.DFS(a, "f", map[string]bool{})
	if got != f {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", f, got)
	}

	got = datastructures.DFS(d, "f", map[string]bool{})
	if got != nil {
		t.Fatal("expected no vertex")
	}

	got = datastructures.DFS(d, "d", map[string]bool{})
	if got != d {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", d, got)
	}
}

func TestBFS(t *testing.T) {
	t.Parallel()

	graph := datastructures.UnweightedGraph()

	a := graph.Vertex("a")
	b := graph.Vertex("b")
	c := graph.Vertex("c")
	d := graph.Vertex("d")
	e := graph.Vertex("e")
	f := graph.Vertex("f")
	g := graph.Vertex("g")

	/*
		graph should look like:
			   a
			/ | \ \
		   b  d  e-g
		  /   |  | /
		 c-----  f/
	*/

	a.AddDirectedVertex(b)
	a.AddDirectedVertex(d)
	b.AddDirectedVertex(c)
	c.AddDirectedVertex(d)
	a.AddDirectedVertex(e)
	a.AddDirectedVertex(g)
	e.AddDirectedVertex(g)
	e.AddDirectedVertex(f)
	f.AddDirectedVertex(g)

	got := datastructures.BFS(a, "f")
	if got != f {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", f, got)
	}

	got = datastructures.BFS(d, "f")
	if got != nil {
		t.Fatal("expected no vertex")
	}

	got = datastructures.BFS(d, "d")
	if got != d {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", d, got)
	}
}
