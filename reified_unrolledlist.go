package container

type UnrolledList_int struct {
	items     [][]int
	size      int
	arraySize int
}

type UnrolledList_int8 struct {
	items     [][]int8
	size      int
	arraySize int
}

type UnrolledList_int16 struct {
	items     [][]int16
	size      int
	arraySize int
}

type UnrolledList_int32 struct {
	items     [][]int32
	size      int
	arraySize int
}

type UnrolledList_int64 struct {
	items     [][]int64
	size      int
	arraySize int
}

type UnrolledList_uint struct {
	items     [][]uint
	size      int
	arraySize int
}

type UnrolledList_uint8 struct {
	items     [][]uint8
	size      int
	arraySize int
}

type UnrolledList_uint16 struct {
	items     [][]uint16
	size      int
	arraySize int
}

type UnrolledList_uint32 struct {
	items     [][]uint32
	size      int
	arraySize int
}

type UnrolledList_uint64 struct {
	items     [][]uint64
	size      int
	arraySize int
}

type UnrolledList_float32 struct {
	items     [][]float32
	size      int
	arraySize int
}

type UnrolledList_float64 struct {
	items     [][]float64
	size      int
	arraySize int
}

type UnrolledList_complex64 struct {
	items     [][]complex64
	size      int
	arraySize int
}

type UnrolledList_complex128 struct {
	items     [][]complex128
	size      int
	arraySize int
}

type UnrolledList_string struct {
	items     [][]string
	size      int
	arraySize int
}

// NewUnrolledList_int creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_int(arraySize int) *UnrolledList_int {
	return &UnrolledList_int{
		arraySize: arraySize,
	}
}

// NewUnrolledList_int8 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_int8(arraySize int) *UnrolledList_int8 {
	return &UnrolledList_int8{
		arraySize: arraySize,
	}
}

// NewUnrolledList_int16 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_int16(arraySize int) *UnrolledList_int16 {
	return &UnrolledList_int16{
		arraySize: arraySize,
	}
}

// NewUnrolledList_int32 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_int32(arraySize int) *UnrolledList_int32 {
	return &UnrolledList_int32{
		arraySize: arraySize,
	}
}

// NewUnrolledList_int64 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_int64(arraySize int) *UnrolledList_int64 {
	return &UnrolledList_int64{
		arraySize: arraySize,
	}
}

// NewUnrolledList_uint creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_uint(arraySize int) *UnrolledList_uint {
	return &UnrolledList_uint{
		arraySize: arraySize,
	}
}

// NewUnrolledList_uint8 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_uint8(arraySize int) *UnrolledList_uint8 {
	return &UnrolledList_uint8{
		arraySize: arraySize,
	}
}

// NewUnrolledList_uint16 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_uint16(arraySize int) *UnrolledList_uint16 {
	return &UnrolledList_uint16{
		arraySize: arraySize,
	}
}

// NewUnrolledList_uint32 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_uint32(arraySize int) *UnrolledList_uint32 {
	return &UnrolledList_uint32{
		arraySize: arraySize,
	}
}

// NewUnrolledList_uint64 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_uint64(arraySize int) *UnrolledList_uint64 {
	return &UnrolledList_uint64{
		arraySize: arraySize,
	}
}

// NewUnrolledList_float32 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_float32(arraySize int) *UnrolledList_float32 {
	return &UnrolledList_float32{
		arraySize: arraySize,
	}
}

// NewUnrolledList_float64 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_float64(arraySize int) *UnrolledList_float64 {
	return &UnrolledList_float64{
		arraySize: arraySize,
	}
}

// NewUnrolledList_complex64 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_complex64(arraySize int) *UnrolledList_complex64 {
	return &UnrolledList_complex64{
		arraySize: arraySize,
	}
}

// NewUnrolledList_complex128 creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_complex128(arraySize int) *UnrolledList_complex128 {
	return &UnrolledList_complex128{
		arraySize: arraySize,
	}
}

