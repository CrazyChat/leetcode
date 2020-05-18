package mix

/*
33. 搜索旋转排序数组
*/

/*
我们发现循环数组存在一个性质：以数组中间点为分区，会将数组分成一个有序数组和一个循环有序数组。

 如果首元素小于 mid，说明前半部分是有序的，后半部分是循环有序数组；
 如果首元素大于 mid，说明后半部分是有序的，前半部分是循环有序的数组；
 如果目标元素在有序数组范围中，使用二分查找；
 如果目标元素在循环有序数组中，设定数组边界后，使用以上方法继续查找。
*/

func Search(nums []int, target int) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 前半部分是有序的，则后部分是循环部分
			// target 在前半部分有序数组里
			if target >= nums[0] && target < nums[mid] {
				r = mid - 1
			} else {
				// target在循环部分里
				l = mid + 1
			}
		} else { // 后半部分是有序的，则前部分是循环部分
			// target 在有序数组里
			if target > nums[mid] && target <= nums[len(nums)-1] {
				l = mid + 1
			} else {
				// target在循环部分里
				r = mid - 1
			}
		}
	}
	return -1
}
