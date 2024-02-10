package shortestpath

import (
	"fmt"

	"github.com/jttait/godsa/graph"
	"github.com/jttait/godsa/mapset"
	"github.com/jttait/godsa/minheap"
)

type Djikstras struct {
	graph graph.Graph
}

func NewDjikstras(g graph.Graph) ShortestPath {
	d := Djikstras{g}
	return &d
}

func (d *Djikstras) Calculate(node int) (map[int]int, error) {
	result := map[int]int{}
	visited := mapset.NewMapSet[int]()
	minHeap := minheap.NewArrayMinHeap[int]()

	visited.Insert(node)
	minHeap.Insert(node)
	result[node] = -1

	for minHeap.Size() > 0 {
		previousCost := result[node]
		node = minHeap.Extract()
		visited.Insert(node)
		result[node] = previousCost + 1
		neighbors, err := d.graph.Neighbors(node)
		if err != nil {
			return result, fmt.Errorf("Failed to get neighbors: %v.", err)
		}
		for _, neighbor := range neighbors.Iter() {
			if !visited.Contains(neighbor) {
				minHeap.Insert(neighbor)
			}
		}
	}
	return result, nil
}
