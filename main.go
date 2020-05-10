package main

import (
	"fmt"
	"github.com/crazychat/leetcode/queue"
)

func main() {
	//tasks := []byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L', 'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'}
	tasks := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	fmt.Println(queue.LeastInterval(tasks, 2))
}
