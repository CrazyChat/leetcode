package array

/*
给定一个整数类型的数组 nums，请编写一个能够返回数组“中心索引”的方法。
我们是这样定义数组中心索引的：数组中心索引的左侧所有元素相加的和等于右侧所有元素相加的和。
如果数组不存在中心索引，那么我们应该返回 -1。如果数组有多个中心索引，那么我们应该返回最靠近左边的那一个。
 */

import "fmt"

func FindCentralIndex(arr []int) int {
	len := len(arr)
	for center := 0; center < len; center++ {
		var leftSum int = 0
		var rightSum int = 0
		for left := 0; left < center; left++ {
			leftSum += arr[left]
		}
		for right := center+1; right < len; right++ {
			rightSum += arr[right]
		}
		fmt.Println("center:", center, ", left:", leftSum, ", right:", rightSum)
		if leftSum == rightSum {
			return center
		}
	}
	return -1
}
