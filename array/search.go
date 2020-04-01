package array

/*
假设按照升序排序的数组在预先未知的某个点上进行了旋转。
( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。
搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。
你可以假设数组中不存在重复的元素。
你的算法时间复杂度必须是 O(log n) 级别。
示例 1:
输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
来源：力扣（LeetCode）33
链接：https://leetcode-cn.com/problems/search-in-rotated-sorted-array
 */
func Search(nums []int, target int) int {
	length := len(nums)
	if length == 0 {
		return -1
	}
	if target >= nums[0] {
		// 从左边开始找
		for i := 0; i < length; i++ {
			if target == nums[i] {
				return i
			}
			if i >= length-1 || nums[i+1] < nums[i] {
				return -1
			}
		}
	} else {
		// 从右边开始找
		for i := length - 1; i >= 0; i-- {
			if target == nums[i] {
				return i
			}
			if i <= 0 || nums[i-1] > nums[i] {
				return -1
			}
		}
	}
	return -1
}