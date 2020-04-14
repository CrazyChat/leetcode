package structure

import "container/heap"

/*
优先级队列
 */

type Item struct {
	Value string	// 优先级队列中的数据，可以是任意类型，这里使用string
	Priority int	// 优先级队列中节点的优先级
	Index int	// index是该节点在堆中的位置
}

// 优先级队列需要实现heap的interface
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].Index = pq[j].Index
}
// 将index置为-1是为了标识该数据已经出了优先级队列了
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	item := old[len(old)-1]
	item.Index = -1
	*pq = old[:len(old)-1]
	return item
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	item.Index = len(*pq)
	*pq = append(*pq, item)
}
// 更新修改了优先级和值的item在优先级队列中的位置
func (pq *PriorityQueue) Update(item *Item, value string, priority int) {
	item.Value = value
	item.Priority = priority
	heap.Fix(pq, item.Index)
}
