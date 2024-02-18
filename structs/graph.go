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

func (graph *Graph[T]) AddEdge(a, b int) {
	graph.Edges[a] = append(graph.Edges[a], b)
	graph.Edges[b] = append(graph.Edges[b], a)
}

func (graph *Graph[T]) AddDirectedEdge(a, b int) {
	graph.Edges[a] = append(graph.Edges[a], b)
}
