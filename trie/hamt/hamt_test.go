package hamt

import (
	"fmt"
	"reflect"
	"testing"
)

// func TestTest(t *testing.T) {
// 	t.Logf("%064b", maxword<<4)
// 	t.Fail()
// }

func Test(t *testing.T) {
	trie := NewHashArrayMappedTrie()

	_, ok := trie.Get(5)
	if ok {
		t.Errorf("expected ok to be false, since nothing has been added")
	}

	isNew := trie.Set(5, 6)
	if !isNew {
		t.Errorf("expected setting a previously unset key to return true")
	}
	v, ok := trie.Get(5)
	if !ok {
		t.Errorf("expected ok to be true")
	}
	if !reflect.DeepEqual(v, 6) {
		t.Errorf("expected value to be 6, got %v", v)
	}

	trie.Set(5, 7)
	v, ok = trie.Get(5)
	if !ok {
		t.Errorf("expected ok to be true")
	}
	if !reflect.DeepEqual(v, 7) {
		t.Errorf("expected value to be 7, got %v", v)
	}

	// trie.Delete(5)
	// _, ok = trie.Get(5)
	// if ok {
	// 	t.Errorf("expected ok to be false, since nothing has been added")
	// }

	fmt.Println(trie.root)
	t.FailNow()
}
