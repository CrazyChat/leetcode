package main

import (
	"github.com/crazychat/leetcode/structure"
)

func main() {
	newGraph := structure.NewGraph()
	newGraph.AddEdge(1,2)
	newGraph.AddEdge(1,3)
	newGraph.AddEdge(2,4)
	newGraph.AddEdge(2,5)
	newGraph.AddEdge(4,6)
	newGraph.AddEdge(4,7)
	newGraph.AddEdge(6,8)
	newGraph.AddEdge(7,9)
	newGraph.AddEdge(8,10)
	newGraph.AddEdge(10,11)
	//newGraph.ShowGraph()
	newGraph.BFS(1)
}