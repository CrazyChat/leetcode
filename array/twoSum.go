package array

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，你不能重复利用这个数组中同样的元素。

示例:

给定 nums = [2, 7, 11, 15], target = 9

因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]

来源：力扣（LeetCode）1
链接：https://leetcode-cn.com/problems/two-sum
 */

/*
时间复杂度：
空间复杂度：
 */
func TwoSum(nums []int, target int) []int {
	length := len(nums)
	if length < 2 {
		return []int{}
	}
	numsMap := make(map[int]int)
	// 存储数组到Map, 同时每次插入查询Map中是否有对应的complement，complement = target - nums[i]
	for i := 0; i < length; i++ {
		complement := target - nums[i]
		// Map是否存在complement
		if index, ok := numsMap[complement]; ok {
			return []int{i, index}
		}
		numsMap[nums[i]] = i
	}
	return []int{}
}
