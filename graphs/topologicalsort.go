package main

import (
	"fmt"
)

type node struct {
	value int
}

type graph struct {
	nodes []node
	edges map[node][]node
}

var (
	stack   []node
	visited = make(map[node]bool)
)

func (g *graph) addNode(n node) {
	g.nodes = append(g.nodes, n)
}

func (n *node) string() string {
	return fmt.Sprintf("%v", n.value)
}

func showGraph(g *graph) {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].string() + " -> "
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].string() + " "
		}
		s += "\n"
	}
	fmt.Println("Graph is", s)
}

func (g *graph) addEdge(n1, n2 node) {
	if g.edges == nil {
		g.edges = make(map[node][]node)
	}
	g.edges[n1] = append(g.edges[n1], n2)
}

func fillGraph() {
	g := graph{}

	n0 := node{5}
	n1 := node{7}
	n2 := node{3}
	n3 := node{11}
	n4 := node{8}
	n5 := node{2}
	n6 := node{9}
	n7 := node{10}
	g.addNode(n0)
	g.addNode(n1)
	g.addNode(n2)
	g.addNode(n3)
	g.addNode(n4)
	g.addNode(n5)
	g.addNode(n6)
	g.addNode(n7)

	g.addEdge(n0, n3)
	g.addEdge(n1, n3)
	g.addEdge(n1, n4)
	g.addEdge(n2, n4)
	g.addEdge(n2, n7)
	g.addEdge(n3, n5)
	g.addEdge(n3, n6)
	g.addEdge(n3, n7)
	g.addEdge(n4, n6)

	showGraph(&g)
	g.topologicalSort()
}

func (g *graph) topologicalSort() {
	inDegree := make(map[node]int, len(g.nodes))
	for _, v := range g.nodes {
		inDegree[v] = 0
	}

	for _, v := range g.nodes {
		near := g.edges[v]
		for i := 0; i < len(near); i++ {
			j := near[i]
			inDegree[j]++
		}
	}

	var queue []node
	for v := range inDegree {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	fmt.Println("queue with zero incoming edges", queue)

	count := 0
	var topSort []node

	for {
		if len(queue) == 0 {
			break
		}
		var u node
		u, queue = queue[0], queue[1:]
		topSort = append(topSort, u)
		near := g.edges[u]
		for i := 0; i < len(near); i++ {
			j := near[i]
			inDegree[j]--
			if inDegree[j] == 0 {
				queue = append(queue, j)
			}

		}
		count++
	}

	if count != len(g.nodes) {
		fmt.Println("There exist a cycle")
	} else {
		fmt.Println("Topological sort", topSort)
	}
}

func main() {

	fillGraph()
}
