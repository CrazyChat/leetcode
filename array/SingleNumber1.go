package array

// 给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素
/*
解法一：数组所有数异或一遍，结果便是不同的那个数
1. 任何数和本身异或则为0
2. 任何数和 0 异或是本身
相同的数抵消为0，0^数就是它本身
*/
func SingleNumber1(nums []int) int {
	len := len(nums)
	if len == 1 {
		return nums[0]
	}
	num := nums[0]
	for i := 1; i < len; i++ {
		num = num ^ nums[i]
	}
	return num
}
/*
解法一： 排序后左右对比
*/
// func SingleNumber(nums []int) int {
// 	len := len(nums)
// 	if len <= 1 {
// 		return nums[0]
// 	}
// 	// 排序
// 	sort.Ints(nums)
// 	// 开始查找
// 	if nums[0] != nums[1] {
// 		return nums[0]
// 	}
// 	if nums[len-1] != nums[len-2] {
// 		return nums[len-1]
// 	}
// 	for i := 1; i < len-2; i++ {
// 		if nums[i] != nums[i-1] && nums[i] != nums[i+1] {
// 			return nums[i]
// 		}
// 	}
// 	return 0
// }