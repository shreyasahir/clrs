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

func (g *graph) recursiveDFS() {

	n := g.nodes[0]
	visited[n] = true
	recursiveDFS(g, n)

}

func recursiveDFS(g *graph, n node) {
	fmt.Println(n)
	//if len(vi)
	near := g.edges[n]
	for i := 0; i < len(near); i++ {
		j := near[i]
		if !visited[j] {
			recursiveDFS(g, j)
		}
	}

}

func (g *graph) dfs() {
	visited := make(map[node]bool, len(g.nodes))
	var stack []node
	n := g.nodes[0]
	stack = append(stack, n)

	for {
		if len(stack) == 0 {
			break
		}
		n = stack[len(stack)-1]
		fmt.Println("node printing", n)
		visited[n] = true
		stack = stack[:len(stack)-1]
		near := g.edges[n]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				stack = append(stack, j)
			}
		}
	}

}

func fillGraph() {
	g := graph{}

	n0 := node{1}
	n1 := node{2}
	n2 := node{3}
	n3 := node{4}
	n4 := node{5}
	n5 := node{6}
	n6 := node{7}
	g.addNode(n0)
	g.addNode(n1)
	g.addNode(n2)
	g.addNode(n3)
	g.addNode(n4)
	g.addNode(n5)
	g.addNode(n6)

	g.addEdge(n0, n1)
	g.addEdge(n0, n2)
	g.addEdge(n1, n3)
	g.addEdge(n1, n4)
	g.addEdge(n2, n5)
	g.addEdge(n2, n6)

	showGraph(&g)
	g.dfs()
	g.recursiveDFS()
}

func main() {

	fillGraph()
}
