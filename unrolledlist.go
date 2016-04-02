package container

// An unrolled list is a linked-list of arrays
type UnrolledList struct {
	items     [][]interface{} /* ITEM */
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
func (list *UnrolledList) Add(item interface{} /* ITEM */) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledList) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList) Contains(item interface{} /* ITEM */) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledList) Fill(items ...interface{} /* ITEM */) {
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
func (list *UnrolledList) Get(idx int) interface{} /* ITEM */ {
	if idx > list.size {
		var def interface{} /* ITEM */
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList) IndexOf(item interface{} /* ITEM */) int {
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
func (list *UnrolledList) Insert(idx int, item interface{} /* ITEM */) {
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
func (list *UnrolledList) Remove(item interface{} /* ITEM */) {
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
func (list *UnrolledList) Set(idx int, item interface{} /* ITEM */) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledList) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledList) get(idx int) interface{} /* ITEM */ {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]interface{} /* ITEM */, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList) set(idx int, item interface{} /* ITEM */) {
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
