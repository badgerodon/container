package hamt

import "testing"

func Test(t *testing.T) {
	trie := NewHashedArrayMappedTrie()
	trie.Insert(5, 5)
	trie.Insert(137, 5)
	t.FailNow()
}
