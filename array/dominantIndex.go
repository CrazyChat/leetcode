package array

/*
在一个给定的数组nums中，总是存在一个最大元素 。
查找数组中的最大元素是否至少是数组中每个其他数字的两倍。
如果是，则返回最大元素的索引，否则返回-1。
 */

// 考虑极端情况，第一个值为最大值的时候如何处理
func DominantIndex(nums []int) int {
	len := len(nums)
	if len <= 1 {
		return 0
	}
	// 找出第一和第二大的数字进行比较
	var (
		maxIndex = 0
		secondIndex = 0
	)
	for i := 0; i < len; i++ {
		if nums[maxIndex] < nums[i] {
			secondIndex = maxIndex
			maxIndex = i
		} else if nums[secondIndex] < nums[i] || maxIndex == secondIndex {
			secondIndex = i
		}
	}
	if nums[maxIndex] >= 2*nums[secondIndex] {
		return maxIndex
	}
	return -1
}