package array

import (
	"fmt"
	"sort"
)

/*
给出一个区间的集合，请合并所有重叠的区间。
示例 1:
输入: [[1,3],[2,6],[8,10],[15,18]]
输出: [[1,6],[8,10],[15,18]]
解释: 区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
示例 2:
输入: [[1,4],[4,5]]
输出: [[1,5]]
解释: 区间 [1,4] 和 [4,5] 可被视为重叠区间。
来源：力扣（LeetCode）56
链接：https://leetcode-cn.com/problems/merge-intervals
 */

func MergeSection(intervals [][]int) [][]int {
	length := len(intervals)
	if length <= 0 {
		return nil
	}
	// 排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	fmt.Println(intervals)
	// 合并数组
	i := 1
	for i < length {
		fmt.Println("intervals: ", intervals)
		fmt.Println("比较", intervals[i-1], "&", intervals[i])
		// 比如[1,3]和[4,5]这种情况不同处理，跳过
		if intervals[i][0] > intervals[i-1][1] {
			i++
			continue
		}
		// 比如[1,3]&[1,3]或者[1,4]&[2,4]这种情况, 直接删除后面的
		if intervals[i][0] >= intervals[i-1][0] && intervals[i][1] <= intervals[i-1][1] {
			intervals = append(intervals[:i], intervals[i+1:]...)
			length--
			continue
		}
		// 只剩下[1,4]&[2,5]这种情况，合并，删除后面的
		intervals[i-1][1] = intervals[i][1]
		intervals = append(intervals[:i], intervals[i+1:]...)
		length--
	}
	return intervals
}
