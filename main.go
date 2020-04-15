package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure/oneWayStack"
)

func main() {
	heights := []int{0,1,0,1}
	fmt.Println(oneWayStack.Trap(heights))
}