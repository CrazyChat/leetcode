package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure/oneWayStack"
)

func main() {
	heights := []int{0,9}
	fmt.Println(oneWayStack.LargestRectangleArea(heights))
}