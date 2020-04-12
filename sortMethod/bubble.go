package sortMethod

/*
+ 原地排序：是
+ 稳定算法：是
+ 时间复杂度：O(N^2)
 */
func Bubble(arr []int) {
	length := len(arr)
	if length <= 1 {
		return
	}
	for i := 0; i < length; i++ {
		// 提前退出冒泡循环的标志位
		flag := false
		for j := 0; j < length - 1 - i; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				flag = true
			}
		}
		if !flag {
			// 满足条件说明本循环一次也没发生交换，说明数组已经排列好了
			break
		}
	}
}
