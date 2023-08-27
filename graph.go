package godsa

import "fmt"

// Graph is a data structure comprising vertices and edges connecting these vertices.
type Graph struct {
	adjacencyList map[int]*Set[int]
}

// NewGraph instantiates a graph and returns a pointer to it.
func NewGraph() *Graph {
	g := Graph{}
	g.adjacencyList = make(map[int]*Set[int])
	return &g
}

// AddVertex adds a vertex to the graph. It returns a Boolean that is true if a vertex with this
// value did not already exist in the graph. It's not possible to have two
// vertices with the same value.
func (g *Graph) AddVertex(i int) bool {
	_, ok := g.adjacencyList[i]
	if !ok {
		g.adjacencyList[i] = NewSet[int]()
		return true
	}
	return false
}

// AddEdge adds an edge between two vertices. It returns an error if either of the vertices do not
// exist in the graph. A Boolean is returned that is true if this edge did not already exist in the
// graph.
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

// Neighbors returns a set of nodes that are connected to the current node by an edge. It returns an
// error if the vertex given does not exist in the graph.
func (g *Graph) Neighbors(i int) (*Set[int], error) {
	_, ok := g.adjacencyList[i]
	if !ok {
		return NewSet[int](), fmt.Errorf("Vertex %v not in graph.", i)
	}
	return g.adjacencyList[i], nil
}
