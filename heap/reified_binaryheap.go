package heap

type

// A BinaryHeapInt implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapInt struct {
	less     func(x, y int) bool
	elements []int
}

type

// A BinaryHeapInt8 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapInt8 struct {
	less     func(x, y int8) bool
	elements []int8
}

type

// A BinaryHeapInt16 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapInt16 struct {
	less     func(x, y int16) bool
	elements []int16
}

type

// A BinaryHeapInt32 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapInt32 struct {
	less     func(x, y int32) bool
	elements []int32
}

type

// A BinaryHeapInt64 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapInt64 struct {
	less     func(x, y int64) bool
	elements []int64
}

type

// A BinaryHeapUint implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapUint struct {
	less     func(x, y uint) bool
	elements []uint
}

type

// A BinaryHeapUint8 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapUint8 struct {
	less     func(x, y uint8) bool
	elements []uint8
}

type

// A BinaryHeapUint16 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapUint16 struct {
	less     func(x, y uint16) bool
	elements []uint16
}

type

// A BinaryHeapUint32 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapUint32 struct {
	less     func(x, y uint32) bool
	elements []uint32
}

type

// A BinaryHeapUint64 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapUint64 struct {
	less     func(x, y uint64) bool
	elements []uint64
}

type

// A BinaryHeapFloat32 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapFloat32 struct {
	less     func(x, y float32) bool
	elements []float32
}

type

// A BinaryHeapFloat64 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapFloat64 struct {
	less     func(x, y float64) bool
	elements []float64
}

type

// A BinaryHeapComplex64 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapComplex64 struct {
	less     func(x, y complex64) bool
	elements []complex64
}

type

// A BinaryHeapComplex128 implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapComplex128 struct {
	less     func(x, y complex128) bool
	elements []complex128
}

type

// A BinaryHeapString implements the binary heap data structure
// see: wikipedia.org/wiki/Binary_heap
BinaryHeapString struct {
	less     func(x, y string) bool
	elements []string
}

// NewBinaryHeapInt creates a new BinaryHeap
func NewBinaryHeapInt(less func(x, y int) bool) *BinaryHeapInt {
	return &BinaryHeapInt{
		less: less,
	}
}

// NewBinaryHeapInt8 creates a new BinaryHeap
func NewBinaryHeapInt8(less func(x, y int8) bool) *BinaryHeapInt8 {
	return &BinaryHeapInt8{
		less: less,
	}
}

// NewBinaryHeapInt16 creates a new BinaryHeap
func NewBinaryHeapInt16(less func(x, y int16) bool) *BinaryHeapInt16 {
	return &BinaryHeapInt16{
		less: less,
	}
}

// NewBinaryHeapInt32 creates a new BinaryHeap
func NewBinaryHeapInt32(less func(x, y int32) bool) *BinaryHeapInt32 {
	return &BinaryHeapInt32{
		less: less,
	}
}

// NewBinaryHeapInt64 creates a new BinaryHeap
func NewBinaryHeapInt64(less func(x, y int64) bool) *BinaryHeapInt64 {
	return &BinaryHeapInt64{
		less: less,
	}
}

// NewBinaryHeapUint creates a new BinaryHeap
func NewBinaryHeapUint(less func(x, y uint) bool) *BinaryHeapUint {
	return &BinaryHeapUint{
		less: less,
	}
}

// NewBinaryHeapUint8 creates a new BinaryHeap
func NewBinaryHeapUint8(less func(x, y uint8) bool) *BinaryHeapUint8 {
	return &BinaryHeapUint8{
		less: less,
	}
}

// NewBinaryHeapUint16 creates a new BinaryHeap
func NewBinaryHeapUint16(less func(x, y uint16) bool) *BinaryHeapUint16 {
	return &BinaryHeapUint16{
		less: less,
	}
}

// NewBinaryHeapUint32 creates a new BinaryHeap
func NewBinaryHeapUint32(less func(x, y uint32) bool) *BinaryHeapUint32 {
	return &BinaryHeapUint32{
		less: less,
	}
}

// NewBinaryHeapUint64 creates a new BinaryHeap
func NewBinaryHeapUint64(less func(x, y uint64) bool) *BinaryHeapUint64 {
	return &BinaryHeapUint64{
		less: less,
	}
}

// NewBinaryHeapFloat32 creates a new BinaryHeap
func NewBinaryHeapFloat32(less func(x, y float32) bool) *BinaryHeapFloat32 {
	return &BinaryHeapFloat32{
		less: less,
	}
}

// NewBinaryHeapFloat64 creates a new BinaryHeap
func NewBinaryHeapFloat64(less func(x, y float64) bool) *BinaryHeapFloat64 {
	return &BinaryHeapFloat64{
		less: less,
	}
}

