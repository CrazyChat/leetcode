package structure

import (
	"errors"
	"fmt"
)

type Heap struct {
	cap int	// 堆可以存储的最大数据个数
	count int // 堆中已经存储的数据个数
	arr []int	// 数组，从下标1开始存储数据
}
// NewHeap 返回一个新堆
func NewHeap(cap int) Heap {
	return Heap{cap: cap, count: 0, arr: make([]int, cap+1)}
}
// Count 返回heap的个数
func (heap *Heap) Count() int {
	return len(heap.arr)
}
// Cap 返回heap的容量
func (heap *Heap) Cap() int {
	return cap(heap.arr)
}
// Insert 插入一个结点
func (heap *Heap) Insert(v int) {
	if heap.count >= heap.cap {
		return
	}
	heap.count++
	heap.arr[heap.count] = v
	// 自下往上堆化
	i := heap.count
	for i/2 > 0 && heap.arr[i] > heap.arr[i/2] {
		heap.arr[i], heap.arr[i/2] = heap.arr[i/2], heap.arr[i]
		i = i / 2
	}
}
// MoveTop 删除Top结点
func (heap *Heap) MoveTop(v int) error {
	if heap.count == 0 {
		return errors.New("没有数据")
	}
	heap.arr[1] = heap.arr[heap.count]	// 将最后一个数字放到堆顶
	// 从上到下开始堆化
	heap.count--
	heapify(heap.arr, heap.count, 1)
	return nil
}
// heapify 自上往下堆化
func heapify(arr []int, length int, startIndex int) {
	for {
		maxPos := startIndex
		// 跟其左结点比较大小
		if startIndex*2 <= length && arr[startIndex] < arr[startIndex*2] {
			maxPos = startIndex * 2
		}
		// 跟其右结点比较大小
		if startIndex*2+1 <= length && arr[startIndex] < arr[startIndex+1] {
			maxPos = startIndex * 2 + 1
		}
		// 如果比左右结点都大，则满足要求了可以退出
		if maxPos == startIndex {
			break
		}
		arr[maxPos], arr[startIndex] = arr[startIndex], arr[maxPos]
		startIndex = maxPos
	}
}
// Print 获取堆的结点数组
func (heap *Heap) Print() {
	for i := 1; i <= heap.count; i++ {
		fmt.Printf("%d ", heap.arr[i])
	}
	fmt.Println()
}
// GetArr 获取堆的结点数组
func (heap *Heap) GetArr() []int {
	return heap.arr
}

//func main() {
//	h := NewHeap(5)
//	for i := 0; i < 5; i++ {
//		h.Insert(i)
//	}
//	h.Print()
//}
