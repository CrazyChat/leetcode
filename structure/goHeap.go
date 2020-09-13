package structure

/*
	arr := []int{2, 1, 5, 6, 4, 7, 9, 8, 0}
	h := IntHeap{}
	for _, val := range arr {
		heap.Push(&h, val)
	}
	fmt.Println(h)
 */

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	item := old[len(old)-1]
	*h = old[:len(old)-1]
	return item
}

func (h *IntHeap) Push(value interface{}) {
	*h = append(*h, value.(int))
}
