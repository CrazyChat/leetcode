package array

import (
	"fmt"
	"github.com/crazychat/leetcode/sortMethod"
)

/*

 */

/*
时间复杂度：
空间复杂度：
 */
func ThreeSum(nums []int) [][]int {
	// 长度小于3不满足
	length := len(nums)
	if length < 3 {
		return [][]int{}
	}
	// 排序
	sortMethod.QuickSort(nums)
	fmt.Println("nums: ", nums)
	// i从0开始遍历到length-2，l = i+1, r = len-1
	var res [][]int
	for i := 0; i < length - 2; i++ {
		// 排序好的数组，如果 nums[0]>0 则不可能三数和=0
		if nums[0] > 0 {
			return res
		}
		// 跳过重复数，防止产生重复解
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		// 检查三数字是否满足
		left := i + 1
		right := length - 1
		for left < right {
			if nums[i] + nums[left] + nums[right] == 0 {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i] + nums[left] + nums[right] > 0 {
				// > 0 说明右边的数太大了
				right--
			} else {
				left++
			}
		}
	}
	return res
}
