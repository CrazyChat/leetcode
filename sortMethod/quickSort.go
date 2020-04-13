package sortMethod

func QuickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	quickSort(arr, 0, len(arr)-1)
}

func quickSort(arr []int, p, r int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r) // 获取分区点
	quickSort(arr, p, q-1)
	quickSort(arr, q+1, r)
}

func partition(arr []int, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j <= r-1; j++ {
		// < pivot移动i，否则不移动i => i停留的位置一直是大于pivot的数字
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
