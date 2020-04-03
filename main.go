package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure"
)

func main() {
	tree := structure.NewTree()
	arr := []int{6, 3, 8, 2, 5, 1, 7}
	for _, v := range arr {
		tree.Insert(v)
	}
	tree.InOrder()
	fmt.Println(tree.Delete(6))
	tree.InOrder()
	tree.PreOrder()
}