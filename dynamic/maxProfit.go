package dynamic

/*
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。
如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。
注意：你不能在买入股票前卖出股票。

示例 1:
输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

来源：力扣（LeetCode）121 买卖股票的最佳时机
链接：https://leetcode-cn.com/problems/best-time-to-buy-and-sell-stock
 */

// 单调栈
func MaxProfit(prices []int) int {
	maxIncome = 0
	length := len(prices)
	stack := NewStack(length)
	// 开始循环
	for i := 0; i < length; i++ {
		stack.Push(prices[i])
	}
	stack.Push(0)
	return maxIncome
}
// 记录最大收益
var maxIncome int = 0
// 定义一个适合的栈, 0 是栈底，top 是栈顶
type Stack struct {
	data []int
	top int
}

func NewStack(len int) Stack {
	return Stack{make([]int, len+1), 0}
}

func (s *Stack) Push(val int) {
	// 空栈或者大于栈顶元素，直接放入
	if s.top == 0 || val > s.data[s.top-1] {
		s.data[s.top] = val
		s.top = s.top + 1
		return
	}
	// 小于，先记录最大值，再弹出，再重新入栈
	income := s.data[s.top-1] - s.data[0]
	if income > maxIncome {
		maxIncome = income
	}
	s.top = s.top - 1	// 出栈
	s.Push(val)
}
func (s *Stack) Pop() {
	s.top = s.top - 1
}

