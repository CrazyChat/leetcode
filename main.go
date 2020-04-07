package main

import (
	"fmt"
	"github.com/crazychat/leetcode/array"
)

func main() {
	nums1 := []int{4,1,2}
	nums2 := []int{2}
	fmt.Println(array.NextGreaterElement(nums1, nums2))
}