package heap

import "github.com/badgerodon/goreify/generics"

//go:generate goreify github.com/badgerodon/container/heap.BinaryHeap,binaryHeapNode numeric,string

// A BinaryHeap implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
type BinaryHeap struct {
	less     func(x, y generics.T1) bool
	elements []generics.T1
}

// NewBinaryHeap creates a new BinaryHeap
func NewBinaryHeap(less func(x, y generics.T1) bool) *BinaryHeap {
	return &BinaryHeap{
		less: less,
	}
}

// Peek returns the minimum element in the heap
func (h *BinaryHeap) Peek() (el generics.T1, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Pop removes the minumum element from the heap
func (h *BinaryHeap) Pop() (el generics.T1, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	// save the root element
	el = h.elements[0]

	// replace the root with the last element in the heap
	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)
		// no children
		if l >= len(h.elements) {
			break
		}

		min := l
		// if there's a right element, and it's smaller than the left element, use the right
		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {
			// left element is less than i, so swap
			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {
			// we're in order, so all done
			break
		}
	}

	return el, true
}

// Push pushes the element onto the heap
func (h *BinaryHeap) Push(el generics.T1) {
	i := len(h.elements)

	// add the element to the list
	h.elements = append(h.elements, el)

	// move element up the tree until it's not less than it's parent
	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

func (h *BinaryHeap) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeap) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeap) rightIndex(i int) int {
	return (2 * i) + 2
}
