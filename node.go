package wordsfilter

import (
	"sync"
)

type node struct {
	mux       sync.RWMutex
	children  map[rune]*node
	character rune
	isRoot    bool
	isPathEnd bool
}

func newChildren(character rune) *node {
	return &node{
		character: character,
		children:  make(map[rune]*node),
	}
}

// check node is leaf node
func (n *node) isLeaf() bool {
	n.mux.RLock()
	defer n.mux.RUnlock()

	return len(n.children) == 0
}

func (n *node) addChild(child *node) {
	n.mux.Lock()
	defer n.mux.Unlock()

	n.children[child.character] = child
}

func (n *node) removeChild(child *node) {
	n.mux.Lock()
	defer n.mux.Unlock()

	delete(n.children, child.character)
}

func (n *node) getChild(character rune) *node {
	n.mux.RLock()
	defer n.mux.RUnlock()

	return n.children[character]
}
