package problem2

type Node struct {
	Value  int
	Parent *Node
	Left   *Node
	Right  *Node
}

type BSTree struct {
	root *Node
}

func NewBSTree() *BSTree {
	return &BSTree{root: nil}
}

func (bst *BSTree) Root() *Node {
	return bst.root
}

func (bst *BSTree) Insert(val int) {
	temPtr := bst.root
	var trailingPtr *Node
	for temPtr != nil {
		trailingPtr = temPtr
		if val <= temPtr.Value {
			temPtr = temPtr.Left
		} else {
			temPtr = temPtr.Right
		}
	}

	newEl := &Node{Value: val, Parent: trailingPtr, Left: nil, Right: nil}
	if trailingPtr == nil {
		bst.root = newEl
		return
	}
	if val <= trailingPtr.Value {
		trailingPtr.Left = newEl
	} else {
		trailingPtr.Right = newEl
	}

}
