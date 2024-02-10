package shortestpath

type ShortestPath interface {
	Calculate(node int) (map[int]int, error)
}
