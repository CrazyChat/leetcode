package main

import (
	"fmt"
	"github.com/crazychat/leetcode/structure"
)

func main() {
	obj := structure.ConstructorMinStack()
	obj.Push(4)
	obj.Push(3)
	obj.Push(1)
	obj.Push(1)
	obj.Push(2)
	fmt.Println("最小值：", obj.GetMin())
	obj.Pop()
	fmt.Println("最小值：", obj.GetMin())
	obj.Pop()
	fmt.Println("最小值：", obj.GetMin())
	obj.Pop()
	fmt.Println("最小值：", obj.GetMin())
}