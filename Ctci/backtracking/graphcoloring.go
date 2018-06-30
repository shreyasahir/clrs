package main

import (
	"fmt"
)

type vertex struct {
	value int
	color int
}
type graph struct {
	vertices []vertex
	edges    map[vertex][]vertex
}

func (g *graph) addEdges(v1, v2 vertex) {
	if g.edges == nil {
		g.edges = make(map[vertex][]vertex)
	}
	g.edges[v1] = append(g.edges[v1], v2)
}

func (g *graph) addVertex(v vertex) {
	g.vertices = append(g.vertices, v)
}

func showGraph(g *graph) {
	s := ""
	for i := 0; i < len(g.vertices); i++ {
		s += g.vertices[i].string() + " -> "
		near := g.edges[g.vertices[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].string() + " "
		}
		s += "\n"
	}
	fmt.Println("Graph is", s)
}
func (n *vertex) string() string {
	return fmt.Sprintf("%v", n.value)
}

func (n *vertex) stringColor() string {
	return fmt.Sprintf("%v", n.value)
}

func graphColoring(g *graph, m int) bool {
	if graphColoringUtil(g, m, 0) == false {
		fmt.Println("yes")
		return false
	}
	fmt.Println("color is")
	g.showColor()
	return true
}

func (g *graph) showColor() {
	//s := ""
	for i := 0; i < len(g.vertices); i++ {
		fmt.Println("color is", g.vertices[i].color)
	}
	//fmt.Println("color is", s)
}

func graphColoringUtil(g *graph, m int, v int) bool {

	if v == len(g.vertices) {
		return true
	}

	for i := 1; i <= m; i++ {
		if issafe(g, v, i) == true {
			g.vertices[v].color = i
			if graphColoringUtil(g, m, v+1) == true {
				return true
			}
			g.vertices[v].color = 0

		}
	}
	return false
}

func issafe(g *graph, v int, c int) bool {
	vertex := g.vertices[v]

	for _, v := range g.edges[vertex] {
		fmt.Println("vertex", v)
		fmt.Println("color", c)
		fmt.Println("color", v.color)
		if v.color == c {
			return false
		}
	}
	return true
}
func main() {
	g := &graph{}
	v0 := vertex{value: 0}
	g.addVertex(v0)
	v1 := vertex{value: 1}
	g.addVertex(v1)
	v2 := vertex{value: 2}
	g.addVertex(v2)
	v3 := vertex{value: 3}
	g.addVertex(v3)
	g.addEdges(v0, v1)
	g.addEdges(v0, v2)
	g.addEdges(v0, v3)
	g.addEdges(v1, v2)
	g.addEdges(v1, v0)
	g.addEdges(v2, v0)
	g.addEdges(v2, v1)
	g.addEdges(v2, v3)
	g.addEdges(v3, v0)
	g.addEdges(v3, v2)
	showGraph(g)
	graphColoring(g, 3)

}
