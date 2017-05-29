package sparse

import "github.com/badgerodon/goreify/generics"

// A Node is a sparse array
type Node interface {
	Contains(idx uint) bool
	Delete(idx uint) (Node, bool)
	Get(idx uint) (generics.T1, bool)
	Len() uint
	Set(idx uint, val generics.T1) (Node, bool)
	Slice() []generics.T1
}


// A Node8x0 is a sparse array node
type Node8x0 struct {
	Bitmap uint8
	Array [0]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x0) Add(idx uint, val generics.T1) (newNode *Node8x1, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x1)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x0) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x0) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node, false

}

// Get returns the element at the given index
func (node *Node8x0) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x0) Len() uint {
	return 0
}

// Replace replaces the value
func (node *Node8x0) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x0) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x0) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x1 is a sparse array node
type Node8x1 struct {
	Bitmap uint8
	Array [1]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x1) Add(idx uint, val generics.T1) (newNode *Node8x2, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x2)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x1) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x1) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x1) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x1) Len() uint {
	return 1
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x1) Remove(idx uint) (newNode *Node8x0, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x0)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x1) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x1) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x1) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x2 is a sparse array node
type Node8x2 struct {
	Bitmap uint8
	Array [2]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x2) Add(idx uint, val generics.T1) (newNode *Node8x3, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x3)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x2) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x2) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x2) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x2) Len() uint {
	return 2
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x2) Remove(idx uint) (newNode *Node8x1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x1)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x2) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x2) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x2) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x3 is a sparse array node
type Node8x3 struct {
	Bitmap uint8
	Array [3]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x3) Add(idx uint, val generics.T1) (newNode *Node8x4, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x4)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x3) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x3) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x3) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x3) Len() uint {
	return 3
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x3) Remove(idx uint) (newNode *Node8x2, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x2)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x3) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x3) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x3) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x4 is a sparse array node
type Node8x4 struct {
	Bitmap uint8
	Array [4]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x4) Add(idx uint, val generics.T1) (newNode *Node8x5, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x5)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x4) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x4) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x4) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x4) Len() uint {
	return 4
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x4) Remove(idx uint) (newNode *Node8x3, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x3)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x4) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x4) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x4) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x5 is a sparse array node
type Node8x5 struct {
	Bitmap uint8
	Array [5]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x5) Add(idx uint, val generics.T1) (newNode *Node8x6, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x6)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x5) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x5) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x5) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x5) Len() uint {
	return 5
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x5) Remove(idx uint) (newNode *Node8x4, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x4)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x5) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x5) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x5) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x6 is a sparse array node
type Node8x6 struct {
	Bitmap uint8
	Array [6]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x6) Add(idx uint, val generics.T1) (newNode *Node8x7, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x7)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x6) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x6) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x6) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x6) Len() uint {
	return 6
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x6) Remove(idx uint) (newNode *Node8x5, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x5)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x6) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x6) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x6) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x7 is a sparse array node
type Node8x7 struct {
	Bitmap uint8
	Array [7]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node8x7) Add(idx uint, val generics.T1) (newNode *Node8x8, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node8x8)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition8(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x7) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x7) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x7) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x7) Len() uint {
	return 7
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x7) Remove(idx uint) (newNode *Node8x6, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x6)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x7) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x7) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node8x7) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node8x8 is a sparse array node
type Node8x8 struct {
	Bitmap uint8
	Array [8]generics.T1
}

// Contains returns whether or not there's an element at the idx
func (node *Node8x8) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node8x8) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node8x8) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node8x8) Len() uint {
	return 8
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node8x8) Remove(idx uint) (newNode *Node8x7, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	newNode = new(Node8x7)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node8x8) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition8(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node8x8) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node, false

}

func (node *Node8x8) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x0 is a sparse array node
type Node16x0 struct {
	Bitmap uint16
	Array [0]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x0) Add(idx uint, val generics.T1) (newNode *Node16x1, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x1)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x0) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x0) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node, false

}

// Get returns the element at the given index
func (node *Node16x0) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x0) Len() uint {
	return 0
}

