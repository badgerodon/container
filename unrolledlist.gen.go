package container
import "bytes"



// An unrolled list is a linked-list of arrays
type UnrolledListOfUint8 struct {
	items     [][]uint8
	size      int
	arraySize int
}

// NewUnrolledListOfUint8 creates a new UnrolledListOfUint8. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfUint8(arraySize int) *UnrolledListOfUint8 {
	return &UnrolledListOfUint8{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfUint8) Add(item uint8) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfUint8) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfUint8) Contains(item uint8) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfUint8) Fill(items ...uint8) {
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
func (list *UnrolledListOfUint8) Get(idx int) uint8 {
	if idx > list.size {
		var def uint8
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfUint8) IndexOf(item uint8) int {
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
func (list *UnrolledListOfUint8) Insert(idx int, item uint8) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfUint8) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfUint8) Remove(item uint8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfUint8) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfUint8) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfUint8) Set(idx int, item uint8) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfUint8) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfUint8) get(idx int) uint8 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfUint8) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint8, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfUint8) set(idx int, item uint8) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfUint8) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfUint16 struct {
	items     [][]uint16
	size      int
	arraySize int
}

// NewUnrolledListOfUint16 creates a new UnrolledListOfUint16. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfUint16(arraySize int) *UnrolledListOfUint16 {
	return &UnrolledListOfUint16{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfUint16) Add(item uint16) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfUint16) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfUint16) Contains(item uint16) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfUint16) Fill(items ...uint16) {
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
func (list *UnrolledListOfUint16) Get(idx int) uint16 {
	if idx > list.size {
		var def uint16
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfUint16) IndexOf(item uint16) int {
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
func (list *UnrolledListOfUint16) Insert(idx int, item uint16) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfUint16) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfUint16) Remove(item uint16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfUint16) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfUint16) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfUint16) Set(idx int, item uint16) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfUint16) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfUint16) get(idx int) uint16 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfUint16) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint16, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfUint16) set(idx int, item uint16) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfUint16) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfUint32 struct {
	items     [][]uint32
	size      int
	arraySize int
}

// NewUnrolledListOfUint32 creates a new UnrolledListOfUint32. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfUint32(arraySize int) *UnrolledListOfUint32 {
	return &UnrolledListOfUint32{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfUint32) Add(item uint32) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfUint32) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfUint32) Contains(item uint32) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfUint32) Fill(items ...uint32) {
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
func (list *UnrolledListOfUint32) Get(idx int) uint32 {
	if idx > list.size {
		var def uint32
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfUint32) IndexOf(item uint32) int {
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
func (list *UnrolledListOfUint32) Insert(idx int, item uint32) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfUint32) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfUint32) Remove(item uint32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfUint32) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfUint32) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfUint32) Set(idx int, item uint32) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfUint32) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfUint32) get(idx int) uint32 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfUint32) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint32, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfUint32) set(idx int, item uint32) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfUint32) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfUint64 struct {
	items     [][]uint64
	size      int
	arraySize int
}

// NewUnrolledListOfUint64 creates a new UnrolledListOfUint64. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfUint64(arraySize int) *UnrolledListOfUint64 {
	return &UnrolledListOfUint64{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfUint64) Add(item uint64) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfUint64) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfUint64) Contains(item uint64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfUint64) Fill(items ...uint64) {
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
func (list *UnrolledListOfUint64) Get(idx int) uint64 {
	if idx > list.size {
		var def uint64
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfUint64) IndexOf(item uint64) int {
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
func (list *UnrolledListOfUint64) Insert(idx int, item uint64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfUint64) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfUint64) Remove(item uint64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfUint64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfUint64) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfUint64) Set(idx int, item uint64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfUint64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfUint64) get(idx int) uint64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfUint64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfUint64) set(idx int, item uint64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfUint64) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfInt8 struct {
	items     [][]int8
	size      int
	arraySize int
}

// NewUnrolledListOfInt8 creates a new UnrolledListOfInt8. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfInt8(arraySize int) *UnrolledListOfInt8 {
	return &UnrolledListOfInt8{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfInt8) Add(item int8) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfInt8) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfInt8) Contains(item int8) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfInt8) Fill(items ...int8) {
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
func (list *UnrolledListOfInt8) Get(idx int) int8 {
	if idx > list.size {
		var def int8
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfInt8) IndexOf(item int8) int {
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
func (list *UnrolledListOfInt8) Insert(idx int, item int8) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfInt8) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfInt8) Remove(item int8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfInt8) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfInt8) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfInt8) Set(idx int, item int8) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfInt8) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfInt8) get(idx int) int8 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfInt8) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int8, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfInt8) set(idx int, item int8) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfInt8) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfInt16 struct {
	items     [][]int16
	size      int
	arraySize int
}

