package sortMethod

/*
+ 原地排序：是
+ 稳定算法：是
+ 时间复杂度：O(N^2)
 */
func Bubble(arr []int) {
	if len(arr) <= 1 {
		return
	}
	for i := 0; i < len(arr); i++ {
		flag := false
		for j := 1; j < len(arr) - i; j++ {
			if arr[j] < arr[j-1] {
				flag = true
				arr[j], arr[j-1] = arr[j-1], arr[j]
			}
		}
		if !flag {
			return
		}
	}
}
