package array

/*
一个整型数组 nums 里除两个数字之外，其他数字都出现了两次。
请写程序找出这两个只出现一次的数字。要求时间复杂度是O(n)，空间复杂度是O(1)。
 */
func SingleNumbers3(nums []int) []int {
	/*
	思路：1. 整个数组异或运算出xorNumber, xorNumber二进制形式即是两个只出现一次数的不同位的地方
	     2. 计算lowbit为最低位的1, lowbit=x&(～x+1)=x&(−x)
		   3. 将所有数和lowbit与运算，结果为0的为1组，其他为1组，这样子就分成了两个组，其中每组满足SingleNumber1的例子
			 3. 运用SingleNumber1的解题找出每组只出现一次的数字
	 */
	length := len(nums)
	if length == 2 {
		return nums
	}
	xorNumber := 0
	num1 := 0
	num2 := 0
	for i := 0; i < length; i++ {
		xorNumber = xorNumber ^ nums[i]
	}
	mask := xorNumber & (-xorNumber)
	for i := 0; i < length; i++ {
		if nums[i]&mask == 0 {
			num1 = num1 ^ nums[i]
		} else {
			num2 = num2 ^ nums[i]
		}
	}
	return []int{num1, num2}
}