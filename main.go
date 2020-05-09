package main

import (
	"fmt"
	"github.com/crazychat/leetcode/stack"
)

func main() {
	queue := stack.MyQueueConstructor()
	queue.Push(1)
	queue.Push(2)
	fmt.Println(queue.Peek())
	fmt.Println(queue.Pop())
	fmt.Println(queue.Empty())
}