// NewUnrolledList_string creates a new UnrolledList. arraySize defines how many
// items are added to each buffer
func NewUnrolledList_string(arraySize int) *UnrolledList_string {
	return &UnrolledList_string{
		arraySize: arraySize,
	}
}

// Add adds item to the list
func (list *UnrolledList_int) Add(item int) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_int8) Add(item int8) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_int16) Add(item int16) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_int32) Add(item int32) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_int64) Add(item int64) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_uint) Add(item uint) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_uint8) Add(item uint8) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_uint16) Add(item uint16) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_uint32) Add(item uint32) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_uint64) Add(item uint64) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_float32) Add(item float32) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_float64) Add(item float64) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_complex64) Add(item complex64) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_complex128) Add(item complex128) {
	list.Fill(item)
}

// Add adds item to the list
func (list *UnrolledList_string) Add(item string) {
	list.Fill(item)
}

// Clear removes all the items from the list
func (list *UnrolledList_int) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_int8) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_int16) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_int32) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_int64) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_uint) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_uint8) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_uint16) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_uint32) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_uint64) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_float32) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_float64) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_complex64) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_complex128) Clear() {
	list.items = nil
	list.size = 0
}

// Clear removes all the items from the list
func (list *UnrolledList_string) Clear() {
	list.items = nil
	list.size = 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_int) Contains(item int) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_int8) Contains(item int8) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_int16) Contains(item int16) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_int32) Contains(item int32) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_int64) Contains(item int64) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_uint) Contains(item uint) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_uint8) Contains(item uint8) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_uint16) Contains(item uint16) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_uint32) Contains(item uint32) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_uint64) Contains(item uint64) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_float32) Contains(item float32) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_float64) Contains(item float64) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_complex64) Contains(item complex64) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_complex128) Contains(item complex128) bool {
	return list.IndexOf(item) >= 0
}

// Contains returns whether or not the list contains the item
func (list *UnrolledList_string) Contains(item string) bool {
	return list.IndexOf(item) >= 0
}

// Fill adds all the items to the end of the list
func (list *UnrolledList_int) Fill(items ...int) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_int8) Fill(items ...int8) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_int16) Fill(items ...int16) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_int32) Fill(items ...int32) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_int64) Fill(items ...int64) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_uint) Fill(items ...uint) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_uint8) Fill(items ...uint8) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_uint16) Fill(items ...uint16) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_uint32) Fill(items ...uint32) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_uint64) Fill(items ...uint64) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_float32) Fill(items ...float32) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_float64) Fill(items ...float64) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_complex64) Fill(items ...complex64) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_complex128) Fill(items ...complex128) {
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

// Fill adds all the items to the end of the list
func (list *UnrolledList_string) Fill(items ...string) {
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
func (list *UnrolledList_int) Get(idx int) int {
	if idx > list.size {
		var def int
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_int8) Get(idx int) int8 {
	if idx > list.size {
		var def int8
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_int16) Get(idx int) int16 {
	if idx > list.size {
		var def int16
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_int32) Get(idx int) int32 {
	if idx > list.size {
		var def int32
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_int64) Get(idx int) int64 {
	if idx > list.size {
		var def int64
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_uint) Get(idx int) uint {
	if idx > list.size {
		var def uint
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_uint8) Get(idx int) uint8 {
	if idx > list.size {
		var def uint8
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_uint16) Get(idx int) uint16 {
	if idx > list.size {
		var def uint16
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_uint32) Get(idx int) uint32 {
	if idx > list.size {
		var def uint32
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_uint64) Get(idx int) uint64 {
	if idx > list.size {
		var def uint64
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_float32) Get(idx int) float32 {
	if idx > list.size {
		var def float32
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_float64) Get(idx int) float64 {
	if idx > list.size {
		var def float64
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_complex64) Get(idx int) complex64 {
	if idx > list.size {
		var def complex64
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_complex128) Get(idx int) complex128 {
	if idx > list.size {
		var def complex128
		return def
	}
	return list.get(idx)
}

// Get returns the item at the index in the list
func (list *UnrolledList_string) Get(idx int) string {
	if idx > list.size {
		var def string
		return def
	}
	return list.get(idx)
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_int) IndexOf(item int) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_int8) IndexOf(item int8) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_int16) IndexOf(item int16) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_int32) IndexOf(item int32) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_int64) IndexOf(item int64) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_uint) IndexOf(item uint) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_uint8) IndexOf(item uint8) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_uint16) IndexOf(item uint16) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_uint32) IndexOf(item uint32) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_uint64) IndexOf(item uint64) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_float32) IndexOf(item float32) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_float64) IndexOf(item float64) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_complex64) IndexOf(item complex64) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_complex128) IndexOf(item complex128) int {
	for i, items := range list.items {
		for j, t := range items {
			if t == item {
				return i*list.arraySize + j
			}
		}
	}
	return -1
}

// IndexOf returns the index of the item in the list. If the item doesn't
// exist in the list -1 is returned
func (list *UnrolledList_string) IndexOf(item string) int {
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
func (list *UnrolledList_int) Insert(idx int, item int) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_int8) Insert(idx int, item int8) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_int16) Insert(idx int, item int16) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_int32) Insert(idx int, item int32) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_int64) Insert(idx int, item int64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_uint) Insert(idx int, item uint) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_uint8) Insert(idx int, item uint8) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_uint16) Insert(idx int, item uint16) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_uint32) Insert(idx int, item uint32) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_uint64) Insert(idx int, item uint64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_float32) Insert(idx int, item float32) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_float64) Insert(idx int, item float64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_complex64) Insert(idx int, item complex64) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_complex128) Insert(idx int, item complex128) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Insert inserts the item into the list
func (list *UnrolledList_string) Insert(idx int, item string) {
	list.grow(idx + 1)
	for i := idx + 1; i < list.size-1; i++ {
		list.set(i+1, list.get(i))
	}
	list.set(idx, item)
}

