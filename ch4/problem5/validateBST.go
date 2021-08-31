package problem5

type Node struct {
	Value  int
	Left   *Node
	Right  *Node
	Parent *Node
}

type BSTree struct {
	root *Node
}

func NewNode(val int, left, right, parent *Node) *Node {
	return &Node{Value: val, Left: left, Right: right, Parent: parent}
}

func NewBSTree() *BSTree {
	return &BSTree{root: nil}
}

func (bst *BSTree) Root() *Node {
	return bst.root
}

func (bst *BSTree) Insert(val int) {
	currentPtr := bst.root

	var trailingPtr *Node
	for currentPtr != nil {
		trailingPtr = currentPtr
		if currentPtr.Value >= val {
			currentPtr = currentPtr.Left
		} else {
			currentPtr = currentPtr.Right
		}
	}
	if trailingPtr == nil {
		bst.root = NewNode(val, nil, nil, nil)
	} else {
		element := NewNode(val, nil, nil, trailingPtr)
		if element.Value <= trailingPtr.Value {
			trailingPtr.Left = element
		} else {
			trailingPtr.Right = element
		}
	}
}

func validateHelper(root *Node, max interface{}, min interface{}) bool {
	if root == nil {
		return true
	}
	// If the converse is true, we return false
	if (max != nil && root.Value >= max.(int)) || (min != nil && root.Value <= min.(int)) {
		return false
	}
	return validateHelper(root.Left, root.Value, min) && validateHelper(root.Right, max, root.Value)
}

// My solution is the same as the book :)
func ValidateBST(root *Node) bool {
	return validateHelper(root, nil, nil)
}
