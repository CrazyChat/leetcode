package array

/*
给定一个未排序的整数数组，找出最长连续序列的长度。

要求算法的时间复杂度为 O(n)。

示例:

输入: [100, 4, 200, 1, 3, 2]
输出: 4
解释: 最长连续序列是 [1, 2, 3, 4]。它的长度为 4。

来源：力扣（LeetCode）128
链接：https://leetcode-cn.com/problems/longest-consecutive-sequence
 */

func LongestConsecutive(nums []int) int {
	length := len(nums)
	if length < 1 {
		return 0
	}
	if length == 1 {
		return 1
	}
	numsMap := make(map[int]bool)
	// 存储所有数字到map
	for i := 0; i < length; i++ {
		numsMap[nums[i]] = true
	}
	maxLength := 1
	for val, _ := range numsMap {
		// 只对 当前数字-1 不在哈希表里的数字作为连续序列的第一个数字去找对应的最长序列，这是因为其他数字一定已经出现在了某个序列里。
		if !numsMap[val-1] {
			step := 1
			curLength := 1
			if numsMap[val+1] {
				for numsMap[val+step] {
					curLength++
					step++
				}
			}
			maxLength = max(maxLength, curLength)
		}
	}
	return maxLength
}