// Len returns the number of items in the list
func (list *UnrolledList_int) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_int8) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_int16) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_int32) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_int64) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_uint) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_uint8) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_uint16) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_uint32) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_uint64) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_float32) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_float64) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_complex64) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_complex128) Len() int {
	return list.size
}

// Len returns the number of items in the list
func (list *UnrolledList_string) Len() int {
	return list.size
}

// Remove removes the item from the list
func (list *UnrolledList_int) Remove(item int) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_int8) Remove(item int8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_int16) Remove(item int16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_int32) Remove(item int32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_int64) Remove(item int64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_uint) Remove(item uint) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_uint8) Remove(item uint8) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_uint16) Remove(item uint16) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_uint32) Remove(item uint32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_uint64) Remove(item uint64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_float32) Remove(item float32) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_float64) Remove(item float64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_complex64) Remove(item complex64) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_complex128) Remove(item complex128) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// Remove removes the item from the list
func (list *UnrolledList_string) Remove(item string) {
	idx := list.IndexOf(item)
	if idx < 0 {
		return
	}
	list.RemoveAt(idx)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_int) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_int8) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_int16) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_int32) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_int64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_uint) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_uint8) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_uint16) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_uint32) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_uint64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_float32) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_float64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_complex64) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_complex128) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// RemoveAt removes the item at the index from the list
func (list *UnrolledList_string) RemoveAt(idx int) {
	for i := idx; i < list.size-1; i++ {
		list.set(i, list.get(i+1))
	}
	list.shrink(list.size - 1)
}

// Reverse reverses the list
func (list *UnrolledList_int) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_int8) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_int16) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_int32) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_int64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_uint) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_uint8) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_uint16) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_uint32) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_uint64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_float32) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_float64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_complex64) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_complex128) Reverse() {
	Reverse(list)
}

// Reverse reverses the list
func (list *UnrolledList_string) Reverse() {
	Reverse(list)
}

