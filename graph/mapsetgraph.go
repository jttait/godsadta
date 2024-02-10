package graph

import (
	"fmt"

	"github.com/jttait/godsa/set"
)

// Graph is a data structure comprising vertices and edges connecting these vertices.
type MapSetGraph struct {
	adjacencyList map[int]set.Set[int]
}

// NewGraph instantiates a graph and returns a pointer to it.
func NewMapSetGraph() *MapSetGraph {
	g := MapSetGraph{}
	g.adjacencyList = make(map[int]set.Set[int])
	return &g
}

// AddVertex adds a vertex to the graph. It returns a Boolean that is true if a vertex with this
// value did not already exist in the graph. It's not possible to have two
// vertices with the same value.
func (g *MapSetGraph) AddVertex(i int) bool {
	_, ok := g.adjacencyList[i]
	if !ok {
		g.adjacencyList[i] = set.NewMapSet[int]()
		return true
	}
	return false
}

// RemoveVertex removes a vertex from the graph. It returns a Boolean that is true if the vertex
// existed in the graph. This method also removes any edges that have the removed node as their
// source or destination.
func (g *MapSetGraph) RemoveVertex(i int) bool {
	_, ok := g.adjacencyList[i]
	if !ok {
		return false
	}
	delete(g.adjacencyList, i)
	return true
}

// AddEdge adds an edge between two vertices. It returns an error if either of the vertices do not
// exist in the graph. A Boolean is returned that is true if this edge did not already exist in the
// graph.
func (g *MapSetGraph) AddEdge(i, j int) (bool, error) {
	_, ok := g.adjacencyList[i]
	if !ok {
		return false, fmt.Errorf("Vertex %v not in graph.", i)
	}
	_, ok = g.adjacencyList[j]
	if !ok {
		return false, fmt.Errorf("Vertex %v not in graph.", j)
	}
	result := g.adjacencyList[i].Insert(j)
	return result, nil
}

// RemoveEdge removes an edge between two vertices. It returns a Boolean that is true if the edge
// existed in the graph.
func (g *MapSetGraph) RemoveEdge(i, j int) bool {
	_, oki := g.adjacencyList[i]
	_, okj := g.adjacencyList[j]
	if !oki || !okj {
		return false
	}
	if !g.adjacencyList[i].Contains(j) {
		return false
	}
	g.adjacencyList[i].Remove(j)
	return true
}

// Neighbors returns a set of nodes that are connected to the current node by an edge. It returns an
// error if the vertex given does not exist in the graph.
func (g *MapSetGraph) Neighbors(i int) (set.Set[int], error) {
	_, ok := g.adjacencyList[i]
	if !ok {
		return set.NewMapSet[int](), fmt.Errorf("Vertex %v not in graph.", i)
	}
	return g.adjacencyList[i], nil
}

// ContainsVertex returns true if the graph contains this vertex.
func (g *MapSetGraph) ContainsVertex(i int) bool {
	_, ok := g.adjacencyList[i]
	return ok
}
