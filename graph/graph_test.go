package graph

import (
	"testing"

	"github.com/jttait/godsa/assert"
	"github.com/jttait/godsa/set"
)

func TestShouldBeTrueWhenAddingNonExistingVertexToGraph(t *testing.T) {
	g := NewGraph()
	result := g.AddVertex(1)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseWhenAddingExistingVertexToGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	result := g.AddVertex(1)
	assert.AssertFalse(result, t)
}

func TestShouldRemoveVertexFromGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.RemoveVertex(1)
	result := g.ContainsVertex(1)
	assert.AssertFalse(result, t)
}

func TestShouldReturnTrueWhenRemovingExistentVertexFromGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	result := g.RemoveVertex(1)
	assert.AssertTrue(result, t)
}

func TestShouldReturnFalseWhenRemovingNonExistentVertexFromGraph(t *testing.T) {
	g := NewGraph()
	result := g.RemoveVertex(1)
	assert.AssertFalse(result, t)
}

func TestShouldRemoveEdgeWhenRemovingVertexFromGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	_ = g.RemoveVertex(1)
	result, _ := g.Neighbors(2)
	want := set.NewMapSet[int]()
	assert.AssertTrue(result.Equals(want), t)
}

func TestShouldBeTrueWhenAddingNonExistingEdgeToGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	result, _ := g.AddEdge(1, 2)
	assert.AssertTrue(result, t)
}

func TestShouldBeFalseWhenAddingExistingEdgeToGraph(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result, _ := g.AddEdge(1, 2)
	assert.AssertFalse(result, t)
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
	result, _ = result.(*set.MapSet[int])
	want := set.NewMapSet[int]()
	assert.AssertTrue(result.Equals(want), t)
}

func TestShouldReturnSetIfNeighbors(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result, _ := g.Neighbors(1)
	want := set.NewMapSet[int](2)
	assert.AssertTrue(result.Equals(want), t)
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
	result, _ = result.(*set.MapSet[int])
	want := set.NewMapSet[int]()
	assert.AssertTrue(result.Equals(want), t)
}

func TestShouldReturnFalseWhenRemovingNonExistentEdge(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	result := g.RemoveEdge(1, 2)
	assert.AssertFalse(result, t)
}

func TestShouldReturnTrueWhenRemovingExistentEdge(t *testing.T) {
	g := NewGraph()
	_ = g.AddVertex(1)
	_ = g.AddVertex(2)
	_, _ = g.AddEdge(1, 2)
	result := g.RemoveEdge(1, 2)
	assert.AssertTrue(result, t)
}

func TestShouldReturnFalseWhenRemovingEdgeWithNonExistentVertices(t *testing.T) {
	g := NewGraph()
	result := g.RemoveEdge(1, 2)
	assert.AssertFalse(result, t)
}
