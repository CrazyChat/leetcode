package main

import (
	"fmt"
	"github.com/crazychat/leetcode/array"
)

func main() {
	arr := []int{7,3,2,1,5,6,4,23}
	a := array.FindKthLargest(arr, 2)
	fmt.Println(a)
	fmt.Println(arr)
}
