package wordsfilter

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Trie keyword tree
type Trie struct {
	root *node
}

// NewTrie init trie
func NewTrie() *Trie {
	return &Trie{
		root: &node{
			isRoot:   true,
			children: make(map[rune]*node),
		},
	}
}

// AddWords add keywords
func (t *Trie) AddWords(words ...string) {
	for _, w := range words {
		t.add(w)
	}
}

// Load common method to add words
func (t *Trie) Load(rd io.Reader) error {
	buf := bufio.NewReader(rd)
	for {
		line, _, err := buf.ReadLine()
		if err != nil {
			if err != io.EOF {
				return err
			}
			break
		}
		t.add(string(line))
	}

	return nil
}

// LoadWordDict load local dict
func (t *Trie) LoadWordDict(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return t.Load(file)
}

// LoadNetWordDict load net dict
func (t *Trie) LoadNetWordDict(url string) error {
	c := http.Client{
		Timeout: 5 * time.Second,
	}
	rsp, err := c.Get(url)
	if err != nil {
		return err
	}
	defer rsp.Body.Close()

	return t.Load(rsp.Body)
}

// Show print Trie struct
func (t *Trie) Show() {
	fmt.Println("root")
	fmt.Println("|")
	for _, n := range t.root.getChilds() {
		n.show()
	}
}

func (t *Trie) add(word string) {
	var (
		current = t.root
		runes   = []rune(word)
	)

	for _, v := range runes {
		if next := current.getChild(v); next != nil {
			current = next
			continue
		}

		newNode := newChildren(v)
		current.addChild(newNode)
		current = newNode
	}

	current.isPathEnd = true
}
