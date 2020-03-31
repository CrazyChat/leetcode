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
// AddNode 增加结点
func (g *Graph) AddNode(val int) {
	node := GraphNode{val}
	g.nodes = append(g.nodes, node)
}
// AddEdge 增加边
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
// ShowGraph 输出邻接表图结构
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
// BFS 广度搜索，开始 -> 结束
func (g *Graph) BFSLayer(start, end int) {
	visited := map[GraphNode]bool{} // 以及访问的结点
	waitVisit := []*GraphNode{} // 等待访问的结点
	startNode := GraphNode{start}
	fmt.Printf("%d\n", start)
	visited[startNode] = true
	// 输出第一层关系
	for _, node := range g.edges[startNode] {
		fmt.Printf("%d ", node.Value)
		if node.Value == end {
			return
		}
		visited[*node] = true
		waitVisit = append(waitVisit, node)
	}
	fmt.Printf("\n")
	// 循环waitVisit 直到没有待访问的Node
	for {
		tempWaitVisit := []*GraphNode{}
		for _, val := range waitVisit {
			for _, newNode := range g.edges[*val] {
				if _, ok := visited[*newNode]; ok {
					continue
				}
				fmt.Printf("%d ", newNode.Value)
				if newNode.Value == end {
					return
				}
				visited[*newNode] = true
				if len(waitVisit) == 1 {
					return
				}
				tempWaitVisit = append(tempWaitVisit, newNode)
			}
		}
		fmt.Printf("\n")
		waitVisit = tempWaitVisit
	}
}
// BFS 广度优先遍历
func (g *Graph) BFS(start int) {
	visited := map[GraphNode]bool{} // 以及访问的结点
	waitVisit := []*GraphNode{} // 等待访问的结点
	startNode := GraphNode{start}
	fmt.Printf("%d ", start)
	visited[startNode] = true
	waitVisit = append(waitVisit, &startNode)
	for {
		node := waitVisit[0]
		for _, subNode := range g.edges[*node] {
			if _, ok := visited[*subNode]; ok {
				continue
			}
			fmt.Printf("%d ", subNode.Value)
			visited[*subNode] = true
			waitVisit = append(waitVisit, subNode)
		}
		if len(waitVisit) == 1 {
			return
		}
		waitVisit = waitVisit[1:]
	}
}
// DFS 深度搜索
func (g *Graph) DFS(start, end int) {

}