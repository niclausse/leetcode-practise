package tree

type BinarySearchTree struct {
	root *BinaryTreeNode
}

func (b *BinarySearchTree) search(data int) *BinaryTreeNode {
	node := b.root

	for node != nil {
		if data < node.Val {
			node = node.Lft
		} else if data > node.Val {
			node = node.Rgt
		} else {
			return node
		}
	}

	return nil
}

func (b *BinarySearchTree) insert(data int) {
	if b.root == nil {
		b.root = &BinaryTreeNode{Val: data}
		return
	}

	node := b.root
	for node != nil {
		if node.Val < data {
			if node.Rgt == nil {
				node.Rgt = &BinaryTreeNode{Val: data}
				return
			}

			node = node.Rgt
		} else if node.Val > data {
			if node.Lft == nil {
				node.Lft = &BinaryTreeNode{Val: data}
				return
			}

			node = node.Lft
		} else {
			return
		}
	}
}
