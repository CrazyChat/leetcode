package oneWayStack

// 单调递增或递减栈
type OneWayStack []int

func NewOneWayStack() OneWayStack {
	return OneWayStack{}
}

func (s OneWayStack) Len() int {
	return len(s)
}

func (s OneWayStack) Top() int {
	return s[len(s)-1]
}

func (s OneWayStack) Bottom() int {
	return s[0]
}

func (s *OneWayStack) Push(val int) {
	*s = append(*s, val)
}

func (s *OneWayStack) Pop() int {
	old := *s
	index := old[len(old)-1]
	*s = old[:len(old)-1]
	return index
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}
