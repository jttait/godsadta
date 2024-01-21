package godsa

type Graph interface {
	AddVertex(int) bool
	RemoveVertex(int) bool
	AddEdge(int, int) (bool, error)
	RemoveEdge(int, int) bool
	Neighbors(int) (Set[int], error)
	ContainsVertex(int) bool
}
