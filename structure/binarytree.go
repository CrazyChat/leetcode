package structure

import (
	"errors"
	"fmt"
)

/*
This is a binary search tree
 */

type Tree struct {
	root *TreeNode
}
type TreeNode struct {
	value int
	left *TreeNode
	right *TreeNode
}
// NewTree 创建一颗空树
func NewTree() Tree {
	return Tree{root: nil}
}
// Insert 插入结点
func (tree *Tree) Insert(val int) {
	newNode := &TreeNode{value: val}
	// 如果树为空，则直接作为根结点就可以了
	if tree.root == nil {
		tree.root = newNode
		return
	}
	// 其他情况需要左右查找
	curNode := tree.root
	for curNode != nil {
		// 比结点大，向右边
		if val > curNode.value {
			// 如果右结点为空，那就插入该位置，否则继续查找
			if curNode.right == nil {
				curNode.right = newNode
				return
			} else {
				curNode = curNode.right
			}
		} else {	// 比结点小，向左边
			// 如果左结点为空，那就插入该位置，否则继续查找
			if curNode.left == nil {
				curNode.left = newNode
				return
			} else {
				curNode = curNode.left
			}
		}
	}
}
// Find 查找结点
func (tree *Tree) Find(val int) *TreeNode {
	curNode := tree.root
	for curNode != nil {
		// 等于该结点，返回
		if val == curNode.value {
			return curNode
		} else if val < curNode.value {
			// 小于结点，向左边
			curNode = curNode.left
		} else {
			// 大于结点，向右边
			curNode = curNode.right
		}
	}
	return nil
}
// Delete 删除结点
func (tree *Tree) Delete(val int) error {
	// 树为空，找不到错误
	if tree.root == nil {
		return errors.New("the tree is empty")
	}
	// curNode 代表要删除的节点，curNodeFather 代表要curNode的父节点
	curNode := tree.root
	var curNodeFather *TreeNode
	for curNode != nil && val != curNode.value {
		curNodeFather = curNode
		if val > curNode.value {
			curNode = curNode.right
			continue
		}
		curNode = curNode.left
	}
	// 没有找到
	if curNode == nil {
		return errors.New("the node no exists")
	}
	// 删除的节点如果有两个子节点，处理成删除另一个节点
	if curNode.left != nil && curNode.right != nil {
		// 找到curNode右子树的最小节点
		minNode := curNode.right
		minNodeFather := curNode
		for minNode.left != nil {
			minNodeFather = minNode
			minNode = minNode.left
		}
		// 将最小节点的值替换curNode的值
		curNode.value = minNode.value
		// 并且将curNode换到最小节点，即现在变成了要删除最小节点
		curNode = minNode
		curNodeFather = minNodeFather
	}
	// 删除节点肯定是叶子节点或者仅有一个子节点，因为上面的情况已经处理转化了
	// 先保存curNode的子节点(如果有的话)
	var curNodeChild *TreeNode
	if curNode.left != nil {
		curNodeChild = curNode.left
	} else if curNode.right != nil {
		curNodeChild = curNode.right
	}
	// 开始删除
	if curNodeFather == nil {	// 删除的是根节点
		tree.root = curNodeChild
	} else if curNodeFather.left == curNode {
		curNodeFather.left = curNodeChild
	} else {
		curNodeFather.right = curNodeChild
	}
	return nil
}
// PreOrder 先序遍历
func (tree *Tree) PreOrder() {
	PreNode(tree.root)
	fmt.Println()
}
func PreNode(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.value)
	PreNode(node.left)
	PreNode(node.right)
}
// InOrder 中序遍历
func (tree *Tree) InOrder() {
	InNode(tree.root)
	fmt.Println()
}
func InNode(node *TreeNode) {
	if node == nil {
		return
	}
	InNode(node.left)
	fmt.Printf("%d ", node.value)
	InNode(node.right)
}
// PostOrder 后序遍历
func (tree *Tree) PostOrder() {
	PostNode(tree.root)
	fmt.Println()
}
func PostNode(node *TreeNode) {
	if node == nil {
		return
	}
	fmt.Printf("%d ", node.value)
	PostNode(node.left)
	PostNode(node.right)
}
// Height 获取树的高度
func (tree *Tree) Height() int {
	return getHeight(tree.root)
}
func getHeight(node *TreeNode) int {
	if node == nil {
		return 0
	}
	// 返回左右边比较高的节点+1
	leftHeight := getHeight(node.left)
	rightHeight := getHeight(node.right)
	if leftHeight >= rightHeight {
		return leftHeight + 1
	}
	return rightHeight + 1
}
// Min 最小值
func (tree *Tree) Min() (int, error) {
	curNode := tree.root
	if curNode == nil {
		return 0, errors.New("The tree is empty.")
	}
	for curNode.left != nil {
		curNode = curNode.left
	}
	return curNode.value, nil
}
// Max 最大值
func (tree *Tree) Max() (int, error) {
	curNode := tree.root
	if curNode == nil {
		return 0, errors.New("The tree is empty.")
	}
	for curNode.right != nil {
		curNode = curNode.right
	}
	return curNode.value, nil
}