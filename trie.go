package wordsfilter

type trie struct {
	root *node
}

func newTrie() *trie {
	return &trie{
		root: &node{
			isRoot:   true,
			children: make(map[rune]*node),
		},
	}
}

// AddWords add words
func (t *trie) AddWords(words ...string) {
	for _, word := range words {
		t.add(word)
	}
}

func (t *trie) add(word string) {
	var (
		current = t.root
		runes   = []rune(word)
	)

	for _, v := range runes {
		if next := current.getChild(v); next != nil {
			current = next
		}

		newNode := newChildren(v)
		current.addChild(newNode)
		current = newNode
	}

	current.isPathEnd = true
}

// DelWords delete words
func (t *trie) DelWords(words ...string) {
	for _, word := range words {
		t.del(word)
	}
}

func (t *trie) del(word string) {
	var (
		current  = t.root
		runes    = []rune(word)
		nodeList []*node
	)

	for _, v := range runes {
		next := current.getChild(v)
		if next == nil {
			return
		}

		list := make([]*node, 0, len(nodeList)+1)
		nodeList = append(append(list, next), nodeList...)
		current = next
	}

	for k, v := range nodeList {
		if k != 0 && v.isPathEnd {
			return
		} else if k == 0 && !v.isPathEnd {
			return
		}

		nodeList[k+1].removeChild(v)
	}
}

func (t *trie) replace(text string, character rune) string {
	var (
		parent  = t.root
		current *node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.children[runes[position]]

		if !found || (!current.isPathEnd && position == length-1) {
			parent = t.root
			position = left
			left++
			continue
		}

		if current.isPathEnd && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = character
			}
		}

		parent = current
	}

	return string(runes)
}

func (t *trie) filter(text string) string {
	var (
		parent      = t.root
		current     *node
		left        = 0
		found       bool
		runes       = []rune(text)
		length      = len(runes)
		resultRunes = make([]rune, 0, length)
	)

	for position := 0; position < length; position++ {
		current, found = parent.children[runes[position]]

		if !found || (!current.isPathEnd && position == length-1) {
			resultRunes = append(resultRunes, runes[left])
			parent = t.root
			position = left
			left++
			continue
		}

		if current.isPathEnd {
			left = position + 1
			parent = t.root
		} else {
			parent = current
		}

	}

	resultRunes = append(resultRunes, runes[left:]...)
	return string(resultRunes)
}

func (t *trie) validate(text string) (bool, string) {
	var (
		parent  = t.root
		current *node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < len(runes); position++ {
		current, found = parent.children[runes[position]]

		if !found || (!current.isPathEnd && position == length-1) {
			parent = t.root
			position = left
			left++
			continue
		}

		if current.isPathEnd && left <= position {
			return false, string(runes[left : position+1])
		}

		parent = current
	}

	return true, ""
}

func (t *trie) findIn(text string) (bool, string) {
	validated, first := t.validate(text)
	return !validated, first
}

func (t *trie) findAll(text string) []string {
	var matches []string
	var (
		parent  = t.root
		current *node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
		found   bool
	)

	for position := 0; position < length; position++ {
		current, found = parent.children[runes[position]]

		if !found {
			parent = t.root
			position = left
			left++
			continue
		}

		if current.isPathEnd && left <= position {
			matches = append(matches, string(runes[left:position+1]))
		}

		if position == length-1 {
			parent = t.root
			position = left
			left++
			continue
		}

		parent = current
	}

	var i = 0
	if count := len(matches); count > 0 {
		set := make(map[string]struct{})
		for i < count {
			_, ok := set[matches[i]]
			if !ok {
				set[matches[i]] = struct{}{}
				i++
				continue
			}
			count--
			copy(matches[i:], matches[i+1:])
		}
		return matches[:count]
	}

	return nil
}
