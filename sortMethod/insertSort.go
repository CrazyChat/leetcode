package sortMethod

/*
+ 原地排序：是
+ 稳定算法：是
+ 时间复杂度：O(N^2)
 */

func InsertSort(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 1; i < length; i++ {
		j := i - 1
		value := arr[i]
		for ; j >= 0; j-- {
			if arr[j] > value {
				arr[j+1] = arr[j]
			} else {
				break
			}
		}
		arr[j+1] = value
	}
}