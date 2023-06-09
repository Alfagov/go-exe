package pov

func (tr *Tree) Value() string {
	return tr.value
}

func (tr *Tree) Children() []*Tree {
	return tr.children
}

func (tr *Tree) hasChild(node *Tree) bool {
	for _, child := range tr.children {
		if child == node {
			return true
		}
	}
	return false
}

// Converts tree to string representation
func (tr *Tree) String() string {
    if tr == nil {
        return "nil"
    }
    result := tr.Value()
    if len(tr.Children()) == 0 {
        return result
    }
    for _, ch := range tr.Children() {
        result += " " + ch.String()
    }
    return "(" + result + ")"
}