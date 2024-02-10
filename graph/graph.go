// Package graph provides interface and implementations for the graph data structure
package graph

import (
	"github.com/jttait/godsa/set"
)

type Graph interface {
	AddVertex(int) bool
	RemoveVertex(int) bool
	AddEdge(int, int) (bool, error)
	RemoveEdge(int, int) bool
	Neighbors(int) (set.Set[int], error)
	ContainsVertex(int) bool
}
