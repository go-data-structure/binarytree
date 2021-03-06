package binarytree

import (
	"testing"
)

// newTestBinaryTree new binary tree for test
func newTestBinaryTree() *BinaryTree {
	bt := new(BinaryTree)
	bt.Root = &Node{Val: "F"}
	bt.Root.Left = &Node{Val: "B"}
	bt.Root.Right = &Node{Val: "G"}

	bt.Root.Left.Left = &Node{Val: "A"}
	bt.Root.Left.Right = &Node{Val: "D"}

	bt.Root.Left.Right.Left = &Node{Val: "C"}
	bt.Root.Left.Right.Right = &Node{Val: "E"}

	bt.Root.Right.Right = &Node{Val: "I"}
	bt.Root.Right.Right.Left = &Node{Val: "H"}
	return bt
}

func TestPreorderTraversal(t *testing.T) {
	bt := newTestBinaryTree()
	bt.PreorderTraversal(func(n *Node) {
		t.Log(n.Val.(string))
	})
}

func TestPostorderTraversal(t *testing.T) {
	bt := newTestBinaryTree()
	bt.PostorderTraversal(func(n *Node) {
		t.Log(n.Val.(string))
	})
}

func TestInorderTraversal(t *testing.T) {
	bt := newTestBinaryTree()
	bt.InorderTraversal(func(n *Node) {
		t.Log(n.Val.(string))
	})
}
