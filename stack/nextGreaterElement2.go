package stack

/*
503. 下一个更大元素 II
给定一个循环数组（最后一个元素的下一个元素是数组的第一个元素），输出每个元素的下一个更大元素。数字 x 的下一个更大的元素是按数组遍历顺序，这个数字之后的第一个比它更大的数，这意味着你应该循环地搜索它的下一个更大的数。如果不存在，则输出 -1。

示例 1:

输入: [1,2,1]
输出: [2,-1,2]
解释: 第一个 1 的下一个更大的数是 2；
数字 2 找不到下一个更大的数；
第二个 1 的下一个最大的数需要循环搜索，结果也是 2。
注意: 输入数组的长度不会超过 10000。
*/

// 单调递减栈
func NextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) <= 1 {
		nums[0] = -1
		return nums
	}
	resultNums := make([]int, len(nums))
	stack := NewOneWayStack()
	stack.Push(0)
	// 先走一遍，初步存储数据
	for i := 1; i < len(nums); i++ {
		for stack.Len() != 0 && nums[i] > nums[stack.Top()] {
			resultNums[stack.Pop()] = nums[i]
		}
		stack.Push(i)
	}
	// 如果最大值和最小值相等，则表示操作完成了，因为数组的最大值 必定找不到 下一个大值
	offIndex := stack.Bottom()
	if nums[stack.Top()] == nums[offIndex] {
		for stack.Len() > 0 {
			resultNums[stack.Pop()] = -1
		}
		return resultNums
	}
	// 否则再走一遍，从头开始找比 栈中值 大的值
	i := 0
	for nums[stack.Top()] != nums[stack.Bottom()] && i <= offIndex {
		if nums[i] > nums[stack.Top()] {
			resultNums[stack.Pop()] = nums[i]
		} else {
			i++
		}
	}
	for stack.Len() > 0 {
		resultNums[stack.Pop()] = -1
	}
	return resultNums
}
