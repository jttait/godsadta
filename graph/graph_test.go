package graph

import (
	"testing"

	"github.com/jttait/godsa/set"
)

func TestShouldBeTrueWhenAddingNonExistingVertexToGraph(t *testing.T) {
	g := NewGraph()
	result := g.AddVertex(1)
	want := true
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseWhenAddingExistingVertexToGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	result := g.AddVertex(1)
	want := false
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldRemoveVertexFromGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.RemoveVertex(1)
	result := g.ContainsVertex(1)
	want := false
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnTrueWhenRemovingExistentVertexFromGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	result := g.RemoveVertex(1)
	want := true
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnFalseWhenRemovingNonExistentVertexFromGraph(t *testing.T) {
	g := NewGraph()
	result := g.RemoveVertex(1)
	want := false
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldRemoveEdgeWhenRemovingVertexFromGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	_ = g.RemoveVertex(1)
	result, _ := g.Neighbors(2)
	want := set.NewSet[int]()
	if !result.Equals(want) {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeTrueWhenAddingNonExistingEdgeToGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	result, _ := g.AddEdge(1, 2)
	want := true
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeFalseWhenAddingExistingEdgeToGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result, _ := g.AddEdge(1, 2)
	want := false
	if result != want {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldBeErrorWhenAddingEdgeBetweenNonExistingSourceVertex(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(2)
	_, err := g.AddEdge(1, 2)
	if err == nil {
		t.Fatalf("Want error. Got no error.")
	}
}

func TestShouldBeErrorWhenAddingEdgeBetweenNonExistingDestinationVertex(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_, err := g.AddEdge(1, 2)
	if err == nil {
		t.Fatalf("Want error. Got no error.")
	}
}

func TestShouldBeErrorWhenAddingEdgeBetweenTwoNonExistingVertices(t *testing.T) {
	g := NewGraph()
	_, err := g.AddEdge(1, 2)
	if err == nil {
		t.Fatalf("Want error. Got no error.")
	}
}

func TestShouldReturnEmptySetIfNoNeighbors(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	result, _ := g.Neighbors(1)
	want := set.NewSet[int]()
	if !result.Equals(want) {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnSetIfNeighbors(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result, _ := g.Neighbors(1)
	want := set.NewSet[int](2)
	if !result.Equals(want) {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnErrorIfVertexForNeighborsDoesNotExist(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	_, err := g.Neighbors(1)
	if err == nil {
		t.Fatalf("Got no error. Want error.")
	}
}

func TestShouldReturnNoErrorIfVertexForNeighborsDoesExist(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_, _ = g.AddEdge(1, 2)
	_, err := g.Neighbors(1)
	if err != nil {
		t.Fatalf("Got error. Want no error.")
	}
}

func TestShouldRemoveEdge(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	_ = g.RemoveEdge(1, 2)
	result, _ := g.Neighbors(1)
	want := set.NewSet[int]()
	if !result.Equals(want) {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnFalseWhenRemovingNonExistentEdge(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	result := g.RemoveEdge(1, 2)
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnTrueWhenRemovingExistentEdge(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result := g.RemoveEdge(1, 2)
	want := true
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnFalseWhenRemovingEdgeWithNonExistentVertices(t *testing.T) {
	g := NewGraph()
	result := g.RemoveEdge(1, 2)
	want := false
	if want != result {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}
