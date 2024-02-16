package structs

type Graph[T any] struct {
	Nodes []T
	Edges map[int][]int
}

func (graph *Graph[T]) New() {
	graph.Edges = make(map[int][]int)
}

func (graph *Graph[T]) AddNode(node T) {
	graph.Nodes = append(graph.Nodes, node)
}
