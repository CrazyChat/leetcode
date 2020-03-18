package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure"
)

func main() {
	stack := structure.CreateArrayStack()
	for i := 0; i < 20; i++ {
		stack.Push(3)
	}
	fmt.Println(stack)
}
