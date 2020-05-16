package recursion

/*
543. 二叉树的直径
*/

var result int

func diameterOfBinaryTree(root *TreeNode) int {
	result := 1
	depth(root)
	return result - 1
}

func depth(node *TreeNode) int {
	if node == nil {
		return 0
	}
	left := depth(node.Left)
	right := depth(node.Right)
	result = max(result, left+right+1)
	return max(left, right) + 1
}
