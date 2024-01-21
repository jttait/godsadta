package godsa

type ShortestPath interface {
	Calculate(node int) (map[int]int, error)
}