// Replace replaces the value
func (node *Node16x0) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x0) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x0) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x1 is a sparse array node
type Node16x1 struct {
	Bitmap uint16
	Array [1]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x1) Add(idx uint, val generics.T1) (newNode *Node16x2, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x2)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x1) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x1) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x1) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x1) Len() uint {
	return 1
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x1) Remove(idx uint) (newNode *Node16x0, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x0)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x1) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x1) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x1) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x2 is a sparse array node
type Node16x2 struct {
	Bitmap uint16
	Array [2]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x2) Add(idx uint, val generics.T1) (newNode *Node16x3, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x3)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x2) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x2) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x2) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x2) Len() uint {
	return 2
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x2) Remove(idx uint) (newNode *Node16x1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x1)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x2) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x2) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x2) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x3 is a sparse array node
type Node16x3 struct {
	Bitmap uint16
	Array [3]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x3) Add(idx uint, val generics.T1) (newNode *Node16x4, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x4)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x3) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x3) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x3) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x3) Len() uint {
	return 3
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x3) Remove(idx uint) (newNode *Node16x2, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x2)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x3) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x3) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x3) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x4 is a sparse array node
type Node16x4 struct {
	Bitmap uint16
	Array [4]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x4) Add(idx uint, val generics.T1) (newNode *Node16x5, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x5)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x4) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x4) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x4) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x4) Len() uint {
	return 4
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x4) Remove(idx uint) (newNode *Node16x3, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x3)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x4) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x4) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x4) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x5 is a sparse array node
type Node16x5 struct {
	Bitmap uint16
	Array [5]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x5) Add(idx uint, val generics.T1) (newNode *Node16x6, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x6)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x5) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x5) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x5) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x5) Len() uint {
	return 5
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x5) Remove(idx uint) (newNode *Node16x4, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x4)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x5) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x5) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x5) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x6 is a sparse array node
type Node16x6 struct {
	Bitmap uint16
	Array [6]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x6) Add(idx uint, val generics.T1) (newNode *Node16x7, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x7)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x6) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x6) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x6) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x6) Len() uint {
	return 6
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x6) Remove(idx uint) (newNode *Node16x5, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x5)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x6) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x6) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x6) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x7 is a sparse array node
type Node16x7 struct {
	Bitmap uint16
	Array [7]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x7) Add(idx uint, val generics.T1) (newNode *Node16x8, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x8)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x7) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x7) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x7) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x7) Len() uint {
	return 7
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x7) Remove(idx uint) (newNode *Node16x6, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x6)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x7) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x7) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x7) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x8 is a sparse array node
type Node16x8 struct {
	Bitmap uint16
	Array [8]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x8) Add(idx uint, val generics.T1) (newNode *Node16x9, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x9)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x8) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x8) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x8) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x8) Len() uint {
	return 8
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x8) Remove(idx uint) (newNode *Node16x7, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x7)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x8) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x8) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x8) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x9 is a sparse array node
type Node16x9 struct {
	Bitmap uint16
	Array [9]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x9) Add(idx uint, val generics.T1) (newNode *Node16x10, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x10)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x9) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x9) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x9) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x9) Len() uint {
	return 9
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x9) Remove(idx uint) (newNode *Node16x8, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x8)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x9) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x9) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x9) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x10 is a sparse array node
type Node16x10 struct {
	Bitmap uint16
	Array [10]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x10) Add(idx uint, val generics.T1) (newNode *Node16x11, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x11)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x10) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x10) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x10) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x10) Len() uint {
	return 10
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x10) Remove(idx uint) (newNode *Node16x9, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x9)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x10) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x10) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x10) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x11 is a sparse array node
type Node16x11 struct {
	Bitmap uint16
	Array [11]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x11) Add(idx uint, val generics.T1) (newNode *Node16x12, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x12)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x11) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x11) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x11) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x11) Len() uint {
	return 11
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x11) Remove(idx uint) (newNode *Node16x10, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x10)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x11) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x11) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x11) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x12 is a sparse array node
type Node16x12 struct {
	Bitmap uint16
	Array [12]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x12) Add(idx uint, val generics.T1) (newNode *Node16x13, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x13)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x12) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x12) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x12) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x12) Len() uint {
	return 12
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x12) Remove(idx uint) (newNode *Node16x11, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x11)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x12) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x12) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x12) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x13 is a sparse array node
type Node16x13 struct {
	Bitmap uint16
	Array [13]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x13) Add(idx uint, val generics.T1) (newNode *Node16x14, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x14)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x13) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x13) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x13) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x13) Len() uint {
	return 13
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x13) Remove(idx uint) (newNode *Node16x12, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x12)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x13) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x13) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x13) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x14 is a sparse array node
type Node16x14 struct {
	Bitmap uint16
	Array [14]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x14) Add(idx uint, val generics.T1) (newNode *Node16x15, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x15)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x14) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x14) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x14) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x14) Len() uint {
	return 14
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x14) Remove(idx uint) (newNode *Node16x13, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x13)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x14) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x14) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x14) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x15 is a sparse array node
type Node16x15 struct {
	Bitmap uint16
	Array [15]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node16x15) Add(idx uint, val generics.T1) (newNode *Node16x16, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node16x16)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition16(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x15) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x15) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x15) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x15) Len() uint {
	return 15
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x15) Remove(idx uint) (newNode *Node16x14, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x14)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x15) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x15) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node16x15) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node16x16 is a sparse array node
type Node16x16 struct {
	Bitmap uint16
	Array [16]generics.T1
}

// Contains returns whether or not there's an element at the idx
func (node *Node16x16) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node16x16) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node16x16) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node16x16) Len() uint {
	return 16
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node16x16) Remove(idx uint) (newNode *Node16x15, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	newNode = new(Node16x15)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node16x16) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition16(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node16x16) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node, false

}

func (node *Node16x16) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x0 is a sparse array node
type Node32x0 struct {
	Bitmap uint32
	Array [0]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x0) Add(idx uint, val generics.T1) (newNode *Node32x1, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x1)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x0) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x0) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node, false

}

// Get returns the element at the given index
func (node *Node32x0) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x0) Len() uint {
	return 0
}

