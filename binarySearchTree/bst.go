package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{data: i}
}

// Insert inserts an int into the BinarySearchTree
func (bst *BinarySearchTree) Insert(i int) {
	if i > bst.data {
		if bst.right == nil {
			bst.right = NewBst(i)
		} else {
			bst.right.Insert(i)
		}
	}
	if i <= bst.data {
		if bst.left == nil {
			bst.left = NewBst(i)
		} else {
			bst.left.Insert(i)
		}
	}
}

func (bst *BinarySearchTree) SortedData() []int {
	var sorted []int
	bst.sortRec(&sorted)
	return sorted
}

func (bst *BinarySearchTree) sortRec(acc *[]int) {
	// do in order
	if bst.left != nil {
		bst.left.sortRec(acc)
	}
	*acc = append(*acc, bst.data)
	if bst.right != nil {
		bst.right.sortRec(acc)
	}
}
