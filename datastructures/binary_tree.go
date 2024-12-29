package datastructures

type BinaryTree struct {
	Root *BinaryNode
}

func NewBinaryTree(root BinaryNode) *BinaryTree {
	return &BinaryTree{
		Root: &root,
	}
}
