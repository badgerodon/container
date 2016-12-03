package container

type ArrayList_int struct {
	items []int
}

type ArrayList_int8 struct {
	items []int8
}

type ArrayList_int16 struct {
	items []int16
}

type ArrayList_int32 struct {
	items []int32
}

type ArrayList_int64 struct {
	items []int64
}

type ArrayList_uint struct {
	items []uint
}

type ArrayList_uint8 struct {
	items []uint8
}

type ArrayList_uint16 struct {
	items []uint16
}

type ArrayList_uint32 struct {
	items []uint32
}

type ArrayList_uint64 struct {
	items []uint64
}

type ArrayList_float32 struct {
	items []float32
}

type ArrayList_float64 struct {
	items []float64
}

type ArrayList_complex64 struct {
	items []complex64
}

type ArrayList_complex128 struct {
	items []complex128
}

type ArrayList_string struct {
	items []string
}

// NewArrayList_int creates a new ArrayList
func NewArrayList_int() *ArrayList_int {
	return new(ArrayList_int)
}

// NewArrayList_int8 creates a new ArrayList
func NewArrayList_int8() *ArrayList_int8 {
	return new(ArrayList_int8)
}

// NewArrayList_int16 creates a new ArrayList
func NewArrayList_int16() *ArrayList_int16 {
	return new(ArrayList_int16)
}

// NewArrayList_int32 creates a new ArrayList
func NewArrayList_int32() *ArrayList_int32 {
	return new(ArrayList_int32)
}

// NewArrayList_int64 creates a new ArrayList
func NewArrayList_int64() *ArrayList_int64 {
	return new(ArrayList_int64)
}

// NewArrayList_uint creates a new ArrayList
func NewArrayList_uint() *ArrayList_uint {
	return new(ArrayList_uint)
}

// NewArrayList_uint8 creates a new ArrayList
func NewArrayList_uint8() *ArrayList_uint8 {
	return new(ArrayList_uint8)
}

// NewArrayList_uint16 creates a new ArrayList
func NewArrayList_uint16() *ArrayList_uint16 {
	return new(ArrayList_uint16)
}

// NewArrayList_uint32 creates a new ArrayList
func NewArrayList_uint32() *ArrayList_uint32 {
	return new(ArrayList_uint32)
}

// NewArrayList_uint64 creates a new ArrayList
func NewArrayList_uint64() *ArrayList_uint64 {
	return new(ArrayList_uint64)
}

// NewArrayList_float32 creates a new ArrayList
func NewArrayList_float32() *ArrayList_float32 {
	return new(ArrayList_float32)
}

// NewArrayList_float64 creates a new ArrayList
func NewArrayList_float64() *ArrayList_float64 {
	return new(ArrayList_float64)
}

// NewArrayList_complex64 creates a new ArrayList
func NewArrayList_complex64() *ArrayList_complex64 {
	return new(ArrayList_complex64)
}

// NewArrayList_complex128 creates a new ArrayList
func NewArrayList_complex128() *ArrayList_complex128 {
	return new(ArrayList_complex128)
}

// NewArrayList_string creates a new ArrayList
func NewArrayList_string() *ArrayList_string {
	return new(ArrayList_string)
}

// Add adds an item to the end of the list
func (list *ArrayList_int) Add(item int) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_int8) Add(item int8) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_int16) Add(item int16) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_int32) Add(item int32) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_int64) Add(item int64) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_uint) Add(item uint) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_uint8) Add(item uint8) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_uint16) Add(item uint16) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_uint32) Add(item uint32) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_uint64) Add(item uint64) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_float32) Add(item float32) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_float64) Add(item float64) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_complex64) Add(item complex64) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_complex128) Add(item complex128) {
	list.items = append(list.items, item)
}

// Add adds an item to the end of the list
func (list *ArrayList_string) Add(item string) {
	list.items = append(list.items, item)
}

// Clear removes all items from the list
func (list *ArrayList_int) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_int8) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_int16) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_int32) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_int64) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_uint) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_uint8) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_uint16) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_uint32) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_uint64) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_float32) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_float64) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_complex64) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_complex128) Clear() {
	list.items = nil
}

// Clear removes all items from the list
func (list *ArrayList_string) Clear() {
	list.items = nil
}

// Contains determines whether an element is in the list
func (list *ArrayList_int) Contains(item int) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_int8) Contains(item int8) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_int16) Contains(item int16) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_int32) Contains(item int32) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_int64) Contains(item int64) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_uint) Contains(item uint) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_uint8) Contains(item uint8) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_uint16) Contains(item uint16) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_uint32) Contains(item uint32) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_uint64) Contains(item uint64) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_float32) Contains(item float32) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_float64) Contains(item float64) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_complex64) Contains(item complex64) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_complex128) Contains(item complex128) bool {
	return list.IndexOf(item) >= 0
}

// Contains determines whether an element is in the list
func (list *ArrayList_string) Contains(item string) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *ArrayList_int) Fill(items []int) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_int8) Fill(items []int8) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_int16) Fill(items []int16) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_int32) Fill(items []int32) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_int64) Fill(items []int64) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_uint) Fill(items []uint) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_uint8) Fill(items []uint8) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_uint16) Fill(items []uint16) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_uint32) Fill(items []uint32) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_uint64) Fill(items []uint64) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_float32) Fill(items []float32) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_float64) Fill(items []float64) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_complex64) Fill(items []complex64) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_complex128) Fill(items []complex128) {
	list.items = append(list.items, items...)
}

// Fill adds all the items to the end of the list
func (list *ArrayList_string) Fill(items []string) {
	list.items = append(list.items, items...)
}

