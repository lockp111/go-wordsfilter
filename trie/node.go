package trie

// trie tree node
type node struct {
	children  map[rune]*node
	character rune
	isRoot    bool
	isPathEnd bool
}

func newNode(character rune) *node {
	return &node{
		character: character,
		children:  make(map[rune]*node),
	}
}

func newRoot() *node {
	return &node{
		isRoot:   true,
		children: make(map[rune]*node),
	}
}

func (n *node) isLeafNode() bool {
	return len(n.children) == 0
}
