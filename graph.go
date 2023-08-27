package godsa

import "fmt"

type Graph struct {
	adjacencyList map[int]*Set
}

func NewGraph() *Graph {
	g := Graph{}
	g.adjacencyList = make(map[int]*Set)
	return &g
}

func (g *Graph) AddVertex(i int) bool {
	_, ok := g.adjacencyList[i]
	if !ok {
		g.adjacencyList[i] = NewSet()
		return true
	}
	return false
}

func (g *Graph) AddEdge(i, j int) (bool, error) {
	_, ok := g.adjacencyList[i]
	if !ok {
		return false, fmt.Errorf("Vertex %v not in graph.", i)
	}
	_, ok = g.adjacencyList[j]
	if !ok {
		return false, fmt.Errorf("Vertex %v not in graph.", j)
	}
	result := g.adjacencyList[i].Add(j)
	return result, nil
}

func (g *Graph) Neighbours(i int) (*Set, error) {
	_, ok := g.adjacencyList[i]
	if !ok {
		return NewSet(), fmt.Errorf("Vertex %v not in graph.", i)
	}
	return g.adjacencyList[i], nil
}
