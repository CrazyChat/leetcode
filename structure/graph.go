package structure

import (
	"fmt"
)

// 无向图
type GraphNode struct {
	Value int
}

type Graph struct {
	nodes []GraphNode			// 所有节点
	edges map[GraphNode][]*GraphNode	// 邻接表表示的无向图
}

func NewGraph() *Graph {
	return &Graph{[]GraphNode{}, map[GraphNode][]*GraphNode{} }
}

func (g *Graph) AddNode(val int) {
	node := GraphNode{val}
	g.nodes = append(g.nodes, node)
}

// 增加边
func (g *Graph) AddEdge(start, end int) {
	startNode := GraphNode{start}
	endNode := GraphNode{end}
	if _, ok := g.edges[startNode]; !ok {
		g.AddNode(start)
	}
	if _, ok := g.edges[endNode]; !ok {
		g.AddNode(end)
	}
	g.edges[startNode] = append(g.edges[startNode], &endNode)
	g.edges[endNode] = append(g.edges[endNode], &startNode)
}

// 输出邻接表图
func (g *Graph) ShowGraph() {
	fmt.Println(g.nodes)
	for i := 0; i < len(g.nodes); i++ {
		node := g.nodes[i]
		fmt.Printf("%d ", node.Value)
		for _, subNode := range g.edges[node] {
			fmt.Printf("-> %d ", subNode.Value)
		}
		fmt.Printf("\n")
	}
}

func (g *Graph) String() {

}