// NewUnrolledListOfInt16 creates a new UnrolledListOfInt16. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfInt16(arraySize int) *UnrolledListOfInt16 {
	return &UnrolledListOfInt16{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfInt16) Add(item int16) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfInt16) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfInt16) Contains(item int16) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfInt16) Fill(items ...int16) {
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
func (list *UnrolledListOfInt16) Get(idx int) int16 {
	if idx > list.size {
		var def int16
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfInt16) IndexOf(item int16) int {
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
func (list *UnrolledListOfInt16) Insert(idx int, item int16) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfInt16) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfInt16) Remove(item int16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfInt16) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfInt16) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfInt16) Set(idx int, item int16) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfInt16) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfInt16) get(idx int) int16 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfInt16) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int16, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfInt16) set(idx int, item int16) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfInt16) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfInt32 struct {
	items     [][]int32
	size      int
	arraySize int
}

// NewUnrolledListOfInt32 creates a new UnrolledListOfInt32. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfInt32(arraySize int) *UnrolledListOfInt32 {
	return &UnrolledListOfInt32{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfInt32) Add(item int32) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfInt32) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfInt32) Contains(item int32) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfInt32) Fill(items ...int32) {
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
func (list *UnrolledListOfInt32) Get(idx int) int32 {
	if idx > list.size {
		var def int32
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfInt32) IndexOf(item int32) int {
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
func (list *UnrolledListOfInt32) Insert(idx int, item int32) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfInt32) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfInt32) Remove(item int32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfInt32) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfInt32) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfInt32) Set(idx int, item int32) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfInt32) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfInt32) get(idx int) int32 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfInt32) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int32, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfInt32) set(idx int, item int32) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfInt32) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfInt64 struct {
	items     [][]int64
	size      int
	arraySize int
}

// NewUnrolledListOfInt64 creates a new UnrolledListOfInt64. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfInt64(arraySize int) *UnrolledListOfInt64 {
	return &UnrolledListOfInt64{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfInt64) Add(item int64) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfInt64) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfInt64) Contains(item int64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfInt64) Fill(items ...int64) {
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
func (list *UnrolledListOfInt64) Get(idx int) int64 {
	if idx > list.size {
		var def int64
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfInt64) IndexOf(item int64) int {
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
func (list *UnrolledListOfInt64) Insert(idx int, item int64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfInt64) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfInt64) Remove(item int64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfInt64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfInt64) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfInt64) Set(idx int, item int64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfInt64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfInt64) get(idx int) int64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfInt64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfInt64) set(idx int, item int64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfInt64) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfFloat32 struct {
	items     [][]float32
	size      int
	arraySize int
}

// NewUnrolledListOfFloat32 creates a new UnrolledListOfFloat32. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfFloat32(arraySize int) *UnrolledListOfFloat32 {
	return &UnrolledListOfFloat32{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfFloat32) Add(item float32) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfFloat32) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfFloat32) Contains(item float32) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfFloat32) Fill(items ...float32) {
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
func (list *UnrolledListOfFloat32) Get(idx int) float32 {
	if idx > list.size {
		var def float32
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfFloat32) IndexOf(item float32) int {
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
func (list *UnrolledListOfFloat32) Insert(idx int, item float32) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfFloat32) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfFloat32) Remove(item float32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfFloat32) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfFloat32) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfFloat32) Set(idx int, item float32) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfFloat32) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfFloat32) get(idx int) float32 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfFloat32) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]float32, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfFloat32) set(idx int, item float32) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfFloat32) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfFloat64 struct {
	items     [][]float64
	size      int
	arraySize int
}

