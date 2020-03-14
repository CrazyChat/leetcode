package array

/*
数组中有一个数字出现的次数超过数组长度的一半，请找出这个数字。你可以假设数组是非空的，并且给定的数组总是存在多数元素。
限制：1 <= 数组长度 <= 50000
[1, 2, 3, 2, 2, 2, 5, 4, 2] => 2
 */

/*
1. 哈希表统计法： 遍历数组 nums ，用 HashMap 统计各数字的数量，最终超过数组长度一半的数字则为众数。
	此方法时间、空间复杂度均为 O(N)O(N)
2. 数组排序法： 将数组 nums 排序，由于众数的数量超过数组长度一半，因此 数组中点的元素 一定为众数。
	此方法时间复杂度 O(N log_2 N)O(NlogN)
3. 摩尔投票法： 核心理念为 “正负抵消” 。时间复杂度为 O(N)O(N) ，空间复杂度为 O(1)O(1) 。是本题的最佳解法。
 */

func MajorityElement(nums []int) int {
	length := len(nums)
	votes := 0
	var mode int
	for i := 0; i < length; i++ {
		if votes == 0 {
			mode = nums[i]
		}
		if nums[i] == mode {
			votes++
			if votes > length/2 {
				return mode
			}
		} else {
			votes--
		}
	}
	return mode
}
