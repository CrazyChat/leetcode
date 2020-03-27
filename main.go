package main

import "github.com/crazychat/leetcode/structure"

func main() {
	g := structure.Graph{}
	n1, n2, n3, n4, n5 := structure.GraphNode{1}, structure.GraphNode{2}, structure.GraphNode{3}, structure.GraphNode{4}, structure.GraphNode{5}

	g.AddNode(&n1)
	g.AddNode(&n2)
	g.AddNode(&n3)
	g.AddNode(&n4)
	g.AddNode(&n5)

	g.AddEdge(&n1, &n2)
	g.AddEdge(&n1, &n5)
	g.AddEdge(&n2, &n3)
	g.AddEdge(&n2, &n4)
	g.AddEdge(&n2, &n5)
	g.AddEdge(&n3, &n4)
	g.AddEdge(&n4, &n5)

	g.String()
}