// Get returns the item at the index in the list
func (list *ArrayList_int) Get(idx int) int {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_int8) Get(idx int) int8 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_int16) Get(idx int) int16 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_int32) Get(idx int) int32 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_int64) Get(idx int) int64 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_uint) Get(idx int) uint {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_uint8) Get(idx int) uint8 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_uint16) Get(idx int) uint16 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_uint32) Get(idx int) uint32 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_uint64) Get(idx int) uint64 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_float32) Get(idx int) float32 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_float64) Get(idx int) float64 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_complex64) Get(idx int) complex64 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_complex128) Get(idx int) complex128 {
	return list.items[idx]
}

// Get returns the item at the index in the list
func (list *ArrayList_string) Get(idx int) string {
	return list.items[idx]
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_int) IndexOf(item int) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_int8) IndexOf(item int8) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_int16) IndexOf(item int16) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_int32) IndexOf(item int32) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_int64) IndexOf(item int64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_uint) IndexOf(item uint) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_uint8) IndexOf(item uint8) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_uint16) IndexOf(item uint16) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_uint32) IndexOf(item uint32) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_uint64) IndexOf(item uint64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_float32) IndexOf(item float32) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_float64) IndexOf(item float64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_complex64) IndexOf(item complex64) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_complex128) IndexOf(item complex128) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// IndexOf searches for the specified item and returns the
// index of the first occurence, or -1 if not found
func (list *ArrayList_string) IndexOf(item string) int {
	for i, t := range list.items {
		if t == item {
			return i
		}
	}
	return -1
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_int) Insert(idx int, item int) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_int8) Insert(idx int, item int8) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_int16) Insert(idx int, item int16) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_int32) Insert(idx int, item int32) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_int64) Insert(idx int, item int64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_uint) Insert(idx int, item uint) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_uint8) Insert(idx int, item uint8) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_uint16) Insert(idx int, item uint16) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_uint32) Insert(idx int, item uint32) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_uint64) Insert(idx int, item uint64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_float32) Insert(idx int, item float32) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_float64) Insert(idx int, item float64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_complex64) Insert(idx int, item complex64) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_complex128) Insert(idx int, item complex128) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Insert inserts the item into the list at the specified index
func (list *ArrayList_string) Insert(idx int, item string) {
	if idx > len(list.items) {
		list.Truncate(idx + 1)
	}
	list.items[idx] = item
}

// Len returns the number of items in the list
func (list *ArrayList_int) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_int8) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_int16) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_int32) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_int64) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_uint) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_uint8) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_uint16) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_uint32) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_uint64) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_float32) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_float64) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_complex64) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_complex128) Len() int {
	return len(list.items)
}

// Len returns the number of items in the list
func (list *ArrayList_string) Len() int {
	return len(list.items)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_int) Remove(item int) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_int8) Remove(item int8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_int16) Remove(item int16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_int32) Remove(item int32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_int64) Remove(item int64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_uint) Remove(item uint) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_uint8) Remove(item uint8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_uint16) Remove(item uint16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_uint32) Remove(item uint32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_uint64) Remove(item uint64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_float32) Remove(item float32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_float64) Remove(item float64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_complex64) Remove(item complex64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_complex128) Remove(item complex128) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the first occurence of the item from the list
func (list *ArrayList_string) Remove(item string) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_int) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_int8) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_int16) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_int32) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_int64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_uint) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_uint8) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_uint16) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_uint32) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_uint64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_float32) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_float64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_complex64) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_complex128) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// RemoveAt removes the element at the specified index of the list
func (list *ArrayList_string) RemoveAt(idx int) {
	list.items = append(list.items[:idx], list.items[idx+1:]...)

	if cap(list.items) > 2*len(list.items) {
		list.Truncate(len(list.items))
	}
}

// Reverse reverses the list
func (list *ArrayList_int) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_int8) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_int16) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_int32) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_int64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_uint) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_uint8) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_uint16) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_uint32) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_uint64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_float32) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_float64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_complex64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_complex128) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *ArrayList_string) Reverse() {
	Reverse(list)
}

// Swap swaps items in the list
func (list *ArrayList_int) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_int8) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_int16) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_int32) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_int64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_uint) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_uint8) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_uint16) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_uint32) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_uint64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_float32) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_float64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_complex64) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_complex128) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Swap swaps items in the list
func (list *ArrayList_string) Swap(i, j int) {
	list.items[i], list.items[j] = list.items[j], list.items[i]
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_int) Truncate(sz int) {
	items := make([]int, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_int8) Truncate(sz int) {
	items := make([]int8, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_int16) Truncate(sz int) {
	items := make([]int16, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_int32) Truncate(sz int) {
	items := make([]int32, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_int64) Truncate(sz int) {
	items := make([]int64, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_uint) Truncate(sz int) {
	items := make([]uint, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_uint8) Truncate(sz int) {
	items := make([]uint8, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_uint16) Truncate(sz int) {
	items := make([]uint16, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_uint32) Truncate(sz int) {
	items := make([]uint32, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_uint64) Truncate(sz int) {
	items := make([]uint64, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_float32) Truncate(sz int) {
	items := make([]float32, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_float64) Truncate(sz int) {
	items := make([]float64, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_complex64) Truncate(sz int) {
	items := make([]complex64, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_complex128) Truncate(sz int) {
	items := make([]complex128, sz)
	copy(items, list.items)
	list.items = items
}

// Truncate changes the size of the list to sz, any excess items are removed
func (list *ArrayList_string) Truncate(sz int) {
	items := make([]string, sz)
	copy(items, list.items)
	list.items = items
}
