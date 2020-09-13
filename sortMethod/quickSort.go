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

/*
i: 停留在大于 pivot 的索引
j: 循环整个数组，当 j<pivot 时，将i、j位置数值交换，i++
*/
func partition(arr []int, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j <= r-1; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}
