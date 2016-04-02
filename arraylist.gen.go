package container
import "bytes"



// An ArrayListOfUint8 is a growable, dynamic array
type ArrayListOfUint8 struct {
	items []uint8
}

// NewArrayListOfUint8 creates a new ArrayListOfUint8
func NewArrayListOfUint8() *ArrayListOfUint8 {
	return new(ArrayListOfUint8)
}

// Add adds an item to the end of the list
func (list *ArrayListOfUint8) Add(item uint8) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfUint8) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfUint8) Contains(item uint8) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfUint8) Fill(items []uint8) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfUint8) Get(idx int) uint8 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfUint8) IndexOf(item uint8) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfUint8) Insert(idx int, item uint8) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfUint8) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfUint8) Remove(item uint8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfUint8) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfUint8) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfUint8) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfUint8) Truncate(sz int) {
	items := make([]uint8, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfUint16 is a growable, dynamic array
type ArrayListOfUint16 struct {
	items []uint16
}

// NewArrayListOfUint16 creates a new ArrayListOfUint16
func NewArrayListOfUint16() *ArrayListOfUint16 {
	return new(ArrayListOfUint16)
}

// Add adds an item to the end of the list
func (list *ArrayListOfUint16) Add(item uint16) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfUint16) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfUint16) Contains(item uint16) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfUint16) Fill(items []uint16) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfUint16) Get(idx int) uint16 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfUint16) IndexOf(item uint16) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfUint16) Insert(idx int, item uint16) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfUint16) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfUint16) Remove(item uint16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfUint16) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfUint16) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfUint16) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfUint16) Truncate(sz int) {
	items := make([]uint16, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfUint32 is a growable, dynamic array
type ArrayListOfUint32 struct {
	items []uint32
}

// NewArrayListOfUint32 creates a new ArrayListOfUint32
func NewArrayListOfUint32() *ArrayListOfUint32 {
	return new(ArrayListOfUint32)
}

// Add adds an item to the end of the list
func (list *ArrayListOfUint32) Add(item uint32) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfUint32) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfUint32) Contains(item uint32) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfUint32) Fill(items []uint32) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfUint32) Get(idx int) uint32 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfUint32) IndexOf(item uint32) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfUint32) Insert(idx int, item uint32) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfUint32) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfUint32) Remove(item uint32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfUint32) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfUint32) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfUint32) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfUint32) Truncate(sz int) {
	items := make([]uint32, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfUint64 is a growable, dynamic array
type ArrayListOfUint64 struct {
	items []uint64
}

// NewArrayListOfUint64 creates a new ArrayListOfUint64
func NewArrayListOfUint64() *ArrayListOfUint64 {
	return new(ArrayListOfUint64)
}

// Add adds an item to the end of the list
func (list *ArrayListOfUint64) Add(item uint64) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfUint64) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfUint64) Contains(item uint64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfUint64) Fill(items []uint64) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfUint64) Get(idx int) uint64 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfUint64) IndexOf(item uint64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfUint64) Insert(idx int, item uint64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfUint64) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfUint64) Remove(item uint64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfUint64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfUint64) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfUint64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfUint64) Truncate(sz int) {
	items := make([]uint64, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfInt8 is a growable, dynamic array
type ArrayListOfInt8 struct {
	items []int8
}

// NewArrayListOfInt8 creates a new ArrayListOfInt8
func NewArrayListOfInt8() *ArrayListOfInt8 {
	return new(ArrayListOfInt8)
}

// Add adds an item to the end of the list
func (list *ArrayListOfInt8) Add(item int8) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfInt8) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfInt8) Contains(item int8) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfInt8) Fill(items []int8) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfInt8) Get(idx int) int8 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfInt8) IndexOf(item int8) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfInt8) Insert(idx int, item int8) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfInt8) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfInt8) Remove(item int8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfInt8) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfInt8) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfInt8) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfInt8) Truncate(sz int) {
	items := make([]int8, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfInt16 is a growable, dynamic array
type ArrayListOfInt16 struct {
	items []int16
}

// NewArrayListOfInt16 creates a new ArrayListOfInt16
func NewArrayListOfInt16() *ArrayListOfInt16 {
	return new(ArrayListOfInt16)
}

// Add adds an item to the end of the list
func (list *ArrayListOfInt16) Add(item int16) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfInt16) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfInt16) Contains(item int16) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfInt16) Fill(items []int16) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfInt16) Get(idx int) int16 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfInt16) IndexOf(item int16) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfInt16) Insert(idx int, item int16) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfInt16) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfInt16) Remove(item int16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfInt16) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfInt16) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfInt16) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfInt16) Truncate(sz int) {
	items := make([]int16, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfInt32 is a growable, dynamic array
type ArrayListOfInt32 struct {
	items []int32
}

// NewArrayListOfInt32 creates a new ArrayListOfInt32
func NewArrayListOfInt32() *ArrayListOfInt32 {
	return new(ArrayListOfInt32)
}

// Add adds an item to the end of the list
func (list *ArrayListOfInt32) Add(item int32) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfInt32) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfInt32) Contains(item int32) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfInt32) Fill(items []int32) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfInt32) Get(idx int) int32 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfInt32) IndexOf(item int32) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfInt32) Insert(idx int, item int32) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfInt32) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfInt32) Remove(item int32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfInt32) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfInt32) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfInt32) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfInt32) Truncate(sz int) {
	items := make([]int32, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfInt64 is a growable, dynamic array
type ArrayListOfInt64 struct {
	items []int64
}

// NewArrayListOfInt64 creates a new ArrayListOfInt64
func NewArrayListOfInt64() *ArrayListOfInt64 {
	return new(ArrayListOfInt64)
}

// Add adds an item to the end of the list
func (list *ArrayListOfInt64) Add(item int64) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfInt64) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfInt64) Contains(item int64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfInt64) Fill(items []int64) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfInt64) Get(idx int) int64 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfInt64) IndexOf(item int64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfInt64) Insert(idx int, item int64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfInt64) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfInt64) Remove(item int64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfInt64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfInt64) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfInt64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfInt64) Truncate(sz int) {
	items := make([]int64, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfFloat32 is a growable, dynamic array
type ArrayListOfFloat32 struct {
	items []float32
}

// NewArrayListOfFloat32 creates a new ArrayListOfFloat32
func NewArrayListOfFloat32() *ArrayListOfFloat32 {
	return new(ArrayListOfFloat32)
}

// Add adds an item to the end of the list
func (list *ArrayListOfFloat32) Add(item float32) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfFloat32) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfFloat32) Contains(item float32) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfFloat32) Fill(items []float32) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfFloat32) Get(idx int) float32 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfFloat32) IndexOf(item float32) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfFloat32) Insert(idx int, item float32) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfFloat32) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfFloat32) Remove(item float32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfFloat32) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfFloat32) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfFloat32) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfFloat32) Truncate(sz int) {
	items := make([]float32, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfFloat64 is a growable, dynamic array
type ArrayListOfFloat64 struct {
	items []float64
}

// NewArrayListOfFloat64 creates a new ArrayListOfFloat64
func NewArrayListOfFloat64() *ArrayListOfFloat64 {
	return new(ArrayListOfFloat64)
}

// Add adds an item to the end of the list
func (list *ArrayListOfFloat64) Add(item float64) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfFloat64) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfFloat64) Contains(item float64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfFloat64) Fill(items []float64) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfFloat64) Get(idx int) float64 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfFloat64) IndexOf(item float64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfFloat64) Insert(idx int, item float64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfFloat64) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfFloat64) Remove(item float64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfFloat64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfFloat64) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfFloat64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfFloat64) Truncate(sz int) {
	items := make([]float64, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfComplex64 is a growable, dynamic array
type ArrayListOfComplex64 struct {
	items []complex64
}

// NewArrayListOfComplex64 creates a new ArrayListOfComplex64
func NewArrayListOfComplex64() *ArrayListOfComplex64 {
	return new(ArrayListOfComplex64)
}

// Add adds an item to the end of the list
func (list *ArrayListOfComplex64) Add(item complex64) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfComplex64) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfComplex64) Contains(item complex64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfComplex64) Fill(items []complex64) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfComplex64) Get(idx int) complex64 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfComplex64) IndexOf(item complex64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfComplex64) Insert(idx int, item complex64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfComplex64) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfComplex64) Remove(item complex64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfComplex64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfComplex64) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfComplex64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfComplex64) Truncate(sz int) {
	items := make([]complex64, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfComplex128 is a growable, dynamic array
type ArrayListOfComplex128 struct {
	items []complex128
}

// NewArrayListOfComplex128 creates a new ArrayListOfComplex128
func NewArrayListOfComplex128() *ArrayListOfComplex128 {
	return new(ArrayListOfComplex128)
}

// Add adds an item to the end of the list
func (list *ArrayListOfComplex128) Add(item complex128) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfComplex128) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfComplex128) Contains(item complex128) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfComplex128) Fill(items []complex128) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfComplex128) Get(idx int) complex128 {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfComplex128) IndexOf(item complex128) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfComplex128) Insert(idx int, item complex128) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfComplex128) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfComplex128) Remove(item complex128) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfComplex128) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfComplex128) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfComplex128) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfComplex128) Truncate(sz int) {
	items := make([]complex128, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfByte is a growable, dynamic array
type ArrayListOfByte struct {
	items []byte
}

// NewArrayListOfByte creates a new ArrayListOfByte
func NewArrayListOfByte() *ArrayListOfByte {
	return new(ArrayListOfByte)
}

// Add adds an item to the end of the list
func (list *ArrayListOfByte) Add(item byte) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfByte) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfByte) Contains(item byte) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfByte) Fill(items []byte) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfByte) Get(idx int) byte {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfByte) IndexOf(item byte) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfByte) Insert(idx int, item byte) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfByte) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfByte) Remove(item byte) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfByte) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfByte) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfByte) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfByte) Truncate(sz int) {
	items := make([]byte, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfRune is a growable, dynamic array
type ArrayListOfRune struct {
	items []rune
}

// NewArrayListOfRune creates a new ArrayListOfRune
func NewArrayListOfRune() *ArrayListOfRune {
	return new(ArrayListOfRune)
}

// Add adds an item to the end of the list
func (list *ArrayListOfRune) Add(item rune) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfRune) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfRune) Contains(item rune) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfRune) Fill(items []rune) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfRune) Get(idx int) rune {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfRune) IndexOf(item rune) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfRune) Insert(idx int, item rune) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfRune) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfRune) Remove(item rune) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfRune) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfRune) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfRune) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfRune) Truncate(sz int) {
	items := make([]rune, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfUint is a growable, dynamic array
type ArrayListOfUint struct {
	items []uint
}

// NewArrayListOfUint creates a new ArrayListOfUint
func NewArrayListOfUint() *ArrayListOfUint {
	return new(ArrayListOfUint)
}

// Add adds an item to the end of the list
func (list *ArrayListOfUint) Add(item uint) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfUint) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfUint) Contains(item uint) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfUint) Fill(items []uint) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfUint) Get(idx int) uint {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfUint) IndexOf(item uint) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfUint) Insert(idx int, item uint) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfUint) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfUint) Remove(item uint) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfUint) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfUint) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfUint) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfUint) Truncate(sz int) {
	items := make([]uint, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfInt is a growable, dynamic array
type ArrayListOfInt struct {
	items []int
}

// NewArrayListOfInt creates a new ArrayListOfInt
func NewArrayListOfInt() *ArrayListOfInt {
	return new(ArrayListOfInt)
}

// Add adds an item to the end of the list
func (list *ArrayListOfInt) Add(item int) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfInt) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfInt) Contains(item int) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfInt) Fill(items []int) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfInt) Get(idx int) int {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfInt) IndexOf(item int) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfInt) Insert(idx int, item int) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfInt) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfInt) Remove(item int) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfInt) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfInt) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfInt) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfInt) Truncate(sz int) {
	items := make([]int, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfUintptr is a growable, dynamic array
type ArrayListOfUintptr struct {
	items []uintptr
}

// NewArrayListOfUintptr creates a new ArrayListOfUintptr
func NewArrayListOfUintptr() *ArrayListOfUintptr {
	return new(ArrayListOfUintptr)
}

// Add adds an item to the end of the list
func (list *ArrayListOfUintptr) Add(item uintptr) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfUintptr) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfUintptr) Contains(item uintptr) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfUintptr) Fill(items []uintptr) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfUintptr) Get(idx int) uintptr {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfUintptr) IndexOf(item uintptr) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfUintptr) Insert(idx int, item uintptr) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfUintptr) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfUintptr) Remove(item uintptr) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfUintptr) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfUintptr) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfUintptr) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfUintptr) Truncate(sz int) {
	items := make([]uintptr, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfString is a growable, dynamic array
type ArrayListOfString struct {
	items []string
}

// NewArrayListOfString creates a new ArrayListOfString
func NewArrayListOfString() *ArrayListOfString {
	return new(ArrayListOfString)
}

// Add adds an item to the end of the list
func (list *ArrayListOfString) Add(item string) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfString) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfString) Contains(item string) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfString) Fill(items []string) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfString) Get(idx int) string {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfString) IndexOf(item string) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfString) Insert(idx int, item string) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfString) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfString) Remove(item string) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfString) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfString) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfString) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfString) Truncate(sz int) {
	items := make([]string, sz)
	copy(items, list.items)
	list.items = items
}



// An ArrayListOfBytes is a growable, dynamic array
type ArrayListOfBytes struct {
	items [][]byte
}

// NewArrayListOfBytes creates a new ArrayListOfBytes
func NewArrayListOfBytes() *ArrayListOfBytes {
	return new(ArrayListOfBytes)
}

// Add adds an item to the end of the list
func (list *ArrayListOfBytes) Add(item []byte) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayListOfBytes) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayListOfBytes) Contains(item []byte) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayListOfBytes) Fill(items [][]byte) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayListOfBytes) Get(idx int) []byte {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayListOfBytes) IndexOf(item []byte) int {
	for i, t := range list.items {
		if bytes.Equal(t, item) {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayListOfBytes) Insert(idx int, item []byte) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayListOfBytes) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayListOfBytes) Remove(item []byte) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayListOfBytes) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	// reduce the size of the array if we've gotten too big
	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayListOfBytes) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayListOfBytes) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayListOfBytes) Truncate(sz int) {
	items := make([][]byte, sz)
	copy(items, list.items)
	list.items = items
}
