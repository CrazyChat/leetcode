package main

import (
	"fmt"
	"github.com/crazychat/leetcode/array"
)

func main() {
	arr := []int{0,1,1,2,4,4,1,3,3,2}
	arr2 := array.GetLeastNumbers(arr, 6)
	fmt.Println(arr2)
}
