package graph

import "github.com/jttait/godsa"

type Graph interface {
	AddVertex(int) bool
	RemoveVertex(int) bool
	AddEdge(int, int) (bool, error)
	RemoveEdge(int, int) bool
	Neighbors(int) (godsa.Set[int], error)
	ContainsVertex(int) bool
}
