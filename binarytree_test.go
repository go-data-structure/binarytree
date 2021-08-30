package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTreeCount(t *testing.T) {
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

	fmt.Println(bt.Count())
}
