package sortMethod

func MergeSort(arr []int) []int {
	len := len(arr)
	if len <= 1 {
		return arr
	}
	middle := len / 2
	// 不断地进行左右对半划分
	left := MergeSort(arr[:middle])
	right := MergeSort(arr[middle:])
	// 合并
	return merge(left, right)
}
// 从小到大排序.
func merge(left, right []int) []int {
	var result []int
	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] > right[r] {
			result = append(result, right[r])
			r++
		} else {
			result = append(result, left[l])
			l++
		}
	}
	// 把左右边剩余的数移入数组
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
