package container

//go:generate goreify github.com/badgerodon/container.UnrolledList numeric,string

import "github.com/badgerodon/goreify/generics"

// An UnrolledList is a linked-list of arrays
type UnrolledList struct {
	items     [][]generics.T1
	size      int
	arraySize int
}

// NewUnrolledList creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList(arraySize int) *UnrolledList {
	return &UnrolledList{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledList) Add(item generics.T1) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledList) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList) Contains(item generics.T1) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledList) Fill(items ...generics.T1) {
	if len(items) == 0 {
		return
	}

	idx := list.size
	list.grow(list.size + len(items))
	for ; len(items) > 0; idx++ {
		amt := list.arraySize - idx%list.arraySize
		copy(list.items[idx/list.arraySize][idx%list.arraySize:], items)
		if len(items) < amt {
			break
		}
		items = items[amt:]
	}
}

// Get returns the item at the index in the list
func (list *UnrolledList) Get(idx int) generics.T1 {
	if idx > list.size {
		var def generics.T1
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList) IndexOf(item generics.T1) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// Insert inserts the item into the list
func (list *UnrolledList) Insert(idx int, item generics.T1) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledList) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledList) Remove(item generics.T1) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledList) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledList) Set(idx int, item generics.T1) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledList) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledList) get(idx int) generics.T1 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]generics.T1, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList) set(idx int, item generics.T1) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList) shrink(size int) {
	if list.size < size {
		return
	}

	idx := size / list.arraySize
	for idx+1 < len(list.items) {
		list.items[len(list.items)-1] = nil
		list.items = list.items[:len(list.items)-1]
	}

	list.size = size
}
