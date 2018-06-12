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

// func union(parent map[node]int, x, y node) {
// 	xSet := findParent(parent, x)
// 	ySet := findParent(parent, y)
// 	parent[xSet] = ySet
// }
func (g *graph) isCycle() bool {
	var parent = make(map[node]int, len(g.nodes))
	for _, v := range g.nodes {
		parent[v] = -1
	}

	for _, v := range g.nodes {
		for _, v1 := range g.edges[v] {
			x := findParent(parent, v)
			y := findParent(parent, v1)
			if x == y {
				return true
			}

			fmt.Println("parent", parent)
			//union(parent, x, y)
		}
	}
	return false
}

func findParent(parent map[node]int, i node) node {
	if parent[i] == -1 {
		return i
	}
	return findParent(parent, i)

}

func fillGraph() {
	g := graph{}

	n0 := node{1}
	n1 := node{2}
	n2 := node{3}

	g.addNode(n0)
	g.addNode(n1)
	g.addNode(n2)

	g.addEdge(n0, n1)
	g.addEdge(n1, n2)
	g.addEdge(n1, n0)
	g.addEdge(n2, n1)
	g.addEdge(n0, n2)

	g.addEdge(n2, n0)
	showGraph(&g)
	g.isCycle()
}

func main() {

	fillGraph()
}
