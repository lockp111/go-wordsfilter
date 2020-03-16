package wordsfilter

import (
	"fmt"
	"sync"
)

type node struct {
	mux       sync.RWMutex
	children  map[rune]*node
	character rune
	isRoot    bool
	isPathEnd bool
	level     int
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

	child.level = n.level + 1
	n.children[child.character] = child
}

func (n *node) getChild(character rune) *node {
	n.mux.RLock()
	defer n.mux.RUnlock()

	return n.children[character]
}

func (n *node) getChilds() []*node {
	n.mux.RLock()
	defer n.mux.RUnlock()

	childs := make([]*node, 0, len(n.children))
	for _, c := range n.children {
		childs = append(childs, c)
	}
	return childs
}

func (n *node) show() {
	for i := 0; i < n.level-1; i++ {
		fmt.Print("|  ")
	}
	fmt.Print("|--")
	fmt.Println(string(n.character))
	for _, v := range n.getChilds() {
		v.show()
	}
}
