package sortMethod

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	temp := mergeSort(arr)
	for i := 0; i < len(temp); i++ {
		arr[i] = temp[i]
	}
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	middle := len(arr) / 2
	left := Merge(arr[:middle])
	right := Merge(arr[middle:])
	return merge2(left, right)
}
// merge 从小到大合并两个有序数组
func merge(left, right []int) []int {
	result := make([]int, 0)
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
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}
