package problem5

func inorderTraversalHelper(root *Node, prevVal interface{}) bool {
	if root == nil {
		return true
	}
	if !inorderTraversalHelper(root.Left, prevVal) {
		return false
	}

	if prevVal != nil && prevVal.(int) >= root.Value {
		return false
	}
	prevVal = root.Value
	return inorderTraversalHelper(root.Right, prevVal)
}

func InorderTraversalSolution(root *Node) bool {
	return inorderTraversalHelper(root, nil)
}
