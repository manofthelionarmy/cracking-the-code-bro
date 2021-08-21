package problem2

type node struct {
	value  int
	parent *node
	left   *node
	right  *node
}

type BSTree struct {
	root *node
}

func NewBSTree() *BSTree {
	return &BSTree{root: nil}
}

func (bst *BSTree) Root() *node {
	return bst.root
}

func (bst *BSTree) Insert(val int) {
	temPtr := bst.root
	var trailingPtr *node
	for temPtr != nil {
		trailingPtr = temPtr
		if val <= temPtr.value {
			temPtr = temPtr.left
		} else {
			temPtr = temPtr.right
		}
	}

	newEl := &node{value: val, parent: trailingPtr, left: nil, right: nil}
	if trailingPtr == nil {
		bst.root = newEl
		return
	}
	if val <= trailingPtr.value {
		trailingPtr.left = newEl
	} else {
		trailingPtr.right = newEl
	}

}
