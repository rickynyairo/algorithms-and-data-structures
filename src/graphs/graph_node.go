package graphs

import "fmt"

type GraphNode struct {
	name       string
	index      int
	distance   int
	neighbours []*GraphNode
	is_visited bool
	parent     *GraphNode
	weight_map map[*GraphNode]int
	value      int
	next       *GraphNode
}

func (gn *GraphNode) String() string {
	return fmt.Sprint(gn.name)
}

func NewGraphNode(name string, index int) GraphNode {
	node := GraphNode{
		name:       name,
		index:      index,
		distance:   1,
		is_visited: false,
		parent:     nil,
	}
	return node
}

type LinkedList struct {
	first  *GraphNode
	last   *GraphNode
	length int
}

func (l *LinkedList) Append(n *GraphNode) {
	if l.length == 0 {
		l.first = n
		l.last = n
		l.length++
		return
	}

	l.last.next = n
	l.last = n
	l.length++
}

func (l *LinkedList) Prepend(n *GraphNode) {
	if l.length == 0 {
		l.first = n
		l.last = n
		l.length++
		return
	}

	firstNode := l.first
	l.first = n
	l.first.next = firstNode
	l.length++
}

func (l *LinkedList) Lookup() []int {
	if l.length == 0 {
		return nil
	}

	var nodes []int
	node := l.first
	for node != nil {
		nodes = append(nodes, node.value)
		node = node.next
	}

	return nodes
}

func (l *LinkedList) Delete(v int) {
	if l.length == 0 {
		return
	}

	if l.first.value == v {
		l.first = l.first.next
		l.length--
		return
	}

	var prevNode *GraphNode
	node := l.first

	for node.value != v {
		if node == nil {
			return
		}

		prevNode = node
		node = node.next
	}

	prevNode.next = node.next
	l.length--
}
