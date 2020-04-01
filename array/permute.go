package array

/*
给定一个 没有重复 数字的序列，返回其所有可能的全排列。

示例:

输入: [1,2,3]
输出:
[
  [1,2,3],
  [1,3,2],
  [2,1,3],
  [2,3,1],
  [3,1,2],
  [3,2,1]
]

来源：力扣（LeetCode）46
链接：https://leetcode-cn.com/problems/permutations
 */

func Permute(nums []int) [][]int {
	Perm(nums, 0, len(nums)-1)
	return res
}
var res = make([][]int, 0)
func Perm(nums []int, p, q int) {
	if p == q {
		var temp = make([]int, q+1)
		copy(temp, nums)
		res = append(res, temp)
	} else {
		for i := p; i <= q; i++ {
			swap(nums, p, i)
			Perm(nums, p+1, q)
			swap(nums, p, i)
		}
	}
}
func swap(arr []int, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}