// Replace replaces the value
func (node *Node32x0) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x0) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x0) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x1 is a sparse array node
type Node32x1 struct {
	Bitmap uint32
	Array [1]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x1) Add(idx uint, val generics.T1) (newNode *Node32x2, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x2)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x1) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x1) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x1) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x1) Len() uint {
	return 1
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x1) Remove(idx uint) (newNode *Node32x0, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x0)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x1) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x1) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x1) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x2 is a sparse array node
type Node32x2 struct {
	Bitmap uint32
	Array [2]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x2) Add(idx uint, val generics.T1) (newNode *Node32x3, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x3)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x2) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x2) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x2) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x2) Len() uint {
	return 2
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x2) Remove(idx uint) (newNode *Node32x1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x1)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x2) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x2) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x2) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x3 is a sparse array node
type Node32x3 struct {
	Bitmap uint32
	Array [3]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x3) Add(idx uint, val generics.T1) (newNode *Node32x4, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x4)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x3) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x3) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x3) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x3) Len() uint {
	return 3
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x3) Remove(idx uint) (newNode *Node32x2, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x2)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x3) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x3) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x3) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x4 is a sparse array node
type Node32x4 struct {
	Bitmap uint32
	Array [4]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x4) Add(idx uint, val generics.T1) (newNode *Node32x5, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x5)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x4) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x4) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x4) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x4) Len() uint {
	return 4
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x4) Remove(idx uint) (newNode *Node32x3, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x3)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x4) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x4) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x4) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x5 is a sparse array node
type Node32x5 struct {
	Bitmap uint32
	Array [5]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x5) Add(idx uint, val generics.T1) (newNode *Node32x6, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x6)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x5) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x5) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x5) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x5) Len() uint {
	return 5
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x5) Remove(idx uint) (newNode *Node32x4, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x4)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x5) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x5) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x5) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x6 is a sparse array node
type Node32x6 struct {
	Bitmap uint32
	Array [6]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x6) Add(idx uint, val generics.T1) (newNode *Node32x7, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x7)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x6) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x6) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x6) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x6) Len() uint {
	return 6
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x6) Remove(idx uint) (newNode *Node32x5, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x5)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x6) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x6) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x6) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x7 is a sparse array node
type Node32x7 struct {
	Bitmap uint32
	Array [7]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x7) Add(idx uint, val generics.T1) (newNode *Node32x8, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x8)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x7) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x7) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x7) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x7) Len() uint {
	return 7
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x7) Remove(idx uint) (newNode *Node32x6, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x6)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x7) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x7) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x7) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x8 is a sparse array node
type Node32x8 struct {
	Bitmap uint32
	Array [8]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x8) Add(idx uint, val generics.T1) (newNode *Node32x9, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x9)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x8) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x8) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x8) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x8) Len() uint {
	return 8
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x8) Remove(idx uint) (newNode *Node32x7, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x7)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x8) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x8) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x8) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x9 is a sparse array node
type Node32x9 struct {
	Bitmap uint32
	Array [9]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x9) Add(idx uint, val generics.T1) (newNode *Node32x10, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x10)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x9) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x9) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x9) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x9) Len() uint {
	return 9
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x9) Remove(idx uint) (newNode *Node32x8, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x8)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x9) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x9) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x9) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x10 is a sparse array node
type Node32x10 struct {
	Bitmap uint32
	Array [10]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x10) Add(idx uint, val generics.T1) (newNode *Node32x11, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x11)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x10) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x10) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x10) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x10) Len() uint {
	return 10
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x10) Remove(idx uint) (newNode *Node32x9, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x9)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x10) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x10) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x10) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x11 is a sparse array node
type Node32x11 struct {
	Bitmap uint32
	Array [11]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x11) Add(idx uint, val generics.T1) (newNode *Node32x12, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x12)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x11) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x11) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x11) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x11) Len() uint {
	return 11
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x11) Remove(idx uint) (newNode *Node32x10, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x10)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x11) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x11) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x11) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x12 is a sparse array node
type Node32x12 struct {
	Bitmap uint32
	Array [12]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x12) Add(idx uint, val generics.T1) (newNode *Node32x13, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x13)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x12) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x12) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x12) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x12) Len() uint {
	return 12
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x12) Remove(idx uint) (newNode *Node32x11, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x11)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x12) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x12) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x12) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x13 is a sparse array node
type Node32x13 struct {
	Bitmap uint32
	Array [13]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x13) Add(idx uint, val generics.T1) (newNode *Node32x14, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x14)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x13) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x13) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x13) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x13) Len() uint {
	return 13
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x13) Remove(idx uint) (newNode *Node32x12, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x12)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x13) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x13) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x13) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x14 is a sparse array node
type Node32x14 struct {
	Bitmap uint32
	Array [14]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x14) Add(idx uint, val generics.T1) (newNode *Node32x15, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x15)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x14) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x14) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x14) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x14) Len() uint {
	return 14
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x14) Remove(idx uint) (newNode *Node32x13, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x13)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x14) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x14) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x14) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x15 is a sparse array node
type Node32x15 struct {
	Bitmap uint32
	Array [15]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x15) Add(idx uint, val generics.T1) (newNode *Node32x16, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x16)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x15) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x15) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x15) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x15) Len() uint {
	return 15
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x15) Remove(idx uint) (newNode *Node32x14, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x14)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x15) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x15) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x15) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x16 is a sparse array node
type Node32x16 struct {
	Bitmap uint32
	Array [16]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x16) Add(idx uint, val generics.T1) (newNode *Node32x17, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x17)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x16) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x16) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x16) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x16) Len() uint {
	return 16
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x16) Remove(idx uint) (newNode *Node32x15, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x15)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x16) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x16) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x16) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x17 is a sparse array node
type Node32x17 struct {
	Bitmap uint32
	Array [17]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x17) Add(idx uint, val generics.T1) (newNode *Node32x18, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x18)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x17) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x17) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x17) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x17) Len() uint {
	return 17
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x17) Remove(idx uint) (newNode *Node32x16, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x16)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x17) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x17) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x17) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x18 is a sparse array node
type Node32x18 struct {
	Bitmap uint32
	Array [18]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x18) Add(idx uint, val generics.T1) (newNode *Node32x19, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x19)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x18) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x18) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x18) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x18) Len() uint {
	return 18
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x18) Remove(idx uint) (newNode *Node32x17, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x17)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x18) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x18) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x18) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x19 is a sparse array node
type Node32x19 struct {
	Bitmap uint32
	Array [19]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x19) Add(idx uint, val generics.T1) (newNode *Node32x20, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x20)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x19) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x19) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x19) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x19) Len() uint {
	return 19
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x19) Remove(idx uint) (newNode *Node32x18, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x18)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x19) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x19) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x19) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x20 is a sparse array node
type Node32x20 struct {
	Bitmap uint32
	Array [20]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x20) Add(idx uint, val generics.T1) (newNode *Node32x21, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x21)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x20) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x20) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x20) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x20) Len() uint {
	return 20
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x20) Remove(idx uint) (newNode *Node32x19, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x19)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x20) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x20) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x20) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x21 is a sparse array node
type Node32x21 struct {
	Bitmap uint32
	Array [21]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x21) Add(idx uint, val generics.T1) (newNode *Node32x22, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x22)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x21) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x21) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x21) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x21) Len() uint {
	return 21
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x21) Remove(idx uint) (newNode *Node32x20, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x20)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x21) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x21) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x21) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x22 is a sparse array node
type Node32x22 struct {
	Bitmap uint32
	Array [22]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x22) Add(idx uint, val generics.T1) (newNode *Node32x23, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x23)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x22) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x22) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x22) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x22) Len() uint {
	return 22
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x22) Remove(idx uint) (newNode *Node32x21, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x21)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x22) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x22) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x22) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x23 is a sparse array node
type Node32x23 struct {
	Bitmap uint32
	Array [23]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x23) Add(idx uint, val generics.T1) (newNode *Node32x24, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x24)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x23) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x23) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x23) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x23) Len() uint {
	return 23
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x23) Remove(idx uint) (newNode *Node32x22, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x22)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x23) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x23) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x23) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x24 is a sparse array node
type Node32x24 struct {
	Bitmap uint32
	Array [24]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x24) Add(idx uint, val generics.T1) (newNode *Node32x25, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x25)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x24) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x24) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x24) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x24) Len() uint {
	return 24
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x24) Remove(idx uint) (newNode *Node32x23, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x23)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x24) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x24) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x24) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x25 is a sparse array node
type Node32x25 struct {
	Bitmap uint32
	Array [25]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x25) Add(idx uint, val generics.T1) (newNode *Node32x26, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x26)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x25) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x25) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x25) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x25) Len() uint {
	return 25
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x25) Remove(idx uint) (newNode *Node32x24, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x24)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x25) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x25) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x25) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x26 is a sparse array node
type Node32x26 struct {
	Bitmap uint32
	Array [26]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x26) Add(idx uint, val generics.T1) (newNode *Node32x27, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x27)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x26) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x26) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x26) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x26) Len() uint {
	return 26
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x26) Remove(idx uint) (newNode *Node32x25, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x25)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x26) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x26) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x26) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x27 is a sparse array node
type Node32x27 struct {
	Bitmap uint32
	Array [27]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x27) Add(idx uint, val generics.T1) (newNode *Node32x28, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x28)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x27) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x27) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x27) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x27) Len() uint {
	return 27
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x27) Remove(idx uint) (newNode *Node32x26, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x26)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x27) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x27) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x27) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x28 is a sparse array node
type Node32x28 struct {
	Bitmap uint32
	Array [28]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x28) Add(idx uint, val generics.T1) (newNode *Node32x29, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x29)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x28) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x28) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x28) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x28) Len() uint {
	return 28
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x28) Remove(idx uint) (newNode *Node32x27, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x27)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x28) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x28) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x28) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x29 is a sparse array node
type Node32x29 struct {
	Bitmap uint32
	Array [29]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x29) Add(idx uint, val generics.T1) (newNode *Node32x30, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x30)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x29) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x29) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x29) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x29) Len() uint {
	return 29
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x29) Remove(idx uint) (newNode *Node32x28, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x28)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x29) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x29) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x29) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x30 is a sparse array node
type Node32x30 struct {
	Bitmap uint32
	Array [30]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x30) Add(idx uint, val generics.T1) (newNode *Node32x31, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x31)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x30) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x30) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x30) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x30) Len() uint {
	return 30
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x30) Remove(idx uint) (newNode *Node32x29, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x29)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x30) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x30) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x30) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x31 is a sparse array node
type Node32x31 struct {
	Bitmap uint32
	Array [31]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node32x31) Add(idx uint, val generics.T1) (newNode *Node32x32, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node32x32)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition32(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x31) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x31) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x31) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x31) Len() uint {
	return 31
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x31) Remove(idx uint) (newNode *Node32x30, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x30)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x31) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x31) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node32x31) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node32x32 is a sparse array node
type Node32x32 struct {
	Bitmap uint32
	Array [32]generics.T1
}

