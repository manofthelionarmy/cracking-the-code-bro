package problem3

import (
	"crackingthecode/ch4/problem2"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	arr = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60, 61, 62, 63, 64, 65, 66, 67, 68, 69, 70, 71, 72, 73, 74, 75, 76, 77, 78, 79, 80, 81, 82, 83, 84, 85, 86, 87, 88, 89, 90, 91, 92, 93, 94, 95, 96, 97, 98, 99, 100}
)

func TestListOfDepths(t *testing.T) {
	llArr := make([]*linkedList, 0)
	root := problem2.CreateMinimalBST(arr)
	ListOfDepths(root, &llArr, 0)
	// so, len(arr -1)
	// The level of the tree is the depth + 1
	totalDepth := int(math.Log2(float64(len(arr))))
	// len(Arr) == number of levels
	require.Equal(t, len(llArr), totalDepth+1)
}

func TestNewLinkedList(t *testing.T) {
	ll := NewLinkedList()
	require.NotNil(t, ll)
	require.Nil(t, ll.head, ll.tail)
}

func walkLinkedList(n *linkedListNode) []int {
	list := make([]int, 0)
	tempPtr := n
	for tempPtr != nil {
		list = append(list, tempPtr.val)
		tempPtr = tempPtr.next
	}
	return list
}

func countItems(n *linkedListNode) int {
	count := 0
	tempPtr := n
	for tempPtr != nil {
		count++
		tempPtr = tempPtr.next
	}
	return count
}

func TestAddNode(t *testing.T) {
	ll := NewLinkedList()
	// require.ElementsMatch(t, []int{1, 2}, walkLinkedList(ll.head))

	for i := range arr {
		ll.AddNode(arr[i])
	}
	require.ElementsMatch(t, arr, walkLinkedList(ll.head))
}
