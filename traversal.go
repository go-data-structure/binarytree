package binarytree

// TraversalFunc is the function which handle node while traversal binary tree.
type TraversalFunc func(n *Node)

// PreorderTraversal preorder traversal binary tree.
func (bt *BinaryTree) PreorderTraversal(tf TraversalFunc) {
	preorderTraversal(bt.Root, tf)
}

func preorderTraversal(n *Node, tf TraversalFunc) {
	if n == nil {
		return
	}
	tf(n)
	preorderTraversal(n.Left, tf)
	preorderTraversal(n.Right, tf)
}

// PostorderTraversal postorder traversal binary tree.
func (bt *BinaryTree) PostorderTraversal(tf TraversalFunc) {
	postorderTraversal(bt.Root, tf)
}

func postorderTraversal(n *Node, tf TraversalFunc) {
	if n == nil {
		return
	}
	postorderTraversal(n.Left, tf)
	postorderTraversal(n.Right, tf)
	tf(n)
}

// InorderTraversal inorder traversal binary tree.
func (bt *BinaryTree) InorderTraversal(tf TraversalFunc) {
	inorderTraversal(bt.Root, tf)
}

func inorderTraversal(n *Node, tf TraversalFunc) {
	if n == nil {
		return
	}
	inorderTraversal(n.Left, tf)
	tf(n)
	inorderTraversal(n.Right, tf)
}
