package main

import (
	"fmt"
	"github.com/crazychat/leetcode/dynamic"
)

/*
[1,1,-6,1,-2,-4,4,-2,6,-6,0,6],[-3,-3,-6,-2,-6,-2,7,-9,-5,-7,-5,5,1]]
 */

func main() {
	triangle := [][]int{}
	triangle = append(triangle, []int{-7})
	triangle = append(triangle, []int{-2,1})
	triangle = append(triangle, []int{-5,-5,9})
	triangle = append(triangle, []int{-4,-5,4,4})
	triangle = append(triangle, []int{-6,-6,2,-1,-5})
	triangle = append(triangle, []int{3,7,8,-3,7,-9})
	triangle = append(triangle, []int{-9,-1,-9,6,9,0,7})
	triangle = append(triangle, []int{-7,0,-6,-8,7,1,-4,9})
	triangle = append(triangle, []int{-3,2,-6,-9,-7,-6,-9,4,0})
	triangle = append(triangle, []int{-8,-6,-3,-9,-2,-6,7,-5,0,7})
	triangle = append(triangle, []int{-9,-1,-2,4,-2,4,4,-1,2,-5,5})
	triangle = append(triangle, []int{1,1,-6,1,-2,-4,4,-2,6,-6,0,6})
	triangle = append(triangle, []int{-3,-3,-6,-2,-6,-2,7,-9,-5,-7,-5,5,1})
	x := dynamic.MinimumTotal(triangle)
	fmt.Println(x)
}