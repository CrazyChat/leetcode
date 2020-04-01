package array

/*
给定一个未经排序的整数数组，找到最长且连续的的递增序列。

示例 1:

输入: [1,3,5,4,7]
输出: 3
解释: 最长连续递增序列是 [1,3,5], 长度为3。
尽管 [1,3,5,7] 也是升序的子序列, 但它不是连续的，因为5和7在原数组里被4隔开。

来源：力扣（LeetCode）674
链接：https://leetcode-cn.com/problems/longest-continuous-increasing-subsequence
 */

func FindLengthOfLCIS(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}
	maxCount := 0
	curCount := 1
	for j := 1; j < length; j++ {
		if nums[j] > nums[j-1] {
			curCount++
		} else {
			maxCount = max(maxCount, curCount)
			curCount = 1
		}
	}
	return max(maxCount, curCount)
}
