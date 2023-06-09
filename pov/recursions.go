package pov

// Recursion function for FromPov for children nodes
func (node *Tree) fP(prev *Tree) *Tree {
	NewTreeChildren := []*Tree{}
	for _, conNode := range node.connectedNodes() {
		if conNode != prev {
			NewTreeChildren = append(NewTreeChildren, conNode.fP(node))
		}
	}
	return NewTree(node.value, NewTreeChildren...)
}

// Recursion function for PathTo for end node
func (node *Tree) pT(endNode *Tree, path []string) ([]string, bool) {
	if node == endNode {
		return append(path, endNode.value), true
	}
	for _, conNode := range node.connectedNodes() {
		if len(path) > 0 && path[len(path)-1] == conNode.value {
			continue
		}
		if p, ok := conNode.pT(endNode, append(path, node.value)); ok {
			return p, true
		}
	}
	return []string{}, false
}
