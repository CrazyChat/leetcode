package meituan

import "fmt"

/*
空间回廊
时间限制：C/C++语言 1000MS；其他语言 3000MS
内存限制：C/C++语言 65536KB；其他语言 589824KB
题目描述：
有一款叫做空间回廊的游戏，游戏中有着n个房间依次相连，如图，1号房间可以走到2号房间，以此类推，n号房间可以走到1号房间。
这个游戏的最终目的是为了在这些房间中留下尽可能多的烙印，在每个房间里留下烙印所花费的法力值是不相同的，已知他共有m点法力值，这些法力是不可恢复的。
小明刚接触这款游戏，所以只会耿直的玩，所以他的每一个行动都是可以预料的：
1. 一开始小明位于1号房间。
2. 如果他剩余的法力能在当前的房间中留下一个烙印，那么他就会毫不犹豫的花费法力值。
3. 无论是否留下了烙印，下一个时刻他都会进入下一个房间，如果当前位于i房间，则会进入i+1房间，如果在n号房间则会进入1号房间。
4. 当重复经过某一个房间时，可以再次留下烙印。
很显然，这个游戏是会终止的，即剩余的法力值不能在任何房间留下烙印的时候，游戏终止。请问他共能留下多少个烙印。

输入
输入第一行有两个正整数n和m，分别代表房间数量和小明拥有的法力值。(1<=n<=100000,1<=m<=10^18)
输入第二行有n个正整数，分别代表1~n号房间留下烙印的法力值花费。(1<=a_i<=10^9)
输出
输出仅包含一个整数，即最多能留下的烙印。
样例输入
4 21
2 1 4 3
样例输出
9
提示
样例解释：
显然是所有房间都留下两个烙印，然后剩下1点法力值，仅能在2号房间再留下一个烙印.
规则
请尽量在全场考试结束10分钟前调试程序，否则由于密集排队提交，可能查询不到编译结果
点击“调试”亦可保存代码
编程题可以使用本地编译器，此页面不记录跳出次数
 */

/*
sum  初始法力值
roomsCount 房间数
arr 每个房间需要的法力值
 */
func CircleLeaveVal(roomsCount, sum int, arr []int) int {
		// 特例判断
		if roomsCount < 1 {
			return 0
		}
		if roomsCount == 1 {
			return sum/arr[0]
		}
		var leaveRoomsCount int	// 在多少个房间留下了法力值
		var circleSum int	//  走一圈在每个房间留下了法力值需要多少
		for i := range arr {
			circleSum += arr[i]
		}
		/*
			1. 全部留下法力可以多少圈，更新leaveRoomsCount
			2. 取模，检查sum: if sum == 0 结束
			3. 走一圈标记没法留下法力值的房间, 同时更新circleSum
			4. 删除没有留下法力值的房间
			5. 重新循环
		 */
		for len(arr) >= 1 {
			roomsCount = len(arr)
			// 1. 每个房间都留下法力可以多少圈，更新leaveRoomsCount
			circleCount := sum/circleSum
			leaveRoomsCount += circleCount * roomsCount
			// 2. 取模，检查sum: if sum == 0 结束
			sum = sum % circleSum
			if sum == 0 {
				return leaveRoomsCount
			}
			// 3. 走一圈标记没法留下法力值的房间
			isLeaveSum := make([]bool, roomsCount)
			circleSum = 0
			for i := 0; i < roomsCount; i++ {
				// 留下法力值
				if sum >= arr[i] {
					sum -= arr[i]
					leaveRoomsCount++
					isLeaveSum[i] = true
					circleSum+= arr[i]
				}
			}
			// 3. 删除没有留下法力值的房间
			arr = removeFalse(arr, isLeaveSum)
			fmt.Println(arr, circleSum)
		}
		return leaveRoomsCount
}

/*
删除 arr 对应 isLeaveSum 中值为 false 的位置
输入:
arr: [1,2,3,4,5,6]
isLeaveSum: [false, false, true, false, true, false]
输出:
arr: [3, 5]
 */
func removeFalse(arr []int, isLeaveSum []bool) []int {
	i := 0
	j := 0
	for i < len(arr) && j < len(arr) {
		// 找 false
		for i < len(arr) && isLeaveSum[i] == true {
			i++
		}
		if i >= len(arr) {
			break
		}
		// 从 false 后面找 true
		j = i+1
		for j < len(arr) && isLeaveSum[j] == false {
			j++
		}
		if j >= len(arr) {
			break
		}
		arr[i], arr[j] = arr[j], arr[i]
		isLeaveSum[i], isLeaveSum[j] = isLeaveSum[j], isLeaveSum[i]
		i++
	}
	return arr[:i]
}
