package main

import (
	"fmt"
)

type node struct {
	value int
}
type graph struct {
	edges map[node][]node
	nodes []node
}

func (g *graph) addNode(n node) {
	g.nodes = append(g.nodes, n)
}

func (n *node) string() string {
	return fmt.Sprintf("%v", n.value)
}

func (g *graph) addEdge(n1, n2 node) {
	if g.edges == nil {
		g.edges = make(map[node][]node)
	}
	g.edges[n1] = append(g.edges[n1], n2)
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

func (g *graph) isCycleUtil(color map[node]string, i node) bool {
	color[i] = "grey"
	for _, v := range g.edges[i] {
		if color[v] == "grey" {
			return true
		}
		if color[v] == "white" && g.isCycleUtil(color, v) == true {
			return true
		}
	}
	color[i] = "black"
	return false
}

func (g *graph) isCycle() bool {
	color := make(map[node]string, len(g.nodes))
	for _, v := range g.nodes {
		color[v] = "white"
	}

	for _, v := range g.nodes {
		if color[v] == "white" {
			if g.isCycleUtil(color, v) == true {
				return true
			}
		}
	}
	return false
}

func main() {
	g := new(graph)
	n1 := node{1}
	n2 := node{2}
	n3 := node{3}
	n4 := node{4}
	n5 := node{5}
	g.addNode(n1)
	g.addNode(n2)
	g.addNode(n3)
	g.addNode(n4)
	g.addNode(n5)

	g.addEdge(n1, n2)
	g.addEdge(n1, n3)
	g.addEdge(n3, n5)
	g.addEdge(n2, n5)
	g.addEdge(n2, n4)
	g.addEdge(n2, n1)
	showGraph(g)
	fmt.Println("is cycle", g.isCycle())
}
