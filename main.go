package main

import (
	"fmt"
	"github.com/crazychat/leetcode/array"
)

func main() {
	arr := [][]int{{0,1}}
	num := array.MaxAreaOfIsland(arr)
	fmt.Println(num)
}