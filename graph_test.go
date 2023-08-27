package godsa

import "testing"

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

func TestShouldReturnEmptySetIfNoNeighbours(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	result, _ := g.Neighbours(1)
	want := NewSet[int]()
	if !result.Equals(want) {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnSetIfNeighbours(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result, _ := g.Neighbours(1)
	want := NewSet[int]()
	want.Add(2)
	if !result.Equals(want) {
		t.Fatalf("Got %v. Want %v.\n", result, want)
	}
}

func TestShouldReturnErrorIfVertexForNeighboursDoesNotExist(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	_, err := g.Neighbours(1)
	if err == nil {
		t.Fatalf("Got no error. Want error.")
	}
}

func TestShouldReturnNoErrorIfVertexForNeighboursDoesExist(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_, _ = g.AddEdge(1, 2)
	_, err := g.Neighbours(1)
	if err != nil {
		t.Fatalf("Got error. Want no error.")
	}
}
