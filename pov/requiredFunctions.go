package pov 

// Method FromPov takes a string argument from which specifies a node in the tree via its value. 
// It should return a tree with the value from in the root.
func (tr *Tree) FromPov(from string) *Tree {
	if node := tr.findNodeByValue(from); node != nil {
		return node.fP(nil)
	}
	return nil
}

// Method PathTo takes two string arguments from and to which specify two nodes in the tree via their values.
// It should return the shortest path in the tree from the first to the second node.
func (tr *Tree) PathTo(from, to string) []string {
	startNode := tr.findNodeByValue(from)
	endNode := tr.findNodeByValue(to)
	if startNode == nil || endNode == nil {
		return []string{}
	}
	if path, ok := startNode.pT(endNode, []string{}); ok {
		return path
	}
	return []string{}
}