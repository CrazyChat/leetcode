package sortMethod

/*
+ 原地排序：是
+ 稳定算法：否
+ 时间复杂度：O(N^2)
 */

func SelectSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if minIndex != i {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
	}
}