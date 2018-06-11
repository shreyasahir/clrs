package main

import (
	"clrs/elements_of_programming/graphs/graph"
	"fmt"
)

func topologicalSort(g *graph.Graph) {
	var inDegree = make(map[graph.Vertex]int, len(g.Vertices))
	for _, v := range g.Vertices {
		inDegree[v] = 0
	}

	for _, v := range g.Vertices {
		near := g.Edges[v]
		for i := 0; i < len(near); i++ {
			j := near[i]
			inDegree[j]++
		}
	}

	var queue []graph.Vertex
	for v := range inDegree {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	count := 0
	var topSort []graph.Vertex

	for {
		if len(queue) == 0 {
			break
		}
		var u graph.Vertex
		u, queue = queue[0], queue[1:]
		topSort = append(topSort, u)
		near := g.Edges[u]
		for i := 0; i < len(near); i++ {
			j := near[i]
			inDegree[j]--
			if inDegree[j] == 0 {
				queue = append(queue, j)
			}
		}
		count++
	}

	if count != len(g.Vertices) {
		fmt.Println("There exist a cycle")
	} else {
		for _, v := range topSort {
			fmt.Println(v)
		}
	}
}

func bfs(g *graph.Graph) {
	visited := make(map[graph.Vertex]bool)
	var queue []graph.Vertex
	n := g.Vertices[0]
	queue = append(queue, n)
	for {
		if len(queue) == 0 {
			break
		}
		node := queue[0]
		queue = queue[1:]
		visited[node] = true
		near := g.Edges[node]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				queue = append(queue, j)
				visited[j] = true
			}
		}

		fmt.Printf("%v", node)
		fmt.Println()

	}
}

func dfs(g *graph.Graph) {
	visited := make(map[graph.Vertex]bool)
	var stack []graph.Vertex
	n := g.Vertices[0]
	stack = append(stack, n)
	for {
		if len(stack) == 0 {
			break
		}
		node := stack[len(stack)-1]
		fmt.Printf("%v", node)
		fmt.Println()
		stack = stack[:len(stack)-1]
		visited[node] = true
		near := g.Edges[node]
		for i := 0; i < len(near); i++ {
			j := near[i]
			if !visited[j] {
				stack = append(stack, j)
				visited[j] = true
			}
		}
	}
}

func findParent(parent map[graph.Vertex]int, i graph.Vertex) graph.Vertex {
	if parent[i] == -1 {
		return i
	}
	return findParent(parent, i)
}

func isCycle(g *graph.Graph) bool {
	var parent = make(map[graph.Vertex]int, len(g.Vertices))
	for _, v := range g.Vertices {
		parent[v] = -1
	}

	for _, v := range g.Vertices {
		for _, v1 := range g.Edges[v] {
			x := findParent(parent, v)
			y := findParent(parent, v1)

			if x == y {
				return true
			}
		}
	}

	return false
}
func main() {
	g := graph.Graph{}
	v1 := graph.Vertex{Value: 1}
	v2 := graph.Vertex{Value: 2}
	v3 := graph.Vertex{Value: 3}
	v4 := graph.Vertex{Value: 4}
	v5 := graph.Vertex{Value: 5}
	v6 := graph.Vertex{Value: 6}
	v7 := graph.Vertex{Value: 7}
	v8 := graph.Vertex{Value: 8}
	g.AddVertex(v1)
	g.AddVertex(v2)
	g.AddVertex(v3)
	g.AddVertex(v4)
	g.AddVertex(v5)
	g.AddVertex(v6)
	g.AddVertex(v7)
	g.AddVertex(v8)

	g.AddEdges(v1, v2)
	g.AddEdges(v1, v3)
	g.AddEdges(v2, v5)
	g.AddEdges(v3, v5)
	g.AddEdges(v3, v8)
	g.AddEdges(v5, v6)
	g.AddEdges(v2, v4)
	g.AddEdges(v4, v6)
	g.AddEdges(v6, v5)
	g.AddEdges(v8, v7)
	//g.AddEdges(v8, v8)

	graph.ShowGraph(&g)
	bfs(&g)
	fmt.Println("DFS")
	dfs(&g)
	fmt.Println("isCycle", isCycle(&g))
	fmt.Println("topological sort")
	topologicalSort(&g)
}
