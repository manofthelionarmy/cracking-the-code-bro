package problem4

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

func CheckBalanced(root *Node, level int, levelCount *[]int) bool {
	if root == nil {
		return true
	}
	if len(*levelCount) != (level + 1) {
		*levelCount = append(*levelCount, 0)
	}
	(*levelCount)[level] += 1
	value := 1 << level
	if (*levelCount)[level] > value {
		return false
	}
	return CheckBalanced(root.Left, level+1, levelCount) && CheckBalanced(root.Right, level+1, levelCount)
}
