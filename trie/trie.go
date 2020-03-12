package trie

// Trie tree
type Trie struct {
	root *node
}

// NewTrie ...
func NewTrie() *Trie {
	return &Trie{
		root: newRoot(),
	}
}

// Add add words
func (t *Trie) Add(words ...string) {
	for _, word := range words {
		t.add(word)
	}
}

func (t *Trie) add(word string) {
	var current = t.root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.children[r]; ok {
			current = next
		} else {
			newNode := newNode(r)
			current.children[r] = newNode
			current = newNode
		}
		if position == len(runes)-1 {
			current.isPathEnd = true
		}
	}
}

// Del delete words
func (t *Trie) Del(words ...string) {
	for _, word := range words {
		t.del(word)
	}
}

func (t *Trie) del(word string) {
	var current = t.root
	var runes = []rune(word)
	for position := 0; position < len(runes); position++ {
		r := runes[position]
		if next, ok := current.children[r]; !ok {
			return
		} else {
			current = next
		}

		/* if position == len(runes)-1 {
			current.SoftDel()
		} */
	}
}

// Replace text replace
func (t *Trie) Replace(text string, character rune) string {
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

// Filter ...
func (t *Trie) Filter(text string) string {
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

// Validate ...
func (t *Trie) Validate(text string) (bool, string) {
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

// FindIn ...
func (t *Trie) FindIn(text string) (bool, string) {
	validated, first := t.Validate(text)
	return !validated, first
}

// FindAll ...
func (t *Trie) FindAll(text string) []string {
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
