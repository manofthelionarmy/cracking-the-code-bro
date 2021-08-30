package problem4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	input = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
)

func InorderTraversal(root *Node, arr *[]int) {
	if root == nil {
		return
	}
	InorderTraversal(root.Left, arr)
	*arr = append(*arr, root.Value)
	InorderTraversal(root.Right, arr)
}

func TestSolution(t *testing.T) {

	// Draw tree out
	bst := NewBSTree()
	bst.Insert(5)

	// left subtree
	bst.Insert(3)
	bst.Insert(2)
	bst.Insert(4)
	bst.Insert(1)
	bst.Insert(0)

	// right subtree
	bst.Insert(6)
	require.Equal(t, false, CheckBalancedSolution(bst.Root()))
	levelCount := make([]int, 0)
	require.NotEqual(t, false, CheckBalanced(bst.Root(), 0, &levelCount), "We expect this one to return true, but it's incorrect")
	bst.Insert(7)
	bst.Insert(8)
	bst.Insert(9)
	arr := make([]int, 0)
	InorderTraversal(bst.Root(), &arr)
	require.ElementsMatch(t, input, arr)
	// why is it false?
	// Now I know! the difference between one of the subtrees was greater than 1
	require.Equal(t, false, CheckBalancedSolution(bst.Root()))
}
