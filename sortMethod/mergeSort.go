package sortMethod

func MergeSort(arr []int) {
	if len(arr) <= 1 {
		return
	}
	mergeSort(arr, 0, len(arr)-1)
}

func mergeSort(arr []int, p, r int) {
	// 递归终止条件
	if p >= r {
		return
	}
	middle := (p + r) / 2
	// 不断地进行左右对半划分
	mergeSort(arr, p, middle)
	mergeSort(arr, middle+1, r)
	// 合并
	merge(arr, p, middle, r)
}
// merge 从小到大合并两个有序数组
func merge(arr []int, p, middle, r int) {
	temp := make([]int, 0)
	i, j:= p, middle+1
	for i <= middle && j <= r {
		if arr[i] > arr[j] {
			temp = append(temp, arr[j])
			j++
		} else {
				temp = append(temp, arr[i])
				i++
		}
	}
	// 把左右边剩余的数移入数组
	if i > middle {
		temp = append(temp, arr[j:r+1]...)
	} else {
		temp = append(temp, arr[i:middle+1]...)
	}
	// 将temp修改到arr中
	for i, v := range temp {
		arr[p + i] = v
	}
}
