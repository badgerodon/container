package sparse

import "github.com/badgerodon/goreify/generics"

//go:generate go run gen/sparse_nodes.go

// An Array is a compressed sparse array using a bitmap as an index
// on the underlying data array
type Array struct {
	size int
	node Node
}

// NewArray creates a new array with up to `size` items
func NewArray(size int) *Array {
	var node Node
	switch {
	case size == 0:
	case size <= 8:
		node = new(Node8x0)
	case size <= 16:
		node = new(Node16x0)
	case size <= 32:
		node = new(Node32x0)
	case size <= 64:
		node = new(Node64x0)
	default:
		panic("not implemented")
	}

	return &Array{
		size: size,
		node: node,
	}
}

// Get returns an element in the array
func (arr *Array) Get(idx int) (val generics.T1, ok bool) {
	if idx >= arr.size {
		panic("index out of range")
	}
	return arr.node.Get(uint(idx))
}

// Delete delets an element from the array
func (arr *Array) Delete(idx int) {
	if idx >= arr.size {
		panic("index out of range")
	}
	arr.node, _ = arr.node.Delete(uint(idx))
}

// Len returns the number of elements in the sparse array
func (arr *Array) Len() int {
	return int(arr.node.Len())
}

// Set sets an element in the array
func (arr *Array) Set(idx int, val generics.T1) {
	if idx >= arr.size {
		panic("index out of range")
	}
	arr.node, _ = arr.node.Set(uint(idx), val)
}

// Slice returns the elements of the sparse array as a slice.
// This is a pointer to the actual data.
func (arr *Array) Slice() []generics.T1 {
	return arr.node.Slice()
}
