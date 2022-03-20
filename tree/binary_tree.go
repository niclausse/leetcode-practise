package tree

import "fmt"

// BinaryTreeNode 二叉树节点
type BinaryTreeNode struct {
	Val int
	Lft *BinaryTreeNode
	Rgt *BinaryTreeNode
}

// 前序遍历-递归
func preOrder1(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	fmt.Println(root.Val)
	preOrder1(root.Lft)
	preOrder1(root.Rgt)
}

// 前序遍历-迭代
func preorder(root *BinaryTreeNode) (result []int) {
	var (
		stack []*BinaryTreeNode
		cur   = root
	)

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			result = append(result, cur.Val)
			stack = append(stack, cur)
			cur = cur.Lft
		}

		// 取栈顶元素
		cur = stack[len(stack)-1].Rgt
		// 出栈
		stack = stack[:len(stack)-1]
	}

	return result
}

// 中序遍历-递归
func inOrder1(root *BinaryTreeNode) {
	if root == nil {
		return
	}

	inOrder1(root.Lft)
	fmt.Println(root.Val)
	inOrder1(root.Rgt)
}

// 中序遍历-迭代
func inorder(root *BinaryTreeNode) []int {
	var (
		result []int
		stack  []*BinaryTreeNode
		cur    = root
	)

	for cur != nil || len(stack) > 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Lft
		}

		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, cur.Val)
		cur = cur.Rgt
	}
	return result
}

// 后序遍历
func postOrder(root *BinaryTreeNode) {
	postOrder(root.Lft)
	postOrder(root.Rgt)
	fmt.Println(root.Val)
}

// TODO 层序遍历
