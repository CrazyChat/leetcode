package stack

import "github.com/crazychat/leetcode/structure/oneWayStack"

/*
84. 柱状图中最大的矩形
给定 n 个非负整数，用来表示柱状图中各个柱子的高度。每个柱子彼此相邻，且宽度为 1 。

求在该柱状图中，能够勾勒出来的矩形的最大面积。
以上是柱状图的示例，其中每个柱子的宽度为 1，给定的高度为 [2,1,5,6,2,3]。
图中阴影部分为所能勾勒出的最大矩形面积，其面积为 10 个单位。

示例:
输入: [2,1,5,6,2,3]
输出: 10
*/

/*
栈 存储的是 索引
高度是当前heights[i]，问题的关键是求宽度
https://blog.csdn.net/Zolewit/article/details/88863970
*/

func LargestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}
	// 头部加0，充当哨兵，栈空的时候直接进栈
	// 尾部加0，最后0入栈可以清空栈
	zero := []int{0}
	heights = append(zero, heights...)
	heights = append(heights, 0)
	stack := NewOneWayStack()
	stack.Push(0)
	maxArea := 0
	for i, v := range heights {
		for v < heights[stack.Top()] {
			index := stack.Pop()
			tempArea := (i - stack.Top() - 1) * heights[index]
			maxArea = oneWayStack.max(maxArea, tempArea)
		}
		stack.Push(i)
	}
	return maxArea
}