// Set sets the item at the index in the list
func (list *UnrolledList_int) Set(idx int, item int) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_int8) Set(idx int, item int8) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_int16) Set(idx int, item int16) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_int32) Set(idx int, item int32) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_int64) Set(idx int, item int64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_uint) Set(idx int, item uint) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_uint8) Set(idx int, item uint8) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_uint16) Set(idx int, item uint16) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_uint32) Set(idx int, item uint32) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_uint64) Set(idx int, item uint64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_float32) Set(idx int, item float32) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_float64) Set(idx int, item float64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_complex64) Set(idx int, item complex64) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_complex128) Set(idx int, item complex128) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Set sets the item at the index in the list
func (list *UnrolledList_string) Set(idx int, item string) {
	list.grow(idx + 1)
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

// Swap swaps i and j
func (list *UnrolledList_int) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_int8) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_int16) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_int32) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_int64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_uint) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_uint8) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_uint16) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_uint32) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_uint64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_float32) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_float64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_complex64) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_complex128) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

// Swap swaps i and j
func (list *UnrolledList_string) Swap(i, j int) {
	list.items[i/list.arraySize][i%list.arraySize] = list.items[j/list.arraySize][j%list.arraySize]
}

func (list *UnrolledList_int) get(idx int) int {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_int8) get(idx int) int8 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_int16) get(idx int) int16 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_int32) get(idx int) int32 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_int64) get(idx int) int64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_uint) get(idx int) uint {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_uint8) get(idx int) uint8 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_uint16) get(idx int) uint16 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_uint32) get(idx int) uint32 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_uint64) get(idx int) uint64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_float32) get(idx int) float32 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_float64) get(idx int) float64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_complex64) get(idx int) complex64 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_complex128) get(idx int) complex128 {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_string) get(idx int) string {
	return list.items[idx/list.arraySize][idx%list.arraySize]
}

func (list *UnrolledList_int) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_int8) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int8, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_int16) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int16, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_int32) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int32, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_int64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]int64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_uint) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_uint8) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint8, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_uint16) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint16, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_uint32) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint32, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_uint64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]uint64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_float32) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]float32, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_float64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]float64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_complex64) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]complex64, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_complex128) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]complex128, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_string) grow(size int) {
	if list.size > size {
		return
	}

	idx := size / list.arraySize
	for idx >= len(list.items) {
		list.items = append(list.items, make([]string, list.arraySize))
	}

	list.size = size
}

func (list *UnrolledList_int) set(idx int, item int) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_int8) set(idx int, item int8) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_int16) set(idx int, item int16) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_int32) set(idx int, item int32) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_int64) set(idx int, item int64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_uint) set(idx int, item uint) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_uint8) set(idx int, item uint8) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_uint16) set(idx int, item uint16) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_uint32) set(idx int, item uint32) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_uint64) set(idx int, item uint64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_float32) set(idx int, item float32) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_float64) set(idx int, item float64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_complex64) set(idx int, item complex64) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_complex128) set(idx int, item complex128) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_string) set(idx int, item string) {
	list.items[idx/list.arraySize][idx%list.arraySize] = item
}

func (list *UnrolledList_int) shrink(size int) {
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

func (list *UnrolledList_int8) shrink(size int) {
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

func (list *UnrolledList_int16) shrink(size int) {
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

func (list *UnrolledList_int32) shrink(size int) {
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

func (list *UnrolledList_int64) shrink(size int) {
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

func (list *UnrolledList_uint) shrink(size int) {
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

func (list *UnrolledList_uint8) shrink(size int) {
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

func (list *UnrolledList_uint16) shrink(size int) {
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

func (list *UnrolledList_uint32) shrink(size int) {
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

func (list *UnrolledList_uint64) shrink(size int) {
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

func (list *UnrolledList_float32) shrink(size int) {
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

func (list *UnrolledList_float64) shrink(size int) {
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

func (list *UnrolledList_complex64) shrink(size int) {
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

func (list *UnrolledList_complex128) shrink(size int) {
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

func (list *UnrolledList_string) shrink(size int) {
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