// Contains returns whether or not there's an element at the idx
func (node *Node32x32) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node32x32) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node32x32) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node32x32) Len() uint {
	return 32
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node32x32) Remove(idx uint) (newNode *Node32x31, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	newNode = new(Node32x31)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node32x32) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition32(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node32x32) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node, false

}

func (node *Node32x32) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x0 is a sparse array node
type Node64x0 struct {
	Bitmap uint64
	Array [0]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x0) Add(idx uint, val generics.T1) (newNode *Node64x1, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x1)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x0) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x0) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node, false

}

// Get returns the element at the given index
func (node *Node64x0) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x0) Len() uint {
	return 0
}

// Replace replaces the value
func (node *Node64x0) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x0) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x0) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x1 is a sparse array node
type Node64x1 struct {
	Bitmap uint64
	Array [1]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x1) Add(idx uint, val generics.T1) (newNode *Node64x2, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x2)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x1) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x1) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x1) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x1) Len() uint {
	return 1
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x1) Remove(idx uint) (newNode *Node64x0, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x0)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x1) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x1) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x1) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x2 is a sparse array node
type Node64x2 struct {
	Bitmap uint64
	Array [2]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x2) Add(idx uint, val generics.T1) (newNode *Node64x3, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x3)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x2) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x2) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x2) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x2) Len() uint {
	return 2
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x2) Remove(idx uint) (newNode *Node64x1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x1)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x2) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x2) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x2) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x3 is a sparse array node
type Node64x3 struct {
	Bitmap uint64
	Array [3]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x3) Add(idx uint, val generics.T1) (newNode *Node64x4, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x4)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x3) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x3) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x3) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x3) Len() uint {
	return 3
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x3) Remove(idx uint) (newNode *Node64x2, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x2)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x3) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x3) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x3) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x4 is a sparse array node
type Node64x4 struct {
	Bitmap uint64
	Array [4]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x4) Add(idx uint, val generics.T1) (newNode *Node64x5, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x5)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x4) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x4) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x4) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x4) Len() uint {
	return 4
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x4) Remove(idx uint) (newNode *Node64x3, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x3)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x4) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x4) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x4) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x5 is a sparse array node
type Node64x5 struct {
	Bitmap uint64
	Array [5]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x5) Add(idx uint, val generics.T1) (newNode *Node64x6, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x6)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x5) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x5) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x5) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x5) Len() uint {
	return 5
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x5) Remove(idx uint) (newNode *Node64x4, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x4)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x5) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x5) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x5) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x6 is a sparse array node
type Node64x6 struct {
	Bitmap uint64
	Array [6]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x6) Add(idx uint, val generics.T1) (newNode *Node64x7, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x7)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x6) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x6) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x6) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x6) Len() uint {
	return 6
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x6) Remove(idx uint) (newNode *Node64x5, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x5)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x6) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x6) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x6) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x7 is a sparse array node
type Node64x7 struct {
	Bitmap uint64
	Array [7]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x7) Add(idx uint, val generics.T1) (newNode *Node64x8, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x8)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x7) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x7) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x7) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x7) Len() uint {
	return 7
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x7) Remove(idx uint) (newNode *Node64x6, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x6)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x7) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x7) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x7) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x8 is a sparse array node
type Node64x8 struct {
	Bitmap uint64
	Array [8]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x8) Add(idx uint, val generics.T1) (newNode *Node64x9, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x9)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x8) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x8) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x8) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x8) Len() uint {
	return 8
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x8) Remove(idx uint) (newNode *Node64x7, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x7)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x8) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x8) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x8) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x9 is a sparse array node
type Node64x9 struct {
	Bitmap uint64
	Array [9]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x9) Add(idx uint, val generics.T1) (newNode *Node64x10, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x10)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x9) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x9) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x9) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x9) Len() uint {
	return 9
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x9) Remove(idx uint) (newNode *Node64x8, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x8)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x9) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x9) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x9) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x10 is a sparse array node
type Node64x10 struct {
	Bitmap uint64
	Array [10]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x10) Add(idx uint, val generics.T1) (newNode *Node64x11, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x11)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x10) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x10) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x10) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x10) Len() uint {
	return 10
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x10) Remove(idx uint) (newNode *Node64x9, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x9)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x10) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x10) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x10) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x11 is a sparse array node
type Node64x11 struct {
	Bitmap uint64
	Array [11]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x11) Add(idx uint, val generics.T1) (newNode *Node64x12, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x12)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x11) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x11) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x11) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x11) Len() uint {
	return 11
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x11) Remove(idx uint) (newNode *Node64x10, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x10)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x11) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x11) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x11) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x12 is a sparse array node
type Node64x12 struct {
	Bitmap uint64
	Array [12]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x12) Add(idx uint, val generics.T1) (newNode *Node64x13, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x13)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x12) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x12) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x12) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x12) Len() uint {
	return 12
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x12) Remove(idx uint) (newNode *Node64x11, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x11)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x12) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x12) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x12) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x13 is a sparse array node
type Node64x13 struct {
	Bitmap uint64
	Array [13]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x13) Add(idx uint, val generics.T1) (newNode *Node64x14, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x14)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x13) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x13) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x13) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x13) Len() uint {
	return 13
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x13) Remove(idx uint) (newNode *Node64x12, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x12)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x13) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x13) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x13) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x14 is a sparse array node
type Node64x14 struct {
	Bitmap uint64
	Array [14]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x14) Add(idx uint, val generics.T1) (newNode *Node64x15, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x15)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x14) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x14) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x14) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x14) Len() uint {
	return 14
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x14) Remove(idx uint) (newNode *Node64x13, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x13)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x14) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x14) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x14) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x15 is a sparse array node
type Node64x15 struct {
	Bitmap uint64
	Array [15]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x15) Add(idx uint, val generics.T1) (newNode *Node64x16, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x16)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x15) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x15) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x15) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x15) Len() uint {
	return 15
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x15) Remove(idx uint) (newNode *Node64x14, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x14)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x15) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x15) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x15) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x16 is a sparse array node
type Node64x16 struct {
	Bitmap uint64
	Array [16]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x16) Add(idx uint, val generics.T1) (newNode *Node64x17, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x17)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x16) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x16) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x16) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x16) Len() uint {
	return 16
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x16) Remove(idx uint) (newNode *Node64x15, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x15)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x16) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x16) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x16) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x17 is a sparse array node
type Node64x17 struct {
	Bitmap uint64
	Array [17]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x17) Add(idx uint, val generics.T1) (newNode *Node64x18, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x18)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x17) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x17) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x17) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x17) Len() uint {
	return 17
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x17) Remove(idx uint) (newNode *Node64x16, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x16)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x17) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x17) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x17) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x18 is a sparse array node
type Node64x18 struct {
	Bitmap uint64
	Array [18]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x18) Add(idx uint, val generics.T1) (newNode *Node64x19, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x19)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x18) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x18) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x18) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x18) Len() uint {
	return 18
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x18) Remove(idx uint) (newNode *Node64x17, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x17)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x18) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x18) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x18) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x19 is a sparse array node
type Node64x19 struct {
	Bitmap uint64
	Array [19]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x19) Add(idx uint, val generics.T1) (newNode *Node64x20, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x20)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x19) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x19) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x19) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x19) Len() uint {
	return 19
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x19) Remove(idx uint) (newNode *Node64x18, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x18)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x19) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x19) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x19) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x20 is a sparse array node
type Node64x20 struct {
	Bitmap uint64
	Array [20]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x20) Add(idx uint, val generics.T1) (newNode *Node64x21, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x21)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x20) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x20) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x20) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x20) Len() uint {
	return 20
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x20) Remove(idx uint) (newNode *Node64x19, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x19)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x20) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x20) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x20) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x21 is a sparse array node
type Node64x21 struct {
	Bitmap uint64
	Array [21]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x21) Add(idx uint, val generics.T1) (newNode *Node64x22, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x22)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x21) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x21) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x21) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x21) Len() uint {
	return 21
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x21) Remove(idx uint) (newNode *Node64x20, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x20)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x21) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x21) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x21) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x22 is a sparse array node
type Node64x22 struct {
	Bitmap uint64
	Array [22]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x22) Add(idx uint, val generics.T1) (newNode *Node64x23, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x23)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x22) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x22) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x22) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x22) Len() uint {
	return 22
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x22) Remove(idx uint) (newNode *Node64x21, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x21)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x22) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x22) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x22) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x23 is a sparse array node
type Node64x23 struct {
	Bitmap uint64
	Array [23]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x23) Add(idx uint, val generics.T1) (newNode *Node64x24, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x24)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x23) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x23) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x23) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x23) Len() uint {
	return 23
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x23) Remove(idx uint) (newNode *Node64x22, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x22)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x23) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x23) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x23) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x24 is a sparse array node
type Node64x24 struct {
	Bitmap uint64
	Array [24]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x24) Add(idx uint, val generics.T1) (newNode *Node64x25, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x25)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x24) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x24) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x24) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x24) Len() uint {
	return 24
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x24) Remove(idx uint) (newNode *Node64x23, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x23)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x24) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x24) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x24) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x25 is a sparse array node
type Node64x25 struct {
	Bitmap uint64
	Array [25]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x25) Add(idx uint, val generics.T1) (newNode *Node64x26, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x26)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x25) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x25) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x25) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x25) Len() uint {
	return 25
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x25) Remove(idx uint) (newNode *Node64x24, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x24)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x25) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x25) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x25) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x26 is a sparse array node
type Node64x26 struct {
	Bitmap uint64
	Array [26]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x26) Add(idx uint, val generics.T1) (newNode *Node64x27, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x27)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x26) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x26) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x26) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x26) Len() uint {
	return 26
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x26) Remove(idx uint) (newNode *Node64x25, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x25)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x26) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x26) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x26) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x27 is a sparse array node
type Node64x27 struct {
	Bitmap uint64
	Array [27]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x27) Add(idx uint, val generics.T1) (newNode *Node64x28, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x28)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x27) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x27) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x27) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x27) Len() uint {
	return 27
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x27) Remove(idx uint) (newNode *Node64x26, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x26)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x27) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x27) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x27) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x28 is a sparse array node
type Node64x28 struct {
	Bitmap uint64
	Array [28]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x28) Add(idx uint, val generics.T1) (newNode *Node64x29, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x29)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x28) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x28) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x28) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x28) Len() uint {
	return 28
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x28) Remove(idx uint) (newNode *Node64x27, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x27)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x28) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x28) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x28) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x29 is a sparse array node
type Node64x29 struct {
	Bitmap uint64
	Array [29]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x29) Add(idx uint, val generics.T1) (newNode *Node64x30, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x30)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x29) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x29) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x29) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x29) Len() uint {
	return 29
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x29) Remove(idx uint) (newNode *Node64x28, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x28)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x29) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x29) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x29) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x30 is a sparse array node
type Node64x30 struct {
	Bitmap uint64
	Array [30]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x30) Add(idx uint, val generics.T1) (newNode *Node64x31, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x31)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x30) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x30) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x30) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x30) Len() uint {
	return 30
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x30) Remove(idx uint) (newNode *Node64x29, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x29)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x30) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x30) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x30) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x31 is a sparse array node
type Node64x31 struct {
	Bitmap uint64
	Array [31]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x31) Add(idx uint, val generics.T1) (newNode *Node64x32, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x32)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x31) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x31) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x31) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x31) Len() uint {
	return 31
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x31) Remove(idx uint) (newNode *Node64x30, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x30)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x31) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x31) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x31) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x32 is a sparse array node
type Node64x32 struct {
	Bitmap uint64
	Array [32]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x32) Add(idx uint, val generics.T1) (newNode *Node64x33, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x33)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x32) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x32) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x32) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x32) Len() uint {
	return 32
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x32) Remove(idx uint) (newNode *Node64x31, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x31)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x32) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x32) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x32) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x33 is a sparse array node
type Node64x33 struct {
	Bitmap uint64
	Array [33]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x33) Add(idx uint, val generics.T1) (newNode *Node64x34, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x34)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x33) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x33) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x33) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x33) Len() uint {
	return 33
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x33) Remove(idx uint) (newNode *Node64x32, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x32)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x33) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x33) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x33) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x34 is a sparse array node
type Node64x34 struct {
	Bitmap uint64
	Array [34]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x34) Add(idx uint, val generics.T1) (newNode *Node64x35, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x35)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x34) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x34) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x34) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x34) Len() uint {
	return 34
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x34) Remove(idx uint) (newNode *Node64x33, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x33)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x34) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x34) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x34) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x35 is a sparse array node
type Node64x35 struct {
	Bitmap uint64
	Array [35]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x35) Add(idx uint, val generics.T1) (newNode *Node64x36, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x36)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x35) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x35) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x35) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x35) Len() uint {
	return 35
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x35) Remove(idx uint) (newNode *Node64x34, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x34)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x35) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x35) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x35) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x36 is a sparse array node
type Node64x36 struct {
	Bitmap uint64
	Array [36]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x36) Add(idx uint, val generics.T1) (newNode *Node64x37, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x37)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x36) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x36) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x36) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x36) Len() uint {
	return 36
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x36) Remove(idx uint) (newNode *Node64x35, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x35)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x36) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x36) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x36) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x37 is a sparse array node
type Node64x37 struct {
	Bitmap uint64
	Array [37]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x37) Add(idx uint, val generics.T1) (newNode *Node64x38, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x38)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x37) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x37) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x37) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x37) Len() uint {
	return 37
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x37) Remove(idx uint) (newNode *Node64x36, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x36)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x37) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x37) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x37) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x38 is a sparse array node
type Node64x38 struct {
	Bitmap uint64
	Array [38]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x38) Add(idx uint, val generics.T1) (newNode *Node64x39, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x39)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x38) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x38) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x38) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x38) Len() uint {
	return 38
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x38) Remove(idx uint) (newNode *Node64x37, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x37)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x38) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x38) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x38) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x39 is a sparse array node
type Node64x39 struct {
	Bitmap uint64
	Array [39]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x39) Add(idx uint, val generics.T1) (newNode *Node64x40, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x40)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x39) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x39) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x39) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x39) Len() uint {
	return 39
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x39) Remove(idx uint) (newNode *Node64x38, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x38)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x39) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x39) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x39) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x40 is a sparse array node
type Node64x40 struct {
	Bitmap uint64
	Array [40]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x40) Add(idx uint, val generics.T1) (newNode *Node64x41, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x41)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x40) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x40) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x40) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x40) Len() uint {
	return 40
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x40) Remove(idx uint) (newNode *Node64x39, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x39)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x40) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x40) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x40) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x41 is a sparse array node
type Node64x41 struct {
	Bitmap uint64
	Array [41]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x41) Add(idx uint, val generics.T1) (newNode *Node64x42, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x42)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x41) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x41) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x41) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x41) Len() uint {
	return 41
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x41) Remove(idx uint) (newNode *Node64x40, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x40)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x41) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x41) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x41) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x42 is a sparse array node
type Node64x42 struct {
	Bitmap uint64
	Array [42]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x42) Add(idx uint, val generics.T1) (newNode *Node64x43, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x43)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x42) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x42) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x42) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x42) Len() uint {
	return 42
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x42) Remove(idx uint) (newNode *Node64x41, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x41)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x42) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x42) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x42) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x43 is a sparse array node
type Node64x43 struct {
	Bitmap uint64
	Array [43]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x43) Add(idx uint, val generics.T1) (newNode *Node64x44, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x44)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x43) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x43) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x43) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x43) Len() uint {
	return 43
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x43) Remove(idx uint) (newNode *Node64x42, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x42)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x43) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x43) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x43) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x44 is a sparse array node
type Node64x44 struct {
	Bitmap uint64
	Array [44]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x44) Add(idx uint, val generics.T1) (newNode *Node64x45, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x45)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x44) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x44) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x44) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x44) Len() uint {
	return 44
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x44) Remove(idx uint) (newNode *Node64x43, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x43)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x44) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x44) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x44) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x45 is a sparse array node
type Node64x45 struct {
	Bitmap uint64
	Array [45]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x45) Add(idx uint, val generics.T1) (newNode *Node64x46, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x46)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x45) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x45) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x45) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x45) Len() uint {
	return 45
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x45) Remove(idx uint) (newNode *Node64x44, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x44)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x45) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x45) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x45) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x46 is a sparse array node
type Node64x46 struct {
	Bitmap uint64
	Array [46]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x46) Add(idx uint, val generics.T1) (newNode *Node64x47, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x47)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x46) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x46) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x46) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x46) Len() uint {
	return 46
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x46) Remove(idx uint) (newNode *Node64x45, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x45)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x46) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x46) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x46) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x47 is a sparse array node
type Node64x47 struct {
	Bitmap uint64
	Array [47]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x47) Add(idx uint, val generics.T1) (newNode *Node64x48, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x48)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x47) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x47) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x47) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x47) Len() uint {
	return 47
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x47) Remove(idx uint) (newNode *Node64x46, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x46)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x47) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x47) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x47) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x48 is a sparse array node
type Node64x48 struct {
	Bitmap uint64
	Array [48]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x48) Add(idx uint, val generics.T1) (newNode *Node64x49, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x49)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x48) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x48) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x48) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x48) Len() uint {
	return 48
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x48) Remove(idx uint) (newNode *Node64x47, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x47)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x48) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x48) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x48) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x49 is a sparse array node
type Node64x49 struct {
	Bitmap uint64
	Array [49]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x49) Add(idx uint, val generics.T1) (newNode *Node64x50, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x50)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x49) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x49) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x49) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x49) Len() uint {
	return 49
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x49) Remove(idx uint) (newNode *Node64x48, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x48)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x49) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x49) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x49) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x50 is a sparse array node
type Node64x50 struct {
	Bitmap uint64
	Array [50]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x50) Add(idx uint, val generics.T1) (newNode *Node64x51, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x51)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x50) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x50) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x50) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x50) Len() uint {
	return 50
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x50) Remove(idx uint) (newNode *Node64x49, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x49)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x50) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x50) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x50) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x51 is a sparse array node
type Node64x51 struct {
	Bitmap uint64
	Array [51]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x51) Add(idx uint, val generics.T1) (newNode *Node64x52, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x52)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x51) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x51) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x51) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x51) Len() uint {
	return 51
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x51) Remove(idx uint) (newNode *Node64x50, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x50)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x51) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x51) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x51) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x52 is a sparse array node
type Node64x52 struct {
	Bitmap uint64
	Array [52]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x52) Add(idx uint, val generics.T1) (newNode *Node64x53, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x53)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x52) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x52) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x52) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x52) Len() uint {
	return 52
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x52) Remove(idx uint) (newNode *Node64x51, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x51)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x52) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x52) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x52) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x53 is a sparse array node
type Node64x53 struct {
	Bitmap uint64
	Array [53]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x53) Add(idx uint, val generics.T1) (newNode *Node64x54, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x54)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x53) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x53) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x53) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x53) Len() uint {
	return 53
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x53) Remove(idx uint) (newNode *Node64x52, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x52)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x53) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x53) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x53) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x54 is a sparse array node
type Node64x54 struct {
	Bitmap uint64
	Array [54]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x54) Add(idx uint, val generics.T1) (newNode *Node64x55, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x55)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x54) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x54) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x54) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x54) Len() uint {
	return 54
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x54) Remove(idx uint) (newNode *Node64x53, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x53)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x54) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x54) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x54) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x55 is a sparse array node
type Node64x55 struct {
	Bitmap uint64
	Array [55]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x55) Add(idx uint, val generics.T1) (newNode *Node64x56, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x56)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x55) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x55) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x55) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x55) Len() uint {
	return 55
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x55) Remove(idx uint) (newNode *Node64x54, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x54)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x55) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x55) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x55) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x56 is a sparse array node
type Node64x56 struct {
	Bitmap uint64
	Array [56]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x56) Add(idx uint, val generics.T1) (newNode *Node64x57, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x57)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x56) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x56) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x56) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x56) Len() uint {
	return 56
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x56) Remove(idx uint) (newNode *Node64x55, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x55)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x56) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x56) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x56) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x57 is a sparse array node
type Node64x57 struct {
	Bitmap uint64
	Array [57]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x57) Add(idx uint, val generics.T1) (newNode *Node64x58, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x58)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x57) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x57) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x57) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x57) Len() uint {
	return 57
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x57) Remove(idx uint) (newNode *Node64x56, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x56)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x57) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x57) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x57) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x58 is a sparse array node
type Node64x58 struct {
	Bitmap uint64
	Array [58]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x58) Add(idx uint, val generics.T1) (newNode *Node64x59, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x59)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x58) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x58) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x58) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x58) Len() uint {
	return 58
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x58) Remove(idx uint) (newNode *Node64x57, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x57)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x58) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x58) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x58) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x59 is a sparse array node
type Node64x59 struct {
	Bitmap uint64
	Array [59]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x59) Add(idx uint, val generics.T1) (newNode *Node64x60, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x60)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x59) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x59) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x59) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x59) Len() uint {
	return 59
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x59) Remove(idx uint) (newNode *Node64x58, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x58)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x59) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x59) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x59) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x60 is a sparse array node
type Node64x60 struct {
	Bitmap uint64
	Array [60]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x60) Add(idx uint, val generics.T1) (newNode *Node64x61, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x61)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x60) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x60) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x60) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x60) Len() uint {
	return 60
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x60) Remove(idx uint) (newNode *Node64x59, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x59)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x60) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x60) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x60) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x61 is a sparse array node
type Node64x61 struct {
	Bitmap uint64
	Array [61]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x61) Add(idx uint, val generics.T1) (newNode *Node64x62, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x62)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x61) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x61) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x61) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x61) Len() uint {
	return 61
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x61) Remove(idx uint) (newNode *Node64x60, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x60)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x61) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x61) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x61) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x62 is a sparse array node
type Node64x62 struct {
	Bitmap uint64
	Array [62]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x62) Add(idx uint, val generics.T1) (newNode *Node64x63, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x63)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x62) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x62) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x62) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x62) Len() uint {
	return 62
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x62) Remove(idx uint) (newNode *Node64x61, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x61)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x62) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x62) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x62) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x63 is a sparse array node
type Node64x63 struct {
	Bitmap uint64
	Array [63]generics.T1
}

// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *Node64x63) Add(idx uint, val generics.T1) (newNode *Node64x64, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(Node64x64)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition64(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x63) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x63) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x63) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x63) Len() uint {
	return 63
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x63) Remove(idx uint) (newNode *Node64x62, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x62)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x63) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x63) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node.Add(idx, val)

}

func (node *Node64x63) Slice() []generics.T1 {
	return node.Array[:]
}

// A Node64x64 is a sparse array node
type Node64x64 struct {
	Bitmap uint64
	Array [64]generics.T1
}

// Contains returns whether or not there's an element at the idx
func (node *Node64x64) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *Node64x64) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}

	return node.Remove(idx)

}

// Get returns the element at the given index
func (node *Node64x64) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *Node64x64) Len() uint {
	return 64
}

// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *Node64x64) Remove(idx uint) (newNode *Node64x63, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	newNode = new(Node64x63)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}

// Replace replaces the value
func (node *Node64x64) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition64(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *Node64x64) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}

	return node, false

}

func (node *Node64x64) Slice() []generics.T1 {
	return node.Array[:]
}
