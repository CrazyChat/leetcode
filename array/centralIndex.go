package array

/*
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
 */

func PivotIndex(nums []int) int {
	len := len(nums)
	if len == 1 {
		return 0
	}
	sumTotal := 0
	sumLeft := 0
	for i := 0; i < len; i++ {
		sumTotal += nums[i]
	}
	for mid := 0; mid < len; mid++ {
		if sumLeft * 2 == sumTotal - nums[mid] {
			return mid
		}
		sumLeft += nums[mid]
	}
	return -1
}
