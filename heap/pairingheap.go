package heap

import "github.com/badgerodon/goreify/generics"

//go:generate goreify github.com/badgerodon/container/heap.PairingHeap,pairingHeapNode numeric,string

// A PairingHeap implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
type PairingHeap struct {
	less func(generics.T1, generics.T1) bool
	root *pairingHeapNode
}

type pairingHeapNode struct {
	elem                   generics.T1
	parent, child, sibling *pairingHeapNode
}

// NewPairingHeap creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeap(less func(generics.T1, generics.T1) bool) *PairingHeap {
	return &PairingHeap{
		less: less,
	}
}

// Peek returns the minimum element in the heap
func (h *PairingHeap) Peek() (el generics.T1, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Pop returns the minimum element in the heap and removes it
func (h *PairingHeap) Pop() (el generics.T1, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Push adds the element to the heap
func (h *PairingHeap) Push(el generics.T1) {
	h.root = h.merge(h.root, &pairingHeapNode{
		elem: el,
	})
}

// Merge merges two heaps
func (h *PairingHeap) Merge(h2 *PairingHeap) *PairingHeap {
	return &PairingHeap{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

func (h *PairingHeap) merge(n1, n2 *pairingHeapNode) *pairingHeapNode {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeap) mergePairs(n *pairingHeapNode) *pairingHeapNode {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}
