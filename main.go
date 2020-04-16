package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure/oneWayStack"
)

func main() {
	heights := []int{1,2,1}
	fmt.Println(oneWayStack.NextGreaterElements(heights))
}