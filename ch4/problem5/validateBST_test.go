package problem5

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateBST(t *testing.T) {
	bst := NewBSTree()
	bst.Insert(5)
	bst.Insert(3)
	bst.Insert(9)
	bst.Insert(1)
	bst.Insert(6)
	// tree reads as: 1, 3, 5, 6, 9
	require.Equal(t, true, ValidateBST(bst.Root()))
	// now tree reads as: 1, 3, 5, 6, 9, 7
	bst.root.Right.Right = NewNode(7, nil, nil, bst.root.Right)
	require.Equal(t, false, ValidateBST(bst.Root()))
}
