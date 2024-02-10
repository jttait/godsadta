// Package shortestpath provides an interface and implementations for the shortest path algorithm
package shortestpath

type ShortestPath interface {
	Calculate(node int) (map[int]int, error)
}
