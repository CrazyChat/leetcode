package main

import (
	"fmt"
	"github.com/crazychat/leetcode/meituan"
)

func main() {
	arr := []int{20, 1, 1, 1, 1, 1, 1, 1}
	sum := 119
	roomsCount := len(arr)
	fmt.Println(meituan.CircleLeaveVal(roomsCount, sum, arr))
}