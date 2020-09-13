package tree

import "fmt"

// preorderTraversal 递归前序遍历
func PreorderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Println(root.Val)
	PreorderTraversal(root.Left)
	PreorderTraversal(root.Right)
}
// preorderTraversal2 非递归前序遍历
func preorderTraversal2(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 {
		nums := len(stack)
		for i := 0; i < nums; i++ {
			temp := stack[0]
			stack = stack[1:]
			result = append(result, temp.Val)
			if temp.Left != nil {
				stack = append(stack, temp.Left)
			}
			if temp.Right != nil {
				stack = append(stack, temp.Right)
			}
		}
	}
	return result
}
// inorderTraversal 非递归中序遍历
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	if root == nil {
		return result
	}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) != 0 || root != nil {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// 弹出
		tempNode := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, tempNode.Val)
		root = tempNode.Right
	}
	return result
}
// postorderTraversal 非递归后序遍历
//func postorderTraversal(root *TreeNode) []int {
//	result := make([]int, 0)
//	if root == nil {
//		return result
//	}
//	stack := make([]*TreeNode, 0)
//	// 通过 lastVisit 标识右子节点是否已经弹出
//	var lastVisit *TreeNode
//
//}
