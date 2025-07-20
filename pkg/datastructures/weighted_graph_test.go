package datastructures_test

import (
	"math/rand"
	"testing"

	"github.com/fugu-chop/golang_dsa/pkg/datastructures"
)

func TestWeightedVertex(t *testing.T) {
	t.Parallel()

	t.Run("creates a vertex within a graph", func(t *testing.T) {
		t.Parallel()

		g := datastructures.WeightedGraph()

		_ = g.Vertex("home")
		_ = g.Vertex("a")

		got := len(g.ListVertices())
		want := 2
		if got != want {
			t.Fatalf("expected graph to have %d vertices, got: %d", want, got)
		}
	})

	t.Run("does not duplicate vertices when value is same", func(t *testing.T) {
		g := datastructures.WeightedGraph()

		_ = g.Vertex("home")
		_ = g.Vertex("home")

		got := len(g.ListVertices())
		want := 1
		if got != 1 {
			t.Fatalf("expected graph to have %d vertices, got: %d", want, got)
		}
	})
}

func TestWeightedValue(t *testing.T) {
	t.Parallel()

	g := datastructures.WeightedGraph()

	vertexValue := "home"
	v := g.Vertex(vertexValue)

	if v.Value() != vertexValue {
		t.Fatalf("expected vertext to have value of %s, got: %s", vertexValue, v.Value())
	}
}

func TestWeightedWeight(t *testing.T) {
	t.Parallel()

	t.Run("works for directed vertices", func(t *testing.T) {
		t.Parallel()

		g := datastructures.WeightedGraph()

		v := g.Vertex("home")
		x := g.Vertex("away")

		expected := uint(rand.Intn(10))
		v.AddDirectedVertex(x, expected)

		if v.Weight(x) != expected {
			t.Fatalf("expected vertex to have weight of %d, got: %d", expected, v.Weight(x))
		}
	})

	t.Run("works for undirected vertices", func(t *testing.T) {
		t.Parallel()

		g := datastructures.WeightedGraph()

		v := g.Vertex("home")
		x := g.Vertex("away")

		expected := uint(rand.Intn(10))
		v.AddUndirectedVertex(x, expected)

		if v.Weight(x) != expected {
			t.Fatalf("expected vertex to have weight of %d, got: %d", expected, v.Weight(x))
		}
	})

	t.Run("returns -1 if no edge exists", func(t *testing.T) {
		t.Parallel()

		g := datastructures.WeightedGraph()

		v := g.Vertex("home")
		x := g.Vertex("away")

		if v.Weight(x) != 0 {
			t.Fatalf("expected vertex to have weight of -1, got: %d", v.Weight(x))
		}
	})
}

func TestAddUndirectedWeightedVertex(t *testing.T) {
	t.Parallel()

	g := datastructures.WeightedGraph()

	v := g.Vertex("home")
	a := g.Vertex("a")
	b := g.Vertex("b")
	v.AddUndirectedVertex(a, uint(rand.Intn(10)))
	v.AddUndirectedVertex(b, uint(rand.Intn(10)))

	connectedVertexA := datastructures.DFSWeighted(v, "a", map[string]bool{})
	if connectedVertexA == nil {
		t.Fatal("expected home vertex to have edge to a")
	}

	connectedVertexB := datastructures.DFSWeighted(v, "b", map[string]bool{})
	if connectedVertexB == nil {
		t.Fatal("expected home vertex to have edge to b")
	}

	connectedVertex := datastructures.DFSWeighted(a, "home", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex a to have edge to home")
	}

	connectedVertex = datastructures.DFSWeighted(b, "home", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex b to have edge to home")
	}

	connectedVertex = datastructures.DFSWeighted(a, "b", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex a to have edge to b")
	}

	connectedVertex = datastructures.DFSWeighted(b, "a", map[string]bool{})
	if connectedVertex == nil {
		t.Fatal("expected vertex b to have edge to a")
	}
}

func TestWeightedDFS(t *testing.T) {
	t.Parallel()

	graph := datastructures.WeightedGraph()

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

	a.AddDirectedVertex(b, uint(rand.Intn(10)))
	a.AddDirectedVertex(d, uint(rand.Intn(10)))
	b.AddDirectedVertex(c, uint(rand.Intn(10)))
	c.AddDirectedVertex(d, uint(rand.Intn(10)))
	a.AddDirectedVertex(e, uint(rand.Intn(10)))
	a.AddDirectedVertex(g, uint(rand.Intn(10)))
	e.AddDirectedVertex(g, uint(rand.Intn(10)))
	e.AddDirectedVertex(f, uint(rand.Intn(10)))
	f.AddDirectedVertex(g, uint(rand.Intn(10)))

	got := datastructures.DFSWeighted(a, "f", map[string]bool{})
	if got != f {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", f, got)
	}

	got = datastructures.DFSWeighted(d, "f", map[string]bool{})
	if got != nil {
		t.Fatal("expected no vertex")
	}

	got = datastructures.DFSWeighted(d, "d", map[string]bool{})
	if got != d {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", d, got)
	}
}

func TestWeightedBFS(t *testing.T) {
	t.Parallel()

	graph := datastructures.WeightedGraph()

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

	a.AddDirectedVertex(b, uint(rand.Intn(10)))
	a.AddDirectedVertex(d, uint(rand.Intn(10)))
	b.AddDirectedVertex(c, uint(rand.Intn(10)))
	c.AddDirectedVertex(d, uint(rand.Intn(10)))
	a.AddDirectedVertex(e, uint(rand.Intn(10)))
	a.AddDirectedVertex(g, uint(rand.Intn(10)))
	e.AddDirectedVertex(g, uint(rand.Intn(10)))
	e.AddDirectedVertex(f, uint(rand.Intn(10)))
	f.AddDirectedVertex(g, uint(rand.Intn(10)))

	got := datastructures.BFSWeighted(a, "f")
	if got != f {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", f, got)
	}

	got = datastructures.BFSWeighted(d, "f")
	if got != nil {
		t.Fatal("expected no vertex")
	}

	got = datastructures.BFSWeighted(d, "d")
	if got != d {
		t.Fatalf("expected to return vertex with value of: %v, got: %v", d, got)
	}
}

func TestDijkstraShortestPath(t *testing.T) {
	t.Parallel()

	t.Run("returns nil if no edge exists", func(t *testing.T) {
		t.Parallel()

		graph := datastructures.WeightedGraph()

		a := graph.Vertex("a")
		b := graph.Vertex("b")

		result := datastructures.DijkstraShortestPath(a, b)
		if result != nil {
			t.Fatalf("expected no result from unconnected vertices, got: %v", result)
		}
	})
}