// NewBinaryHeapComplex64 creates a new BinaryHeap
func NewBinaryHeapComplex64(less func(x, y complex64) bool) *BinaryHeapComplex64 {
	return &BinaryHeapComplex64{
		less: less,
	}
}

// NewBinaryHeapComplex128 creates a new BinaryHeap
func NewBinaryHeapComplex128(less func(x, y complex128) bool) *BinaryHeapComplex128 {
	return &BinaryHeapComplex128{
		less: less,
	}
}

// NewBinaryHeapString creates a new BinaryHeap
func NewBinaryHeapString(less func(x, y string) bool) *BinaryHeapString {
	return &BinaryHeapString{
		less: less,
	}
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapInt) Peek() (el int, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapInt8) Peek() (el int8, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapInt16) Peek() (el int16, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapInt32) Peek() (el int32, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapInt64) Peek() (el int64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapUint) Peek() (el uint, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapUint8) Peek() (el uint8, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapUint16) Peek() (el uint16, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapUint32) Peek() (el uint32, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapUint64) Peek() (el uint64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapFloat32) Peek() (el float32, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapFloat64) Peek() (el float64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapComplex64) Peek() (el complex64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapComplex128) Peek() (el complex128, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *BinaryHeapString) Peek() (el string, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}
	return h.elements[0], true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapInt) Pop() (el int, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapInt8) Pop() (el int8, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapInt16) Pop() (el int16, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapInt32) Pop() (el int32, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapInt64) Pop() (el int64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapUint) Pop() (el uint, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapUint8) Pop() (el uint8, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapUint16) Pop() (el uint16, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapUint32) Pop() (el uint32, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapUint64) Pop() (el uint64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapFloat32) Pop() (el float32, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapFloat64) Pop() (el float64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapComplex64) Pop() (el complex64, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapComplex128) Pop() (el complex128, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Pop removes the minumum element from the heap
// O(log n)
func (h *BinaryHeapString) Pop() (el string, ok bool) {
	if len(h.elements) == 0 {
		return el, false
	}

	el = h.elements[0]

	h.elements[0] = h.elements[len(h.elements)-1]
	h.elements = h.elements[:len(h.elements)-1]

	i := 0
	for {
		l, r := h.leftIndex(i), h.rightIndex(i)

		if l >= len(h.elements) {
			break
		}

		min := l

		if r < len(h.elements) && !h.less(h.elements[l], h.elements[r]) {
			min = r
		}

		if h.less(h.elements[min], h.elements[i]) {

			h.elements[i], h.elements[min] = h.elements[min], h.elements[i]
			i = min
		} else {

			break
		}
	}

	return el, true
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapInt) Push(el int) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapInt8) Push(el int8) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapInt16) Push(el int16) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapInt32) Push(el int32) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapInt64) Push(el int64) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapUint) Push(el uint) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapUint8) Push(el uint8) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapUint16) Push(el uint16) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapUint32) Push(el uint32) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapUint64) Push(el uint64) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapFloat32) Push(el float32) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapFloat64) Push(el float64) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapComplex64) Push(el complex64) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapComplex128) Push(el complex128) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

// Push pushes the element onto the heap
// O(log n)
func (h *BinaryHeapString) Push(el string) {
	i := len(h.elements)

	h.elements = append(h.elements, el)

	for i > 0 {
		p := h.parentIndex(i)
		if !h.less(h.elements[i], h.elements[p]) {
			break
		}
		h.elements[i], h.elements[p] = h.elements[p], h.elements[i]
		i = p
	}
}

func (h *BinaryHeapInt) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapInt8) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapInt16) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapInt32) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapInt64) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapUint) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapUint8) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapUint16) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapUint32) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapUint64) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapFloat32) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapFloat64) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapComplex64) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapComplex128) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapString) parentIndex(i int) int {
	return (i - 1) / 2
}

func (h *BinaryHeapInt) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapInt8) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapInt16) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapInt32) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapInt64) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapUint) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapUint8) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapUint16) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapUint32) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapUint64) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapFloat32) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapFloat64) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapComplex64) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapComplex128) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapString) leftIndex(i int) int {
	return (2 * i) + 1
}

func (h *BinaryHeapInt) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapInt8) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapInt16) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapInt32) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapInt64) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapUint) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapUint8) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapUint16) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapUint32) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapUint64) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapFloat32) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapFloat64) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapComplex64) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapComplex128) rightIndex(i int) int {
	return (2 * i) + 2
}

func (h *BinaryHeapString) rightIndex(i int) int {
	return (2 * i) + 2
}
