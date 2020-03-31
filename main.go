package main

import (
	"github.com/crazychat/leetcode/structure"
)

func main() {
	newGraph := structure.NewGraph()
	newGraph.AddEdge(1,2)
	newGraph.AddEdge(1,3)
	newGraph.AddEdge(4,3)
	newGraph.AddEdge(2,4)
	newGraph.BFS(1)
}