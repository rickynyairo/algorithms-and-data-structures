package graphs

import "fmt"

type BFSAdjacencyMatrix struct {
	nodes            []GraphNode
	adjacency_matrix [][]int
}

func NewBFSAdjacencyMatrix(nodes []GraphNode) BFSAdjacencyMatrix {
	matrix := make([][]int, len(nodes))
	for i := range matrix {
		matrix[i] = make([]int, len(nodes))
	}
	return BFSAdjacencyMatrix{
		nodes:            nodes,
		adjacency_matrix: matrix,
	}
}

func (b *BFSAdjacencyMatrix) Bfs() {
	for _, node := range b.nodes {
		if !node.is_visited {
			b.bfsVisit(node)
		}
	}
}

func (b *BFSAdjacencyMatrix) bfsVisit(node GraphNode) {
	var queue []GraphNode
	queue = append(queue, node)
	for len(queue) > 0 {
		present_node := queue[0]
		queue = queue[1:]
		present_node.is_visited = true
		fmt.Print(node, " ")

		neighbours := b.getNeighbours(present_node)
		for _, neighbour := range neighbours {
			if !neighbour.is_visited {
				queue = append(queue, neighbour)
				neighbour.is_visited = true
			}
		}
	}
}

func (b *BFSAdjacencyMatrix) getNeighbours(node GraphNode) []GraphNode {
	var neighbours []GraphNode
	var node_index int = node.index
	for i := 0; i < len(b.adjacency_matrix); i++ {
		if b.adjacency_matrix[node_index][i] == 1 {
			neighbours = append(neighbours, b.nodes[i])
		}
	}
	return neighbours

}

func (b *BFSAdjacencyMatrix) AddUndirectedEdge(i int, j int) {
	i--
	j--
	b.adjacency_matrix[i][j] = 1
	b.adjacency_matrix[j][i] = 1
}
