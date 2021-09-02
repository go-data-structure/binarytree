package binarytree

import (
	"fmt"
	"testing"
)

func TestBinaryTreeCount(t *testing.T) {
	bt := newTestBinaryTree()

	fmt.Println(bt.Count())
}

func TestBinaryTreeSerialize(t *testing.T) {
	bt := newTestBinaryTree()

	t.Log(bt.Serialize())
}

func TestBinaryTreeDeserialize(t *testing.T) {
	bt := newTestBinaryTree()
	s := bt.Serialize()

	nbt := Deserialize(s)
	t.Log(nbt.Serialize())
}
