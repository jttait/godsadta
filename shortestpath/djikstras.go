package shortestpath

import (
	"fmt"

	"github.com/jttait/godsa"
)

func djikstras(graph godsa.Graph, node int) (map[int]int, error) {
	result := map[int]int{}
	visited := godsa.NewMapSet[int]()
	minHeap := godsa.NewArrayMinHeap[int]()

	visited.Insert(node)
	minHeap.Insert(node)
	result[node] = -1

	for minHeap.Size() > 0 {
		previousCost := result[node]
		node = minHeap.Extract()
		visited.Insert(node)
		result[node] = previousCost + 1
		neighbors, err := graph.Neighbors(node)
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
