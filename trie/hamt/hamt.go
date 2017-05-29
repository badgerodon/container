package hamt

import (
	"fmt"
	"strings"
	"unsafe"

	"github.com/badgerodon/goreify/generics"
)

type nodeType byte

const (
	nodeTypeInternal = iota + 1
	nodeTypeLeaf
	nodeTypeLink
)

type node struct {
	Type    nodeType
	Pointer unsafe.Pointer
}

func (node node) String() string {
	switch node.Type {
	case nodeTypeInternal:
		return (*internalNode)(node.Pointer).String()
	case nodeTypeLeaf:
		return (*leafNode)(node.Pointer).String()
	}
	return ""
}

type internalNode struct {
	bitmap word
	array  [64]node
}

func newInternalNode(size word) *internalNode {
	mem := make([]byte, unsafe.Sizeof(word(0))+uintptr(size)*unsafe.Sizeof(uintptr(0)))
	return (*internalNode)(unsafe.Pointer(&mem[0]))
}

func (node *internalNode) String() string {
	var els []string
	for i, idx := range getBitmapIndexes(node.bitmap) {
		els = append(els, fmt.Sprintf("%02d:%s", idx, node.array[i]))
		i++
	}
	return "I{" + strings.Join(els, ",") + "}"
}

func (node *internalNode) delete(idx word) *internalNode {
	mask := word(1 << word(idx))
	// if not set, nothing to do
	if mask&node.bitmap == 0 {
		return node
	}

	size := node.len()

	nn := newInternalNode(size - 1)
	nn.bitmap ^= mask
	pos := 0
	for i := word(0); i < size; i++ {
		if i == idx {
			continue
		}
		nn.array[pos] = node.array[i]
		pos++
	}
	return nn
}

func (node *internalNode) get(idx word) node {
	return node.array[getArrayIndex(node.bitmap, idx)]
}

func (node *internalNode) len() word {
	return popcount(node.bitmap)
}

func (node *internalNode) set(idx word, child node) *internalNode {
	pos := getArrayIndex(node.bitmap, idx)
	fmt.Println("INDEX=", idx, "POSITION=", pos)
	size := node.len()
	mask := word(1 << idx)

	// need to add a new element to the array
	if node.bitmap&mask == 0 {
		newNode := newInternalNode(size + 1)
		newNode.bitmap = node.bitmap
		copy(newNode.array[:pos], node.array[:pos])
		copy(newNode.array[pos+1:], node.array[pos:size])
		node = newNode
	}

	fmt.Println("NODE!", node)

	node.bitmap |= mask
	node.array[pos] = child

	return node
}

type leafNode struct {
	key generics.T1
	val generics.T2
}

func (node *leafNode) Hash() word {
	return word(fnv32a(generics.UnsafeBytes(&node.key)))
}

func (node *leafNode) String() string {
	return fmt.Sprintf("L{%v:%v}", node.key, node.val)
}

type linkNode struct {
	leafNode
	next *linkNode
}

func (node *linkNode) add(leaf *leafNode) (isNew bool) {
	for {
		if generics.Equal(leaf.key, node.key) {
			node.leafNode = *leaf
			return false
		}
		if node.next == nil {
			break
		}
		node = node.next
	}
	node.next = &linkNode{
		leafNode: *leaf,
	}
	return true
}

type HashArrayMappedTrie struct {
	w    word
	root *internalNode
}

func NewHashArrayMappedTrie() *HashArrayMappedTrie {
	return &HashArrayMappedTrie{
		w: 6,
	}
}

func (trie *HashArrayMappedTrie) Get(key generics.T1) (val generics.T2, ok bool) {
	if trie.root == nil {
		return val, false
	}
	hash := (&leafNode{key: key}).Hash()
	return trie.get(0, trie.root, key, hash)
}

func (trie *HashArrayMappedTrie) Set(key generics.T1, val generics.T2) (isNew bool) {
	if trie.root == nil {
		trie.root = newInternalNode(0)
	}
	lnode := &leafNode{
		key: key,
		val: val,
	}
	hash := lnode.Hash()
	trie.root, isNew = trie.set(0, trie.root, lnode, hash)
	return isNew
}

func (trie *HashArrayMappedTrie) get(lvl word, parent *internalNode, key generics.T1, hash word) (val generics.T2, ok bool) {
	idx := getHashIndex(hash, trie.w, lvl)
	child := parent.get(idx)
	switch child.Type {
	case nodeTypeInternal:
		inode := (*internalNode)(child.Pointer)
		return trie.get(lvl+1, inode, key, hash)
	case nodeTypeLeaf:
		lnode := (*leafNode)(child.Pointer)
		if generics.Equal(key, lnode.key) {
			return lnode.val, true
		}
	case nodeTypeLink:
		llnode := (*linkNode)(child.Pointer)
		for llnode != nil {
			if generics.Equal(key, llnode.key) {
				return llnode.val, true
			}
			llnode = llnode.next
		}
	}
	return val, false
}

func (trie *HashArrayMappedTrie) set(lvl word, parent *internalNode, lnode *leafNode, hash word) (newNode *internalNode, isNew bool) {
	idx := getHashIndex(hash, trie.w, lvl)
	child := parent.get(idx)
	switch child.Type {
	case nodeTypeInternal:
		inode := (*internalNode)(child.Pointer)
		newNode, isNew = trie.set(lvl+1, inode, lnode, hash)
		newNode = parent.set(idx, node{
			Type:    nodeTypeInternal,
			Pointer: unsafe.Pointer(newNode),
		})
		return newNode, isNew
	case nodeTypeLeaf:
		enode := (*leafNode)(child.Pointer)
		if generics.Equal(lnode.key, enode.key) {
			// replace
			parent.set(idx, node{
				Type:    nodeTypeLeaf,
				Pointer: unsafe.Pointer(lnode),
			})
			return parent, false
		} else if lvl == maxLevel(trie.w) {
			// convert to link
			llnode := &linkNode{
				leafNode: *lnode,
			}
			llnode.add(enode)
			parent.set(idx, node{
				Type:    nodeTypeLink,
				Pointer: unsafe.Pointer(llnode),
			})
			return parent, true
		} else {
			// convert to internal
			inode := newInternalNode(0)
			trie.set(lvl+1, inode, lnode, hash)
			trie.set(lvl+1, inode, enode, enode.Hash())
			parent.set(idx, node{
				Type:    nodeTypeInternal,
				Pointer: unsafe.Pointer(inode),
			})
			return parent, true
		}
	case nodeTypeLink:
		llnode := (*linkNode)(child.Pointer)
		return parent, llnode.add(lnode)
	default:
		fmt.Println("ADD LEAF")
		fmt.Println("|", parent)
		// create a leaf node
		newNode := parent.set(idx, node{
			Type:    nodeTypeLeaf,
			Pointer: unsafe.Pointer(lnode),
		})
		fmt.Println("|", newNode)
		return newNode, true
	}
}
