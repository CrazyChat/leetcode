package array

/*
LeetCode 215
在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素
说明: 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度
[3,2,1,5,6,4], k = 2 => 5
 */
/*
利用快速排序的方法查找第K大, 只处理包含k大值的那一部分
 */
func FindKthLargest(nums []int, k int) int {
	quickSort(nums, k-1)
	return nums[k-1]
}

func quickSort(arr []int, k int) {
	quick_sort(arr, 0, len(arr)-1, k)
}

func quick_sort(arr []int, p, r, k int) {
	if p >= r {
		return
	}
	q := partition(arr, p, r) // 获取分区点
	if q == k {
		return
	} else if  k < q {
		quick_sort(arr, p, q-1, k)
	} else {
		quick_sort(arr, q+1, r, k)
	}
}

func partition(arr []int, p, r int) int {
	pivot := arr[r]
	i := p
	for j := p; j <= r-1; j++ {
		// < pivot移动i，否则不移动
		if arr[j] > pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i+=1
		}
	}
	arr[i], arr[r] = arr[r], arr[i]
	return i
}