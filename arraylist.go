package container

// An ArrayList is a growable, dynamic array
type ArrayList struct {
	items []interface{} /* ITEM */
}

// NewArrayList creates a new ArrayList
func NewArrayList() *ArrayList {
	return new(ArrayList)
}

// Add adds an item to the end of the list
func (list *ArrayList) Add(item interface{} /* ITEM */) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayList) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayList) Contains(item interface{} /* ITEM */) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayList) Fill(items []interface{} /* ITEM */) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayList) Get(idx int) interface{} /* ITEM */ {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList) IndexOf(item interface{} /* ITEM */) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList) Insert(idx int, item interface{} /* ITEM */) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayList) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList) Remove(item interface{} /* ITEM */) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayList) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayList) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList) Truncate(sz int) {
	items := make([]interface{} /* ITEM */, sz)
	copy(items, list.items)
	list.items = items
}
