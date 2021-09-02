package binarytree

import (
	"strings"
)

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

// Serialize serialize the binary tree
func (bt *BinaryTree) Serialize() string {
	builder := &strings.Builder{}

	var serl func(n *Node)
	serl = func(n *Node) {
		if n == nil {
			builder.WriteString("null,")
			return
		}

		builder.WriteString(n.Val.(string))
		builder.WriteString(",")
		serl(n.Left)
		serl(n.Right)
	}
	serl(bt.Root)

	return builder.String()
}

// Deserialize deserialize the binary tree
func Deserialize(data string) *BinaryTree {
	d := strings.Split(data, ",")
	return &BinaryTree{Root: deserl(&d)}
}

func deserl(data *[]string) *Node {
	nodeV := (*data)[0]
	*data = (*data)[1:]

	if nodeV == "null" {
		return nil
	}

	n := &Node{Val: nodeV}
	n.Left = deserl(data)
	n.Right = deserl(data)
	return n
}
