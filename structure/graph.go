package structure

import (
	"fmt"
	"sync"
)

type GraphNode struct {
	Value int
}

type Graph struct {
	nodes []*GraphNode			// 所有节点
	edges map[GraphNode][]*GraphNode	// 邻接表表示的无向图
	lock sync.RWMutex				// 保证线程安全
}

func (g *Graph) AddNode(n *GraphNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	g.nodes = append(g.nodes, n)
}

// 增加边
func (g *Graph) AddEdge(start, end *GraphNode) {
	g.lock.Lock()
	defer g.lock.Unlock()
	// 首次建立图
	if g.edges == nil {
		g.edges = make(map[GraphNode][]*GraphNode)
	}
	g.edges[*start] = append(g.edges[*start], end)
	g.edges[*end] = append(g.edges[*end], start)
}

// 输出邻接表图
func (g *GraphNode) String() string {
	return fmt.Sprintf("%v", g.Value)
}

func (g *Graph) String() {
	g.lock.RLock()
	defer g.lock.RUnlock()
	str := ""
	for _, node := range g.nodes {
		str += node.String() + " -> "
		nexts := g.edges[*node]
		for _, next := range nexts {
			str += next.String() + " "
		}
		str += "\n"
	}
	fmt.Println(str)
}