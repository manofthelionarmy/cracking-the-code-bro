package problem4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewTree(t *testing.T) {

	// TODO: make root and nodes public accessible
	// bst := binarytree.NewBSTree()
	// bst.Insert(2)
	// bst.Insert(3)
	// bst.Insert(1)
	bst := NewBSTree()
	bst.Insert(2)
	require.NotNil(t, bst.Root())
	require.Nil(t, bst.Root().Left)
	require.Nil(t, bst.Root().Right)
	require.Equal(t, 2, bst.Root().Value)
	bst.Insert(1)
	require.NotNil(t, bst.Root().Left)
	require.Equal(t, 1, bst.Root().Left.Value)
	bst.Insert(3)
	require.NotNil(t, bst.Root().Right)
	require.Equal(t, 3, bst.Root().Right.Value)
}

func TestCheckBalance(t *testing.T) {
	// I alway think if you contruct a tree as a BSTtree, the height between
	// each subtree will always be less than or equal to 1 (log(2) == 1)
	bst := NewBSTree()
	bst.Insert(5)

	// left subtree
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(1)

	// right subtree
	bst.Insert(6)
	levelCount := make([]int, 0)
	require.Equal(t, true, CheckBalanced(bst.Root(), 0, &levelCount))
}
