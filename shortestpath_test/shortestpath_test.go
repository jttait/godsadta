package shortestpath_test

import (
	"testing"

	"github.com/jttait/godsa/djikstras"
	"github.com/jttait/godsa/graph"
	"github.com/jttait/godsa/mapsetgraph"
	"github.com/jttait/godsa/shortestpath"
)

func getShortestPathImplementations() [](func(graph.Graph) shortestpath.ShortestPath) {
	return [](func(graph.Graph) shortestpath.ShortestPath){
		djikstras.NewDjikstras,
	}
}

func TestShouldBeOneResultForTwoNodeGraph(t *testing.T) {
	for _, s := range getShortestPathImplementations() {
		g := mapsetgraph.NewMapSetGraph()
		g.AddVertex(1)
		g.AddVertex(2)
		g.AddEdge(1, 2)
		shortestpath := s(g)
		result, err := shortestpath.Calculate(1)
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
}

func TestShouldBeTwoResultForThreeNodeGraph(t *testing.T) {
	for _, s := range getShortestPathImplementations() {
		g := mapsetgraph.NewMapSetGraph()
		g.AddVertex(1)
		g.AddVertex(2)
		g.AddVertex(3)
		g.AddEdge(1, 2)
		g.AddEdge(2, 3)
		shortestpath := s(g)
		result, _ := shortestpath.Calculate(1)
		val, ok := result[3]
		if !ok {
			t.Fatal("Want true. Got false")
		}
		if val != 2 {
			t.Fatalf("Want: 1. Got %v.\n", val)
		}
	}
}
