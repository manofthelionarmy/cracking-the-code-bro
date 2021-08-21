package problem2

// don't redefine a new struct, reuse what I got.
// type TreeNode struct {
// 	left  *TreeNode
// 	right *TreeNode
// 	value int
// }

func NewTreeNode(val int) *node {
	return &node{
		parent: nil,
		right:  nil,
		left:   nil,
		value:  val,
	}
}

// What assumptions are made: because the array is in sorted, increasing order, everything to the left of arr[mid] is less than it,
// and everthing to the right is greater than it. We are calculating the new mid value everytime, which somehow is the next successor
// of each parent in a subtree. We are doing divide and conquer. The frame is getting smaller and smaller until start and end overlap.
func createMinimalBST(arr []int, start, end int) *node {
	if end < start {
		return nil
	}
	mid := (start + end) / 2   // node we return after all of the recursion is the root
	n := NewTreeNode(arr[mid]) // Looks like we are doing preorder traversal, makes sense, we need to create the node first
	n.left = createMinimalBST(arr, start, mid-1)
	n.right = createMinimalBST(arr, mid+1, end)
	return n
}

func CreateMinimalBST(arr []int) *node {
	return createMinimalBST(arr, 0, len(arr)-1)
}

// it makes sense to set up the parents in post walk order. By time we get to the base case, the root.left and root.right are nil
// meaning we just visited an external node or edge. We then go back some recursive steps until the current node is an internal node.
// Therefore, our current node (variable named currentNode in the code) is the parent. We then set the parent of it's left and right
// node to the current node we are at.
func setUpParents(currentNode *node) {
	if currentNode != nil {
		return
	}
	setUpParents(currentNode.left)
	setUpParents(currentNode.right)
	if currentNode.left != nil {
		currentNode.left.parent = currentNode
	}
	if currentNode.right != nil {
		currentNode.right.parent = currentNode
	}
}