// NewUnrolledListOfFloat64 creates a new UnrolledListOfFloat64. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfFloat64(arraySize int) *UnrolledListOfFloat64 {
	return &UnrolledListOfFloat64{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfFloat64) Add(item float64) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfFloat64) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfFloat64) Contains(item float64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfFloat64) Fill(items ...float64) {
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
func (list *UnrolledListOfFloat64) Get(idx int) float64 {
	if idx > list.size {
		var def float64
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfFloat64) IndexOf(item float64) int {
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
func (list *UnrolledListOfFloat64) Insert(idx int, item float64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfFloat64) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfFloat64) Remove(item float64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfFloat64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfFloat64) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfFloat64) Set(idx int, item float64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfFloat64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfFloat64) get(idx int) float64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfFloat64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]float64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfFloat64) set(idx int, item float64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfFloat64) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfComplex64 struct {
	items     [][]complex64
	size      int
	arraySize int
}

// NewUnrolledListOfComplex64 creates a new UnrolledListOfComplex64. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfComplex64(arraySize int) *UnrolledListOfComplex64 {
	return &UnrolledListOfComplex64{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfComplex64) Add(item complex64) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfComplex64) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfComplex64) Contains(item complex64) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfComplex64) Fill(items ...complex64) {
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
func (list *UnrolledListOfComplex64) Get(idx int) complex64 {
	if idx > list.size {
		var def complex64
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfComplex64) IndexOf(item complex64) int {
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
func (list *UnrolledListOfComplex64) Insert(idx int, item complex64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfComplex64) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfComplex64) Remove(item complex64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfComplex64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfComplex64) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfComplex64) Set(idx int, item complex64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfComplex64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfComplex64) get(idx int) complex64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfComplex64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]complex64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfComplex64) set(idx int, item complex64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfComplex64) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfComplex128 struct {
	items     [][]complex128
	size      int
	arraySize int
}

// NewUnrolledListOfComplex128 creates a new UnrolledListOfComplex128. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfComplex128(arraySize int) *UnrolledListOfComplex128 {
	return &UnrolledListOfComplex128{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfComplex128) Add(item complex128) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfComplex128) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfComplex128) Contains(item complex128) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfComplex128) Fill(items ...complex128) {
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
func (list *UnrolledListOfComplex128) Get(idx int) complex128 {
	if idx > list.size {
		var def complex128
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfComplex128) IndexOf(item complex128) int {
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
func (list *UnrolledListOfComplex128) Insert(idx int, item complex128) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfComplex128) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfComplex128) Remove(item complex128) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfComplex128) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfComplex128) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfComplex128) Set(idx int, item complex128) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfComplex128) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfComplex128) get(idx int) complex128 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfComplex128) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]complex128, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfComplex128) set(idx int, item complex128) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfComplex128) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfByte struct {
	items     [][]byte
	size      int
	arraySize int
}

// NewUnrolledListOfByte creates a new UnrolledListOfByte. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfByte(arraySize int) *UnrolledListOfByte {
	return &UnrolledListOfByte{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfByte) Add(item byte) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfByte) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfByte) Contains(item byte) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfByte) Fill(items ...byte) {
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
func (list *UnrolledListOfByte) Get(idx int) byte {
	if idx > list.size {
		var def byte
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfByte) IndexOf(item byte) int {
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
func (list *UnrolledListOfByte) Insert(idx int, item byte) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfByte) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfByte) Remove(item byte) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfByte) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfByte) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfByte) Set(idx int, item byte) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfByte) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfByte) get(idx int) byte {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfByte) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]byte, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfByte) set(idx int, item byte) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfByte) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfRune struct {
	items     [][]rune
	size      int
	arraySize int
}

// NewUnrolledListOfRune creates a new UnrolledListOfRune. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfRune(arraySize int) *UnrolledListOfRune {
	return &UnrolledListOfRune{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfRune) Add(item rune) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfRune) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfRune) Contains(item rune) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfRune) Fill(items ...rune) {
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
func (list *UnrolledListOfRune) Get(idx int) rune {
	if idx > list.size {
		var def rune
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfRune) IndexOf(item rune) int {
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
func (list *UnrolledListOfRune) Insert(idx int, item rune) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfRune) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfRune) Remove(item rune) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfRune) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfRune) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfRune) Set(idx int, item rune) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfRune) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfRune) get(idx int) rune {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfRune) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]rune, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfRune) set(idx int, item rune) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfRune) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfUint struct {
	items     [][]uint
	size      int
	arraySize int
}

