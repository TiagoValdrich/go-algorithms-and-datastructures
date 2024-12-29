package datastructures_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tiagovaldrich/go-algorithms-and-datastructures/datastructures"
)

func createBinaryTestTree() *datastructures.BinaryTree {
	nodeA := datastructures.BinaryNode{
		Value: 50,
	}
	// Left Subtree
	nodeB := datastructures.BinaryNode{
		Parent: &nodeA,
		Value:  25,
	}
	// Right Subtree
	nodeC := datastructures.BinaryNode{
		Parent: &nodeA,
		Value:  75,
	}

	nodeA.Left = &nodeB
	nodeA.Right = &nodeC

	nodeD := datastructures.BinaryNode{
		Parent: &nodeB,
		Value:  15,
	}
	nodeE := datastructures.BinaryNode{
		Parent: &nodeB,
		Value:  30,
	}
	nodeF := datastructures.BinaryNode{
		Parent: &nodeC,
		Value:  100,
	}
	nodeG := datastructures.BinaryNode{
		Parent: &nodeD,
		Value:  5,
	}

	nodeB.Left = &nodeD
	nodeB.Right = &nodeE
	nodeC.Right = &nodeF
	nodeD.Left = &nodeG

	return datastructures.NewBinaryTree(nodeA)
}

func TestBinaryTree(t *testing.T) {
	t.Parallel()

	t.Run("Should iter the binary tree in traversal order", func(t *testing.T) {
		binaryTree := createBinaryTestTree()
		pathTraversed := binaryTree.Root.IterSubtree(nil)

		assert.Equal(t, 50, pathTraversed[0].Value)
		assert.Equal(t, 25, pathTraversed[1].Value)
		assert.Equal(t, 15, pathTraversed[2].Value)
		assert.Equal(t, 5, pathTraversed[3].Value)
		assert.Equal(t, 30, pathTraversed[4].Value)
		assert.Equal(t, 75, pathTraversed[5].Value)
		assert.Equal(t, 100, pathTraversed[6].Value)
	})
}
