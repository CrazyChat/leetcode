package structure

import "fmt"

type BinarytreeType int

type Node struct {
	Data BinarytreeType
	Left *Node
	Right *Node
}

// PreOrder 前序遍历
func (tree *Node) PreOrder() {
	if tree == nil {
		return
	}
	fmt.Println(tree.Data)
	tree.Left.PreOrder()
	tree.Right.PreOrder()
}
// BinarySearchTree 查找结点
func (tree *Node) BinarySearchTree(data BinarytreeType) *Node {
	p := tree
	for p != nil {
		if data < p.Data {
			p = p.Left
		} else if data > p.Data {
			p = p.Right
		} else {
			return p
		}
	}
	return nil
}
// Insert 插入结点
func (tree *Node) Insert(data BinarytreeType) {
	if tree == nil {
		fmt.Println("tree no exists")
		return
	}
	p := tree
	for p != nil {
		if data > p.Data {
			if p.Right == nil {
				p.Right = &Node{data, nil, nil}
				return
			}
			p = p.Right
		} else {
			if p.Left == nil {
				p.Left = &Node{data, nil, nil}
				return
			}
			p = p.Left
		}
	}
}
// Delete 删除一个结点
func (tree *Node) Delete(data BinarytreeType) *Node {
	if tree == nil {
		fmt.Println("tree is null")
	}
	p := tree
	var pPre *Node
	for p != nil && p.Data != data {
		pPre = p
		if data > p.Data {
			p = p.Right
		} else {
			p = p.Left
		}
	}
	if p == nil {
		fmt.Println("data no exists")
	}
	// 要删除的节点有两个子节点
	if p.Left != nil && p.Right != nil {
		// 找出右子树中最小的值
		minP := p.Right
		minPPre := p
		for minP.Left != nil {
			minPPre = minP
			minP = minP.Left
		}
		p.Data = minP.Data	// 将minP的数据替换到p中
		// 下面就变成了删除minP了, 注意此时minP可能有右节点不能让minPPre指向nil来删除
		p = minP
		pPre = minPPre
	}
	// 删除节点是叶子节点或者仅有一个子节点
	var child *Node		// p的子节点
	if p.Left != nil {
		child = p.Left
	} else if p.Right != nil {
		child = p.Right
	} else {
		child = nil
	}
	if pPre == nil {
		// 删除的是根节点
		return child
	} else if pPre.Left == p {
		pPre.Left = child
		return tree
	}
	pPre.Right = child
	return tree
}