// NewUnrolledListOfUint creates a new UnrolledListOfUint. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfUint(arraySize int) *UnrolledListOfUint {
	return &UnrolledListOfUint{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfUint) Add(item uint) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfUint) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfUint) Contains(item uint) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfUint) Fill(items ...uint) {
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
func (list *UnrolledListOfUint) Get(idx int) uint {
	if idx > list.size {
		var def uint
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfUint) IndexOf(item uint) int {
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
func (list *UnrolledListOfUint) Insert(idx int, item uint) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfUint) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfUint) Remove(item uint) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfUint) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfUint) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfUint) Set(idx int, item uint) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfUint) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfUint) get(idx int) uint {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfUint) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfUint) set(idx int, item uint) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfUint) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfInt struct {
	items     [][]int
	size      int
	arraySize int
}

// NewUnrolledListOfInt creates a new UnrolledListOfInt. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfInt(arraySize int) *UnrolledListOfInt {
	return &UnrolledListOfInt{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfInt) Add(item int) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfInt) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfInt) Contains(item int) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfInt) Fill(items ...int) {
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
func (list *UnrolledListOfInt) Get(idx int) int {
	if idx > list.size {
		var def int
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfInt) IndexOf(item int) int {
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
func (list *UnrolledListOfInt) Insert(idx int, item int) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfInt) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfInt) Remove(item int) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfInt) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfInt) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfInt) Set(idx int, item int) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfInt) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfInt) get(idx int) int {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfInt) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfInt) set(idx int, item int) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfInt) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfUintptr struct {
	items     [][]uintptr
	size      int
	arraySize int
}

// NewUnrolledListOfUintptr creates a new UnrolledListOfUintptr. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfUintptr(arraySize int) *UnrolledListOfUintptr {
	return &UnrolledListOfUintptr{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfUintptr) Add(item uintptr) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfUintptr) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfUintptr) Contains(item uintptr) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfUintptr) Fill(items ...uintptr) {
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
func (list *UnrolledListOfUintptr) Get(idx int) uintptr {
	if idx > list.size {
		var def uintptr
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfUintptr) IndexOf(item uintptr) int {
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
func (list *UnrolledListOfUintptr) Insert(idx int, item uintptr) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfUintptr) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfUintptr) Remove(item uintptr) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfUintptr) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfUintptr) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfUintptr) Set(idx int, item uintptr) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfUintptr) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfUintptr) get(idx int) uintptr {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfUintptr) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uintptr, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfUintptr) set(idx int, item uintptr) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfUintptr) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfString struct {
	items     [][]string
	size      int
	arraySize int
}

// NewUnrolledListOfString creates a new UnrolledListOfString. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfString(arraySize int) *UnrolledListOfString {
	return &UnrolledListOfString{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfString) Add(item string) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfString) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfString) Contains(item string) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfString) Fill(items ...string) {
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
func (list *UnrolledListOfString) Get(idx int) string {
	if idx > list.size {
		var def string
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfString) IndexOf(item string) int {
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
func (list *UnrolledListOfString) Insert(idx int, item string) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfString) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfString) Remove(item string) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfString) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfString) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfString) Set(idx int, item string) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfString) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfString) get(idx int) string {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfString) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]string, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfString) set(idx int, item string) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfString) shrink(size int) {
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



// An unrolled list is a linked-list of arrays
type UnrolledListOfBytes struct {
	items     [][][]byte
	size      int
	arraySize int
}

// NewUnrolledListOfBytes creates a new UnrolledListOfBytes. arraySize defines how many
// items are added to each buffer
func NewUnrolledListOfBytes(arraySize int) *UnrolledListOfBytes {
	return &UnrolledListOfBytes{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledListOfBytes) Add(item []byte) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledListOfBytes) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledListOfBytes) Contains(item []byte) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledListOfBytes) Fill(items ...[]byte) {
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
func (list *UnrolledListOfBytes) Get(idx int) []byte {
	if idx > list.size {
		var def []byte
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledListOfBytes) IndexOf(item []byte) int {
	for i, items := range list.items {
		for j, t := range items {
			if bytes.Equal(t, item) {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// Insert inserts the item into the list
func (list *UnrolledListOfBytes) Insert(idx int, item []byte) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledListOfBytes) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledListOfBytes) Remove(item []byte) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledListOfBytes) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledListOfBytes) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledListOfBytes) Set(idx int, item []byte) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledListOfBytes) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledListOfBytes) get(idx int) []byte {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledListOfBytes) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([][]byte, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledListOfBytes) set(idx int, item []byte) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledListOfBytes) shrink(size int) {
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
