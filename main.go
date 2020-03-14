package main

import (
	"fmt"
	"github.com/crazychat/leetcode/sortMethod"
)

func main() {
	arr := []int{6,11,3,9,8}
	sortMethod.QuickSort(arr)
	fmt.Println(arr)
}
