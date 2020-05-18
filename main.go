package main

import (
	"fmt"
	"github.com/crazychat/leetcode/mix"
)

func main() {
	num := []int{0, 1, 4, 5, 8, 9}
	for _, v := range num {
		fmt.Println(v, mix.Search(num, v))
	}
	fmt.Println(mix.Search(num, 6))
}
