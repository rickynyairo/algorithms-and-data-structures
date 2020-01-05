package main

import "fmt"
import "sorts"
import . "graphs"
import . "greedy"
import "dividenconc"

func main() {
	var coins = []int{1, 2, 5, 10, 20, 50, 100, 500, 2000}
	CoinChange(coins, 6799)
	var houses = []int{6, 7, 1, 30, 8, 2, 4}
	fmt.Println("10th fibonacci element ==> ", dividenconc.Fibonacci(10))
	fmt.Println("Number Factor for 6 ==> ", dividenconc.NumberFactor(6))
	fmt.Println("House Thief algo ==> ", dividenconc.HouseThief(houses, 0))
}

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

func mainss() {
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
