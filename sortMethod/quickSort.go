package sortMethod

func QuickSort(arr []int) {
	quick_sort(arr, 0, len(arr)-1)
}

func quick_sort(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r) // 获取分区点
	quick_sort(arr, p, q-1)
	quick_sort(arr, q+1, r)
}

func partition(arr []int, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j <= r-1; j++ {
		// < pivot移动i，否则不移动
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i+=1
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
