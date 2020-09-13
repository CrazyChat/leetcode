package tree

import "github.com/crazychat/leetcode/structure"

const (
	SEP = ","
	NULL = "#"
)

func serialize(node *TreeNode) string {
	if node == nil {
		return ""
	}
	// 初始化队列
	queue := structure.NewQueue(100)
	queue.Enqueue(node.Val)
	
}
