package array

/*
示例:

输入: [0,1,0,2,1,0,1,3,2,1,2,1]
输出: 6
LeetCode 42
 */

// 按列求
func Trap(height []int) int {
	length := len(height)
	sum := 0
	for i := 0; i < length; i++ {
		var maxLeft, maxRight int
		// 1. 找出左边最高
		for j := 0; j < i; j++ {
			if height[j] > maxLeft {
				maxLeft = height[j]
			}
		}
		// 2. 找出右边最高
		for j := i+1; j < length; j++ {
			if height[j] > maxRight {
				maxRight = height[j]
			}
		}
		// 3. 找出两端交矮的
		min := trapMin(maxLeft, maxRight)
		// 4. 只有较小的一端大于当前列的高度才会有水，其他情况不会有水
		if min > height[i] {
			sum = sum + min - height[i]
		}
	}
	return sum
}

func trapMin(i, j int) int {
	if i < j {
		return i
	}
	return j
}