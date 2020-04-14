package main

import (
	"container/heap"
	"fmt"
	"github.com/crazychat/leetcode/structure"
)

func main() {
	// 创建节点并设计他们的优先级
	items := map[string]int{"二毛": 5, "张三": 3, "狗蛋": 9}
	i := 0
	pq := make(structure.PriorityQueue, len(items)) // 创建优先级队列，并初始化
	for k, v := range items {             // 将节点放到优先级队列中
		pq[i] = &structure.Item{
			Value:    k,
			Priority: v,
			Index:    i}
		i++
	}
	heap.Init(&pq) // 初始化堆
	item := &structure.Item{ // 创建一个item
		Value:    "李四",
		Priority: 1,
	}
	heap.Push(&pq, item)           // 入优先级队列
	pq.Update(item, item.Value, 2) // 更新item的优先级
	for len(pq) > 0 {
		item := heap.Pop(&pq).(*structure.Item)
		fmt.Printf("%.2d:%s index:%.2d\n", item.Priority, item.Value, item.Index)
	}
}