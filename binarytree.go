package binarytree

// Node is binary tree node.
type Node struct {
	Left  *Node
	Right *Node
	Val   interface{}
}

// BinaryTree is binary tree
type BinaryTree struct {
	Root *Node
}

// Count return the count of binary tree nodes.
func (bt *BinaryTree) Count() int64 {
	return countBinaryTree(bt.Root)
}

func countBinaryTree(n *Node) int64 {
	if n == nil {
		return 0
	}
	return 1 + countBinaryTree(n.Left) + countBinaryTree(n.Right)
}
