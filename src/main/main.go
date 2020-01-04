package main

import "fmt"
import "sorts"
import . "graphs"

func mains() {
	a := []int{2, 212, 3001, 14, 501, 7800, 9932, 33, 45, 45, 45, 91, 99, 37, 102, 102, 104, 106, 109, 106}
	var b []int
	var c []int
	var d []int
	var e []int
	b = append(b, a...)
	c = append(c, a...)
	d = append(d, a...)
	e = append(e, a...)
	fmt.Println("Original slice ==>", a)
	fmt.Println("Bubble Sort    ==>", sorts.BubbleSort(a))
	fmt.Println("Selection Sort ==>", sorts.SelectionSort(b))
	fmt.Println("Insertion Sort ==>", sorts.InsertionSort(c))
	fmt.Println("Merge Sort     ==>", sorts.MergeSort(d))
	fmt.Println("Quick Sort     ==>", sorts.QuickSort(d))
}

func main() {
	var nodes []GraphNode

	for i := 0; i < 11; i++ {
		node := NewGraphNode(fmt.Sprintf("V%d", i), i-1)
		nodes = append(nodes, node)
	}

	graph := NewBFSAdjacencyMatrix(nodes)
	graph.AddUndirectedEdge(1, 2)
	graph.AddUndirectedEdge(1, 4)
	graph.AddUndirectedEdge(2, 3)
	graph.AddUndirectedEdge(2, 5)
	graph.AddUndirectedEdge(3, 6)
	graph.AddUndirectedEdge(3, 10)
	graph.AddUndirectedEdge(4, 7)
	graph.AddUndirectedEdge(5, 8)
	graph.AddUndirectedEdge(6, 9)
	graph.AddUndirectedEdge(7, 8)
	graph.AddUndirectedEdge(8, 9)
	graph.AddUndirectedEdge(9, 10)

	fmt.Println("Printing Breadth First Search ==>")
	graph.Bfs()

}
