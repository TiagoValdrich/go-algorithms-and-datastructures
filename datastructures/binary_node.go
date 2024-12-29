package datastructures

type BinaryNode struct {
	Value  int
	Parent *BinaryNode
	Left   *BinaryNode
	Right  *BinaryNode
}

func (bn *BinaryNode) IterSubtree(path *[]*BinaryNode) []*BinaryNode {
	if path == nil {
		tempPath := make([]*BinaryNode, 0)
		path = &tempPath
	}

	*path = append(*path, bn)

	if bn.Left != nil {
		bn.Left.IterSubtree(path)
	}

	if bn.Right != nil {
		bn.Right.IterSubtree(path)
	}

	return *path
}
