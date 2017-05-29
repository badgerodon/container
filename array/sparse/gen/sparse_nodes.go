package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	txt := `package sparse

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

`

	for _, i := range []int{8, 16, 32, 64} {
		for j := 0; j <= i; j++ {
			name := `Node` + fmt.Sprint(i) + `x` + fmt.Sprint(j)
			prevName := `Node` + fmt.Sprint(i) + `x` + fmt.Sprint(j-1)
			nextName := `Node` + fmt.Sprint(i) + `x` + fmt.Sprint(j+1)
			txt += `
// A ` + name + ` is a sparse array node
type ` + name + ` struct {
	Bitmap uint` + fmt.Sprint(i) + `
	Array [` + fmt.Sprint(j) + `]generics.T1
}
`

			if j < i {
				txt += `
// Add adds a new element to the sparse array. It returns the new node and true if successful.
func (node *` + name + `) Add(idx uint, val generics.T1) (newNode *` + nextName + `, ok bool) {
	if node.Bitmap & (1 << idx) > 0 {
		return nil, false
	}
	newNode = new(` + nextName + `)
	newNode.Bitmap = node.Bitmap | (1 << idx)
	pos := getArrayPosition` + fmt.Sprint(i) + `(newNode.Bitmap, idx)
	copy(newNode.Array[:], node.Array[:pos])
	newNode.Array[pos] = val
	copy(newNode.Array[pos+1:], node.Array[pos:])
	return newNode, true
}
`
			}

			txt += `
// Contains returns whether or not there's an element at the idx
func (node *` + name + `) Contains(idx uint) bool {
	return node.Bitmap & (1 << idx) > 0
}

// Delete deletes an element from the Node. It returns the new Node
func (node *` + name + `) Delete(idx uint) (Node, bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return node, false
	}
`
			if j == 0 {
				txt += `
	return node, false
`
			} else {
				txt += `
	return node.Remove(idx)
`
			}

			txt += `
}
`

			txt += `
// Get returns the element at the given index
func (node *` + name + `) Get(idx uint) (val generics.T1, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return val, false
	}
	pos := getArrayPosition` + fmt.Sprint(i) + `(node.Bitmap, idx)
	return node.Array[pos], true
}

// Len returns the number of elements in the array
func (node *` + name + `) Len() uint {
	return ` + fmt.Sprint(j) + `
}
`
			if j > 0 {
				txt += `
// Remove removes an element from the sparse array. It returns the new node and true if successful.
func (node *` + name + `) Remove(idx uint) (newNode *` + prevName + `, ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return nil, false
	}
	pos := getArrayPosition` + fmt.Sprint(i) + `(node.Bitmap, idx)
	newNode = new(` + prevName + `)
	newNode.Bitmap = node.Bitmap ^ (1 << idx)
	copy(newNode.Array[:], node.Array[:pos])
	copy(newNode.Array[pos:], node.Array[pos+1:])
	return newNode, true
}
`
			}

			txt += `
// Replace replaces the value
func (node *` + name + `) Replace(idx uint, val generics.T1) (ok bool) {
	if node.Bitmap & (1 << idx) == 0 {
		return false
	}
	pos := getArrayPosition` + fmt.Sprint(i) + `(node.Bitmap, idx)
	node.Array[pos] = val
	return true
}

func (node *` + name + `) Set(idx uint, val generics.T1) (Node, bool) {
	if node.Replace(idx, val) {
		return node, true
	}
`
			if i == j {
				txt += `
	return node, false
`
			} else {
				txt += `
	return node.Add(idx, val)
`
			}

			txt += `
}

func (node *` + name + `) Slice() []generics.T1 {
	return node.Array[:]
}
`
		}
	}

	ioutil.WriteFile("nodes.gen.go", []byte(txt), 0755)
}
