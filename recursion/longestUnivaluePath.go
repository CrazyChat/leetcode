package recursion

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
	687 最长同值路径
*/

func longestUnivaluePath(root *TreeNode) int {
	result := 0
	f(root, &result)
	return result
}

func f(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}
	left := f(node.Left, result)
	right := f(node.Right, result)
	arrowLeft, arrowRight := 0, 0
	if node.Left != nil && node.Left.Val == node.Val {
		arrowLeft = left + 1
	}
	if node.Right != nil && node.Right.Val == node.Val {
		arrowRight = right + 1
	}
	*result = max(*result, arrowLeft+arrowRight)
	return max(arrowLeft, arrowRight)
}

func max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}
