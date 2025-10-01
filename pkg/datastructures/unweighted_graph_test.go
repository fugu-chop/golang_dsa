package datastructures

import (
	"testing"
)

func TestVertex(t *testing.T) {
	t.Parallel()

	t.Run("creates a vertex within a graph", func(t *testing.T) {
		t.Parallel()

		g := UnweightedGraph()

		_ = g.Vertex("home")
		_ = g.Vertex("a")

		got := len(g.ListVertices())
		want := 2
		if got != want {
			t.Fatalf("expected graph to have %d vertices, got: %d", want, got)
		}
	})

	t.Run("does not duplicate vertices when value is same", func(t *testing.T) {
		g := UnweightedGraph()

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

	g := UnweightedGraph()

	vertexValue := "home"
	v := g.Vertex(vertexValue)

	if v.Value() != vertexValue {
		t.Fatalf("expected vertex to have value of %s, got: %s", vertexValue, v.Value())
	}
}

func TestAddDirectedVertex(t *testing.T) {
	t.Parallel()

	g := UnweightedGraph()

	v := g.Vertex("home")
	a := g.Vertex("a")
	b := g.Vertex("b")
	v.AddDirectedVertex(a)
	v.AddDirectedVertex(b)

	connectedVertexA := DFSUnweighted(v, "a", map[string]bool{})
	if connectedVertexA == nil {
		t.Fatal("expected home vertex to have edge to a")
	}

	connectedVertexB := DFSUnweighted(v, "b", map[string]bool{})
	if connectedVertexB == nil {
		t.Fatal("expected home vertex to have edge to b")
	}

	unconnectedVertex := DFSUnweighted(a, "home", map[string]bool{})
	if unconnectedVertex != nil {
		t.Fatalf("expected vertex to have no edge, got: %v", unconnectedVertex)
	}

	unconnectedVertex = DFSUnweighted(b, "home", map[string]bool{})
	if unconnectedVertex != nil {
		t.Fatalf("expected vertex to have no edge, got: %v", unconnectedVertex)
	}
}

func TestAddUndirectedVertex(t *testing.T) {
	t.Parallel()

	g := UnweightedGraph()

	v := g.Vertex("home")
	a := g.Vertex("a")
	b := g.Vertex("b")
	v.AddUndirectedVertex(a)
	v.AddUndirectedVertex(b)

	connectedVertexA := DFSUnweighted(v, "a", map[string]bool{})
	if connectedVertexA == nil {
		t.Fatal("expected home vertex to have edge to a")
	}

	connectedVertexB := DFSUnweighted(v, "b", map[string]bool{})
	if connectedVertexB == nil {
		t.Fatal("expected home vertex to have edge to b")
	}

	connectedVertex := DFSUnweighted(a, "home", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex a to have edge to home")
	}

	connectedVertex = DFSUnweighted(b, "home", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex b to have edge to home")
	}

	connectedVertex = DFSUnweighted(a, "b", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex a to have edge to b")
	}

	connectedVertex = DFSUnweighted(b, "a", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex b to have edge to a")
	}
}

func TestDFS(t *testing.T) {
	t.Parallel()

	graph := UnweightedGraph()

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

	got := DFSUnweighted(a, "f", map[string]bool{})
	if got != f {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", f, got)
	}

	got = DFSUnweighted(d, "f", map[string]bool{})
	if got != nil {
		t.Fatal("expected no vertex")
	}

	got = DFSUnweighted(d, "d", map[string]bool{})
	if got != d {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", d, got)
	}
}

func TestBFS(t *testing.T) {
	t.Parallel()

	graph := UnweightedGraph()

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

	got := BFSUnweighted(a, "f")
	if got != f {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", f, got)
	}

	got = BFSUnweighted(d, "f")
	if got != nil {
		t.Fatal("expected no vertex")
	}

	got = BFSUnweighted(d, "d")
	if got != d {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", d, got)
	}
}
