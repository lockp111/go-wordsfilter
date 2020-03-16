package wordsfilter

import (
	"testing"
)

var tree = NewTrie()

func TestTrieAddwords(t *testing.T) {
	tree.AddWords("垃圾", "测试")
	tree.AddWords("qq", "wechat", "微信")
	tree.AddWords("测试机")
}

func TestTrieShow(t *testing.T) {
	tree.AddWords("垃圾", "测试")
	tree.AddWords("大垃圾", "测试一下")
	tree.Show()
}

func TestFilterReplace(t *testing.T) {
	tree.AddWords("qq", "wechat", "微信")
	filter := New(tree)
	text := filter.Replace("testqq", '*')
	t.Log(text)
}
