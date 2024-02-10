package godsa

import (
	"fmt"

	"github.com/jttait/godsa/arrayminheap"
)

type Djikstras struct {
	graph Graph
}

func NewDjikstras(g Graph) ShortestPath {
	d := Djikstras{g}
	return &d
}

func (d *Djikstras) Calculate(node int) (map[int]int, error) {
	result := map[int]int{}
	visited := NewMapSet[int]()
	minHeap := arrayminheap.NewArrayMinHeap[int]()

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
