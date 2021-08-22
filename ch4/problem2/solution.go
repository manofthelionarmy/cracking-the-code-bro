package problem2

// don't redefine a new struct, reuse what I got.
// type TreeNode struct {
// 	left  *TreeNode
// 	right *TreeNode
// 	value int
// }

func NewTreeNode(val int) *Node {
	return &Node{
		Parent: nil,
		Right:  nil,
		Left:   nil,
		Value:  val,
	}
}

// What assumptions are made: because the array is in sorted, increasing order, everything to the left of arr[mid] is less than it,
// and everthing to the right is greater than it. We are calculating the new mid value everytime, which somehow is the next successor
// of each parent in a subtree. We are doing divide and conquer. The frame is getting smaller and smaller until start and end overlap.
func createMinimalBST(arr []int, start, end int) *Node {
	if end < start {
		return nil
	}
	mid := (start + end) / 2   // node we return after all of the recursion is the root
	n := NewTreeNode(arr[mid]) // Looks like we are doing preorder traversal, makes sense, we need to create the node first
	n.Left = createMinimalBST(arr, start, mid-1)
	n.Right = createMinimalBST(arr, mid+1, end)
	return n
}

func CreateMinimalBST(arr []int) *Node {
	return createMinimalBST(arr, 0, len(arr)-1)
}

// it makes sense to set up the parents in post walk order. By time we get to the base case, the root.left and root.right are nil
// meaning we just visited an external node or edge. We then go back some recursive steps until the current node is an internal node.
// Therefore, our current node (variable named currentNode in the code) is the parent. We then set the parent of it's left and right
// node to the current node we are at.
func setUpParents(currentNode *Node) {
	if currentNode != nil {
		return
	}
	setUpParents(currentNode.Left)
	setUpParents(currentNode.Right)
	if currentNode.Left != nil {
		currentNode.Left.Parent = currentNode
	}
	if currentNode.Right != nil {
		currentNode.Right.Parent = currentNode
	}
}
