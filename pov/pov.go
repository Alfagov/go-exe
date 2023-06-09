package pov

type Tree struct {
	children []*Tree
	parent   *Tree
	value    string
}

func NewTree(value string, children ...*Tree) *Tree {
	tree := &Tree{value: value, children: children}
	for _, child := range children {
		child.parent = tree
	}
	return tree
}

func (tr *Tree) connectedNodes() []*Tree {
	if tr.parent != nil {
		return append(tr.children, tr.parent)
	}
	return tr.children
}

func (tr *Tree) findNodeByValue(value string) *Tree {
	if tr.value == value {
		return tr
	}
	for _, child := range tr.children {
		if node := child.findNodeByValue(value); node != nil {
			return node
		}
	}
	return nil
}
