package hamt

import (
	"unsafe"

	"github.com/badgerodon/goreify/generics"
)

const w = 5

type (
	// an internalNode is an internal node
	internalNode struct {
		internalBMP uint32
		leafBMP     uint32
		array       [32]unsafe.Pointer
	}
)

func (node *internalNode) getInternal(idx uint32) *internalNode {
	mask := uint32(1 << idx)
	if node.internalBMP&mask == 0 {
		return nil
	}
	return (*internalNode)(node.array[idx])
}
func (node *internalNode) setInternal(idx uint32, n *internalNode) {
	mask := uint32(1 << idx)
	node.internalBMP |= mask
	node.leafBMP ^= mask
	node.array[idx] = unsafe.Pointer(n)
}

func (node *internalNode) getLeaf(idx uint32) *leafNode {
	mask := uint32(1 << idx)
	if node.leafBMP&mask == 0 {
		return nil
	}
	return (*leafNode)(node.array[idx])
}
func (node *internalNode) setLeaf(idx uint32, n *leafNode) {
	mask := uint32(1 << idx)
	node.internalBMP ^= mask
	node.leafBMP |= mask
	node.array[idx] = unsafe.Pointer(n)
}

type (
	// a leafNode is a leaf node
	leafNode struct {
		key   generics.T1
		value generics.T2
	}
)

func (node *leafNode) hash() uint32 {
	return fnv32a(generics.UnsafeBytes(&node.key))
}

type (
	// HashedArrayMappedTrie is a trie with two types of nodes: internal nodes and
	// leaf nodes. Leaf nodes contain a key and value. Internal nodes contain an
	// array of points to internal or leaf nodes and a bitmap to improve lookup
	// performance.
	//
	// Keys are hashed and each level uses W bits from the key. This implementation
	// uses 5 for W so that the number can be represented with a 32 bit integer.
	HashedArrayMappedTrie struct {
		root *internalNode
	}
)

// NewHashedArrayMappedTrie returns a new HashedArrayMappedTrie
func NewHashedArrayMappedTrie() *HashedArrayMappedTrie {
	return &HashedArrayMappedTrie{}
}

// Insert inserts an element into the HashedArrayMappedTrie
func (t *HashedArrayMappedTrie) Insert(key generics.T1, value generics.T2) {
	if t.root == nil {
		t.root = new(internalNode)
	}

	leaf := &leafNode{key: key, value: value}

	t.insert(0, t.root, leaf, leaf.hash())
}

func (t *HashedArrayMappedTrie) insert(
	lvl uint32,
	parent *internalNode,
	child *leafNode,
	hash uint32,
) {
	idx := getIndex(hash, lvl)

	// look to see if we're already an internal node, and if so go down a level
	if cur := parent.getInternal(idx); cur != nil {
		t.insert(lvl+1, cur, child, hash)
		return
	}

	// we have a leaf node
	if cur := parent.getLeaf(idx); cur != nil {
		// if we're equal, replace
		if generics.Equal(cur.key, child.key) {
			cur.value = child.value
			return
		}

		// convert to an internal node
		newNode := new(internalNode)
		t.insert(lvl+1, newNode, cur, cur.hash())
		t.insert(lvl+1, newNode, child, hash)
		parent.setInternal(idx, newNode)
		return
	}

	// no node, create a leaf node
	parent.setLeaf(idx, child)
}
