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

func (g *graph) addEdge(n1, n2 node) {
	if g.edges == nil {
		g.edges = make(map[node][]node)
	}
	g.edges[n1] = append(g.edges[n1], n2)
}

func (n *node) string() string {
	return fmt.Sprintf("%v", n.value)
}

func (g *graph) addNode(n node) {
	g.nodes = append(g.nodes, n)
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

	fmt.Println("indegree", inDegree)
	var queue []node
	for v := range inDegree {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

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
		fmt.Println("there exist a Cycle")
		fmt.Println("topo", topSort)
	} else {
		fmt.Println("Topological sort is ", topSort)
	}

}
func main() {
	g := new(graph)
	n0 := node{0}
	n1 := node{1}
	n2 := node{2}
	n3 := node{3}
	n4 := node{4}
	n5 := node{5}
	g.addNode(n0)
	g.addNode(n1)
	g.addNode(n2)
	g.addNode(n3)
	g.addNode(n4)
	g.addNode(n5)

	g.addEdge(n5, n2)
	g.addEdge(n5, n0)
	g.addEdge(n4, n0)
	g.addEdge(n4, n1)
	g.addEdge(n2, n3)
	g.addEdge(n3, n1)
	showGraph(g)
	g.topologicalSort()
}
