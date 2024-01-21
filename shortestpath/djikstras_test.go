package shortestpath

import (
	"testing"

	"github.com/jttait/godsa"
)

func TestShouldBeOneResultForTwoNodeGraph(t *testing.T) {
	g := godsa.NewMapSetGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddEdge(1, 2)
	result, err := djikstras(g, 1)
	if err != nil {
		t.Fatalf("Error not nil: %v.\n", err)
	}
	val, ok := result[2]
	if !ok {
		t.Fatal("Want true. Got false")
	}
	if val != 1 {
		t.Fatalf("Want: 1. Got %v.\n", val)
	}
}

func TestShouldBeTwoResultForThreeNodeGraph(t *testing.T) {
	g := godsa.NewMapSetGraph()
	g.AddVertex(1)
	g.AddVertex(2)
	g.AddVertex(3)
	g.AddEdge(1, 2)
	g.AddEdge(2, 3)
	result, _ := djikstras(g, 1)
	val, ok := result[3]
	if !ok {
		t.Fatal("Want true. Got false")
	}
	if val != 2 {
		t.Fatalf("Want: 1. Got %v.\n", val)
	}
}
