package hamt

// import (
// 	"fmt"
// 	"unsafe"

// 	"github.com/badgerodon/goreify/generics"
// )

// const w = 5

// type (
// 	// an internalNode is an internal node
// 	internalNode struct {
// 		internalBMP uint32
// 		leafBMP     uint32
// 		array       [32]unsafe.Pointer
// 	}
// )

// func (node *internalNode) delete(idx uint32) {
// 	mask := uint32(1 << idx)
// 	node.internalBMP &= ^mask
// 	node.leafBMP &= ^mask
// 	node.array[idx] = nil
// }

// func (node *internalNode) getInternal(idx uint32) *internalNode {
// 	mask := uint32(1 << idx)
// 	if node.internalBMP&mask == 0 {
// 		return nil
// 	}
// 	return (*internalNode)(node.array[idx])
// }

// func (node *internalNode) getLeaf(idx uint32) *leafNode {
// 	mask := uint32(1 << idx)
// 	if node.leafBMP&mask == 0 {
// 		return nil
// 	}
// 	return (*leafNode)(node.array[idx])
// }

// func (node *internalNode) len() int {
// 	return popcountHD(node.internalBMP) + popcountHD(node.leafBMP)
// }

// func (node *internalNode) setInternal(idx uint32, n *internalNode) {
// 	mask := uint32(1 << idx)
// 	node.internalBMP |= mask
// 	node.leafBMP &= ^mask
// 	node.array[idx] = unsafe.Pointer(n)
// }

// func (node *internalNode) setLeaf(idx uint32, n *leafNode) {
// 	mask := uint32(1 << idx)
// 	node.internalBMP &= ^mask
// 	node.leafBMP |= mask
// 	node.array[idx] = unsafe.Pointer(n)
// }

// func (node *internalNode) String() string {
// 	return fmt.Sprintf("InternalNode(%032b,%032b,%v)", node.internalBMP, node.leafBMP, node.array)
// }

// type (
// 	// a leafNode is a leaf node
// 	leafNode struct {
// 		key   generics.T1
// 		value generics.T2
// 	}
// )

// func (node *leafNode) hash() uint32 {
// 	return fnv32a(generics.UnsafeBytes(&node.key))
// }

// func (node *leafNode) String() string {
// 	return fmt.Sprintf("Leaf(%v,%v)", node.key, node.value)
// }

// type (
// 	// HashedArrayMappedTrie is a trie with two types of nodes: internal nodes and
// 	// leaf nodes. Leaf nodes contain a key and value. Internal nodes contain an
// 	// array of points to internal or leaf nodes and a bitmap to improve lookup
// 	// performance.
// 	//
// 	// Keys are hashed and each level uses W bits from the key. This implementation
// 	// uses 5 for W so that the number can be represented with a 32 bit integer.
// 	HashedArrayMappedTrie struct {
// 		root *internalNode
// 	}
// )

// // NewHashedArrayMappedTrie returns a new HashedArrayMappedTrie
// func NewHashedArrayMappedTrie() *HashedArrayMappedTrie {
// 	return &HashedArrayMappedTrie{}
// }

// // Delete removes the key from the trie. It returns true if the key was found.
// func (t *HashedArrayMappedTrie) Delete(key generics.T1) (found bool) {
// 	if t.root == nil {
// 		return false
// 	}
// 	leaf := &leafNode{key: key}
// 	return t.delete(0, t.root, leaf.key, leaf.hash())
// }

// // Get gets a value from the trie.
// func (t *HashedArrayMappedTrie) Get(key generics.T1) (value generics.T2, ok bool) {
// 	if t.root == nil {
// 		return value, false
// 	}
// 	leaf := &leafNode{key: key}
// 	return t.get(0, t.root, leaf.key, leaf.hash())
// }

// // Set sets the key to the value. If it already exists its replaced.
// func (t *HashedArrayMappedTrie) Set(key generics.T1, value generics.T2) {
// 	if t.root == nil {
// 		t.root = new(internalNode)
// 	}
// 	leaf := &leafNode{key: key, value: value}
// 	t.set(0, t.root, leaf, leaf.hash())
// }

// func (t *HashedArrayMappedTrie) delete(
// 	lvl uint32,
// 	parent *internalNode,
// 	key generics.T1,
// 	hash uint32,
// ) (found bool) {
// 	idx := getIndex(hash, lvl)

// 	// if this is an internal node, traverse down
// 	if cur := parent.getInternal(idx); cur != nil {
// 		found = t.delete(lvl+1, cur, key, hash)
// 		if found && cur.len() == 0 {
// 			parent.delete(idx)
// 		}
// 		return found
// 	}

// 	// if this is a leaf node
// 	if cur := parent.getLeaf(idx); cur != nil {
// 		// only return if the keys match
// 		if generics.Equal(cur.key, key) {
// 			parent.delete(idx)
// 			return true
// 		}
// 	}

// 	return false
// }

// func (t *HashedArrayMappedTrie) get(
// 	lvl uint32,
// 	parent *internalNode,
// 	key generics.T1,
// 	hash uint32,
// ) (value generics.T2, ok bool) {
// 	idx := getIndex(hash, lvl)

// 	// if this is an internal node, traverse down
// 	if cur := parent.getInternal(idx); cur != nil {
// 		return t.get(lvl+1, cur, key, hash)
// 	}

// 	// if this is a leaf node
// 	if cur := parent.getLeaf(idx); cur != nil {
// 		// only return if the keys match
// 		if generics.Equal(cur.key, key) {
// 			return cur.value, true
// 		}
// 	}

// 	// guess it doesn't exist
// 	return value, false
// }

// func (t *HashedArrayMappedTrie) set(
// 	lvl uint32,
// 	parent *internalNode,
// 	child *leafNode,
// 	hash uint32,
// ) {
// 	idx := getIndex(hash, lvl)

// 	// look to see if we're already an internal node, and if so go down a level
// 	if cur := parent.getInternal(idx); cur != nil {
// 		t.set(lvl+1, cur, child, hash)
// 		return
// 	}

// 	// we have a leaf node
// 	if cur := parent.getLeaf(idx); cur != nil {
// 		// if we're equal, replace
// 		if generics.Equal(cur.key, child.key) {
// 			cur.value = child.value
// 			return
// 		}

// 		// convert to an internal node
// 		newNode := new(internalNode)
// 		t.set(lvl+1, newNode, cur, cur.hash())
// 		t.set(lvl+1, newNode, child, hash)
// 		parent.setInternal(idx, newNode)
// 		return
// 	}

// 	// no node, create a leaf node
// 	parent.setLeaf(idx, child)
// }
