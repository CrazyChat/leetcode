package array

import (
	"fmt"
	"math"
)

/*
496. 下一个更大元素 I
给定两个没有重复元素的数组 nums1 和 nums2 ，其中nums1 是 nums2 的子集。找到 nums1 中每个元素在 nums2 中的下一个比其大的值。

nums1 中数字 x 的下一个更大元素是指 x 在 nums2 中对应位置的右边的第一个比 x 大的元素。如果不存在，对应位置输出-1。

示例 1:

输入: nums1 = [4,1,2], nums2 = [1,3,4,2].
输出: [-1,3,-1]
解释:
    对于num1中的数字4，你无法在第二个数组中找到下一个更大的数字，因此输出 -1。
    对于num1中的数字1，第二个数组中数字1右边的下一个较大数字是 3。
    对于num1中的数字2，第二个数组中没有下一个更大的数字，因此输出 -1。
 */

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	if len(nums2) <= 1 {
		for i := 0; i < len(nums1); i++ {
			nums1[i] = -1
		}
		return nums1
	}
	dataMap := map[int]int{}
	minStack := NewStack(len(nums2))
	for i := 0; i < len(nums2); i++ {
		// 大于栈顶，更新进dataMap，并出栈
		top := minStack.Top()
		for nums2[i] > top {
			dataMap[top] = nums2[i]
			minStack.Pop()
			top = minStack.Top()
		}
		minStack.Push(nums2[i])
	}
	fmt.Println(dataMap)
	// 遍历nums1赋值
	for i := 0; i < len(nums1); i++ {
		if val, ok := dataMap[nums1[i]]; ok {
			nums1[i] = val
		} else {
			nums1[i] = -1
		}
	}
	return nums1
}

type Stack struct {
	data []int
	length int
}

func NewStack(cap int) Stack {
	return Stack{make([]int, cap), 0}
}

func (s *Stack) Top() int {
	if s.length == 0 {
		return math.MaxInt64	// 哨兵
	}
	return s.data[s.length-1]
}

func (s *Stack) Push(val int) {
	s.data[s.length] = val
	s.length++
}

func (s *Stack) Pop() {
	s.length--
}
