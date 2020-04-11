package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure"
)

func main() {
	stack := structure.NewLinkStack()
	for i := 0; i < 10; i++ {
		stack.Push(i)
	}

	for !stack.IsEmpty() {
		value, _ := stack.Pop()
		fmt.Printf("%d ", value)
	}
	value, err := stack.Top()
	if err != nil {
		fmt.Println()
		fmt.Println(err)
	} else {
		fmt.Println(value)
	}
}