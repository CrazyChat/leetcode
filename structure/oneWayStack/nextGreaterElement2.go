package oneWayStack

import (
	"math"
)

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

func NextGreaterElements(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}
	if len(nums) <= 1 {
		nums[0] = -1
		return nums
	}
	minStack := NewStack(len(nums))
	resultNums := make([]int, len(nums))
	// 走一遍，初步入栈
	for i, v := range nums {
		topIndex, topValue := minStack.Top()
		// 大于栈顶，更新resultNums，并出栈
		for v > topValue {
			resultNums[topIndex] = v
			minStack.Pop()
			topIndex, topValue = minStack.Top()
		}
		// 小于栈顶，入栈
		minStack.Push(i, v)
	}
	// 如果 栈的元素 <= 1，则表示结束了，因为有一个最大值，它肯定找不到接下来的大值
	if minStack.Len() <= 1 {
		topIndex, _ := minStack.Top()
		resultNums[topIndex] = -1
		return resultNums
	}
	// 不然还要重头开始找剩下的
	offIndex := minStack.Bottom()	// 最大值的下标
	for i := 0; i <= offIndex; i++ {
		topIndex, topValue := minStack.Top()
		for nums[i] > topValue {
			resultNums[topIndex] = nums[i]
			minStack.Pop()
			topIndex, topValue = minStack.Top()
		}
	}
	// 栈中剩下的都是找不到 下一个大值的
	for minStack.length > 0 {
		topIndex, _ := minStack.Top()
		resultNums[topIndex] = -1
		minStack.Pop()
	}
	return resultNums
}

type Stack struct {
	data []*Node
	length int
}

type Node struct {
	index int
	value int
}

func NewStack(cap int) Stack {
	return Stack{make([]*Node, cap), 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Top() (int, int) {
	if s.length == 0 {
		return 0, math.MaxInt64	// 哨兵，比任何加入的值都大，相当于空栈，任何值都直接插入
	}
	return s.data[s.length-1].index, s.data[s.length-1].value
}

func (s *Stack) Bottom() int {
	return s.data[0].index
}

func (s *Stack) Push(index, val int) {
	s.data[s.length] = &Node{index, val}
	s.length++
}

func (s *Stack) Pop() {
	s.length--
}
