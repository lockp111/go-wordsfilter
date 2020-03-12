package wordsfilter

import (
	"testing"
)

func TestTrie(t *testing.T) {
	trieTree := newTrie()
	trieTree.AddWords("test")
}
