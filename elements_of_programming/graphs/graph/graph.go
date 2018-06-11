package graph

import (
	"fmt"
)

//Vertex to hold grpah nodes
type Vertex struct {
	Value int
}

//Graph data type
type Graph struct {
	Vertices []Vertex
	Edges    map[Vertex][]Vertex
}

//AddVertex to add node to a Graph
func (g *Graph) AddVertex(n Vertex) {
	g.Vertices = append(g.Vertices, n)
}

func (v *Vertex) string() string {
	return fmt.Sprintf("%v", v.Value)
}

//AddEdges to add edge between two nodes
func (g *Graph) AddEdges(n1, n2 Vertex) {
	if g.Edges == nil {
		g.Edges = make(map[Vertex][]Vertex)
	}
	g.Edges[n1] = append(g.Edges[n1], n2)
}

//ShowGraph to print graph
func ShowGraph(g *Graph) {
	s := ""
	for i := 0; i < len(g.Vertices); i++ {
		s += g.Vertices[i].string() + " -> "
		near := g.Edges[g.Vertices[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].string() + " "
		}
		s += "\n"
	}
	fmt.Println("Graph is", s)
}
