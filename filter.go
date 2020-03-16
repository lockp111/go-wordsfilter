package wordsfilter

import (
	"regexp"
	"sync"
)

// Filter ...
type Filter struct {
	trie  *Trie
	mux   sync.RWMutex
	noise *regexp.Regexp
}

// New ...
func New(trie *Trie) *Filter {
	if trie == nil {
		trie = NewTrie()
	}

	return &Filter{
		trie:  trie,
		noise: regexp.MustCompile(`[\|\s&%$@*]+`),
	}
}

// UpdateNoisePattern 更新去噪模式
func (f *Filter) UpdateNoisePattern(pattern string) {
	f.mux.Lock()
	defer f.mux.Unlock()

	f.noise = regexp.MustCompile(pattern)
}

// Filter ...
func (f *Filter) Filter(text string) string {
	f.mux.RLock()
	defer f.mux.RUnlock()

	var (
		parent      = f.trie.root
		current     *node
		left        = 0
		runes       = []rune(text)
		length      = len(runes)
		resultRunes = make([]rune, 0, length)
	)

	for position := 0; position < length; position++ {
		current = parent.getChild(runes[position])
		if current == nil || (!current.isPathEnd && position == length-1) {
			resultRunes = append(resultRunes, runes[left])
			parent = f.trie.root
			position = left
			left++
			continue
		}

		if current.isPathEnd {
			left = position + 1
			parent = f.trie.root
		} else {
			parent = current
		}

	}

	resultRunes = append(resultRunes, runes[left:]...)
	return string(resultRunes)
}

// Replace ...
func (f *Filter) Replace(text string, repl rune) string {
	f.mux.RLock()
	defer f.mux.RUnlock()

	var (
		parent  = f.trie.root
		current *node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
	)

	for position := 0; position < len(runes); position++ {
		current = parent.getChild(runes[position])

		if current == nil || (!current.isPathEnd && position == length-1) {
			parent = f.trie.root
			position = left
			left++
			continue
		}

		if current.isPathEnd && left <= position {
			for i := left; i <= position; i++ {
				runes[i] = repl
			}
		}

		parent = current
	}
	return string(runes)
}

// FindIn ...
func (f *Filter) FindIn(text string) (bool, string) {
	f.mux.RLock()
	defer f.mux.RUnlock()

	text = f.removeNoise(text)
	validated, first := f.Validate(text)
	return !validated, first
}

// FindAll ...
func (f *Filter) FindAll(text string) []string {
	f.mux.RLock()
	defer f.mux.RUnlock()

	text = f.removeNoise(text)
	var (
		matches []string
		parent  = f.trie.root
		current *node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
	)

	for position := 0; position < length; position++ {
		current = parent.getChild(runes[position])

		if current == nil {
			parent = f.trie.root
			position = left
			left++
			continue
		}

		if current.isPathEnd && left <= position {
			matches = append(matches, string(runes[left:position+1]))
		}

		if position == length-1 {
			parent = f.trie.root
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

// Validate ...
func (f *Filter) Validate(text string) (bool, string) {
	f.mux.RLock()
	defer f.mux.RUnlock()

	text = f.removeNoise(text)

	var (
		parent  = f.trie.root
		current *node
		runes   = []rune(text)
		length  = len(runes)
		left    = 0
	)

	for position := 0; position < len(runes); position++ {
		current = parent.getChild(runes[position])

		if current == nil || (!current.isPathEnd && position == length-1) {
			parent = f.trie.root
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

// RemoveNoise ...
func (f *Filter) RemoveNoise(text string) string {
	f.mux.RLock()
	defer f.mux.RUnlock()

	return f.noise.ReplaceAllString(text, "")
}

func (f *Filter) removeNoise(text string) string {
	return f.noise.ReplaceAllString(text, "")
}

// UpdateTrie ...
func (f *Filter) UpdateTrie(trie *Trie) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.trie = trie
}
