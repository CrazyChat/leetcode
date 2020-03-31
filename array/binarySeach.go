package array

func BinarySeach(nums []int, target int) int {
	length := len(nums)
	left := 0
	right := length - 1
	for left < right {
		mid := (right + left) / 2
		if target == nums[mid] {
			for mid < right && nums[mid+1] == nums[mid] {
				mid++
			}
			return mid
		}
		if target >= nums[mid] {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return -1
}
