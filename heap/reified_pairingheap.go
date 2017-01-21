package heap

type

// A PairingHeapInt implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapInt struct {
	less func(int, int) bool
	root *pairingHeapNodeInt
}

type

// A PairingHeapInt8 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapInt8 struct {
	less func(int8, int8) bool
	root *pairingHeapNodeInt8
}

type

// A PairingHeapInt16 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapInt16 struct {
	less func(int16, int16) bool
	root *pairingHeapNodeInt16
}

type

// A PairingHeapInt32 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapInt32 struct {
	less func(int32, int32) bool
	root *pairingHeapNodeInt32
}

type

// A PairingHeapInt64 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapInt64 struct {
	less func(int64, int64) bool
	root *pairingHeapNodeInt64
}

type

// A PairingHeapUint implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapUint struct {
	less func(uint, uint) bool
	root *pairingHeapNodeUint
}

type

// A PairingHeapUint8 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapUint8 struct {
	less func(uint8, uint8) bool
	root *pairingHeapNodeUint8
}

type

// A PairingHeapUint16 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapUint16 struct {
	less func(uint16, uint16) bool
	root *pairingHeapNodeUint16
}

type

// A PairingHeapUint32 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapUint32 struct {
	less func(uint32, uint32) bool
	root *pairingHeapNodeUint32
}

type

// A PairingHeapUint64 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapUint64 struct {
	less func(uint64, uint64) bool
	root *pairingHeapNodeUint64
}

type

// A PairingHeapFloat32 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapFloat32 struct {
	less func(float32, float32) bool
	root *pairingHeapNodeFloat32
}

type

// A PairingHeapFloat64 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapFloat64 struct {
	less func(float64, float64) bool
	root *pairingHeapNodeFloat64
}

type

// A PairingHeapComplex64 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapComplex64 struct {
	less func(complex64, complex64) bool
	root *pairingHeapNodeComplex64
}

type

// A PairingHeapComplex128 implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapComplex128 struct {
	less func(complex128, complex128) bool
	root *pairingHeapNodeComplex128
}

type

// A PairingHeapString implements the Heap interface using the Pairing Heap data structure
// see: wikipedia.org/wiki/Pairing_heap
PairingHeapString struct {
	less func(string, string) bool
	root *pairingHeapNodeString
}

type pairingHeapNodeInt struct {
	elem                   int
	parent, child, sibling *pairingHeapNodeInt
}

type pairingHeapNodeInt8 struct {
	elem                   int8
	parent, child, sibling *pairingHeapNodeInt8
}

type pairingHeapNodeInt16 struct {
	elem                   int16
	parent, child, sibling *pairingHeapNodeInt16
}

type pairingHeapNodeInt32 struct {
	elem                   int32
	parent, child, sibling *pairingHeapNodeInt32
}

type pairingHeapNodeInt64 struct {
	elem                   int64
	parent, child, sibling *pairingHeapNodeInt64
}

type pairingHeapNodeUint struct {
	elem                   uint
	parent, child, sibling *pairingHeapNodeUint
}

type pairingHeapNodeUint8 struct {
	elem                   uint8
	parent, child, sibling *pairingHeapNodeUint8
}

type pairingHeapNodeUint16 struct {
	elem                   uint16
	parent, child, sibling *pairingHeapNodeUint16
}

type pairingHeapNodeUint32 struct {
	elem                   uint32
	parent, child, sibling *pairingHeapNodeUint32
}

type pairingHeapNodeUint64 struct {
	elem                   uint64
	parent, child, sibling *pairingHeapNodeUint64
}

type pairingHeapNodeFloat32 struct {
	elem                   float32
	parent, child, sibling *pairingHeapNodeFloat32
}

type pairingHeapNodeFloat64 struct {
	elem                   float64
	parent, child, sibling *pairingHeapNodeFloat64
}

type pairingHeapNodeComplex64 struct {
	elem                   complex64
	parent, child, sibling *pairingHeapNodeComplex64
}

type pairingHeapNodeComplex128 struct {
	elem                   complex128
	parent, child, sibling *pairingHeapNodeComplex128
}

type pairingHeapNodeString struct {
	elem                   string
	parent, child, sibling *pairingHeapNodeString
}

// NewPairingHeapInt creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapInt(less func(int, int) bool) *PairingHeapInt {
	return &PairingHeapInt{
		less: less,
	}
}

// NewPairingHeapInt8 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapInt8(less func(int8, int8) bool) *PairingHeapInt8 {
	return &PairingHeapInt8{
		less: less,
	}
}

// NewPairingHeapInt16 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapInt16(less func(int16, int16) bool) *PairingHeapInt16 {
	return &PairingHeapInt16{
		less: less,
	}
}

// NewPairingHeapInt32 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapInt32(less func(int32, int32) bool) *PairingHeapInt32 {
	return &PairingHeapInt32{
		less: less,
	}
}

// NewPairingHeapInt64 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapInt64(less func(int64, int64) bool) *PairingHeapInt64 {
	return &PairingHeapInt64{
		less: less,
	}
}

// NewPairingHeapUint creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapUint(less func(uint, uint) bool) *PairingHeapUint {
	return &PairingHeapUint{
		less: less,
	}
}

// NewPairingHeapUint8 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapUint8(less func(uint8, uint8) bool) *PairingHeapUint8 {
	return &PairingHeapUint8{
		less: less,
	}
}

// NewPairingHeapUint16 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapUint16(less func(uint16, uint16) bool) *PairingHeapUint16 {
	return &PairingHeapUint16{
		less: less,
	}
}

// NewPairingHeapUint32 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapUint32(less func(uint32, uint32) bool) *PairingHeapUint32 {
	return &PairingHeapUint32{
		less: less,
	}
}

// NewPairingHeapUint64 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapUint64(less func(uint64, uint64) bool) *PairingHeapUint64 {
	return &PairingHeapUint64{
		less: less,
	}
}

// NewPairingHeapFloat32 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapFloat32(less func(float32, float32) bool) *PairingHeapFloat32 {
	return &PairingHeapFloat32{
		less: less,
	}
}

// NewPairingHeapFloat64 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapFloat64(less func(float64, float64) bool) *PairingHeapFloat64 {
	return &PairingHeapFloat64{
		less: less,
	}
}

// NewPairingHeapComplex64 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapComplex64(less func(complex64, complex64) bool) *PairingHeapComplex64 {
	return &PairingHeapComplex64{
		less: less,
	}
}

// NewPairingHeapComplex128 creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapComplex128(less func(complex128, complex128) bool) *PairingHeapComplex128 {
	return &PairingHeapComplex128{
		less: less,
	}
}

// NewPairingHeapString creates a new PairingHeap using the less function to define
// the ordering of the elements
func NewPairingHeapString(less func(string, string) bool) *PairingHeapString {
	return &PairingHeapString{
		less: less,
	}
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapInt) Peek() (el int, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapInt8) Peek() (el int8, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapInt16) Peek() (el int16, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapInt32) Peek() (el int32, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapInt64) Peek() (el int64, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapUint) Peek() (el uint, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapUint8) Peek() (el uint8, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapUint16) Peek() (el uint16, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapUint32) Peek() (el uint32, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapUint64) Peek() (el uint64, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapFloat32) Peek() (el float32, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapFloat64) Peek() (el float64, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapComplex64) Peek() (el complex64, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapComplex128) Peek() (el complex128, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Peek returns the minimum element in the heap
// O(1)
func (h *PairingHeapString) Peek() (el string, ok bool) {
	if h.root == nil {
		return el, false
	}
	return h.root.elem, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapInt) Pop() (el int, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapInt8) Pop() (el int8, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapInt16) Pop() (el int16, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapInt32) Pop() (el int32, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapInt64) Pop() (el int64, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapUint) Pop() (el uint, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapUint8) Pop() (el uint8, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapUint16) Pop() (el uint16, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapUint32) Pop() (el uint32, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapUint64) Pop() (el uint64, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapFloat32) Pop() (el float32, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapFloat64) Pop() (el float64, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapComplex64) Pop() (el complex64, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapComplex128) Pop() (el complex128, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Pop returns the minimum element in the heap and removes it
// O(log n)
func (h *PairingHeapString) Pop() (el string, ok bool) {
	if h.root == nil {
		return el, false
	}
	el = h.root.elem
	h.root = h.mergePairs(h.root.child)
	return el, true
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapInt) Push(el int) {
	h.root = h.merge(h.root, &pairingHeapNodeInt{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapInt8) Push(el int8) {
	h.root = h.merge(h.root, &pairingHeapNodeInt8{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapInt16) Push(el int16) {
	h.root = h.merge(h.root, &pairingHeapNodeInt16{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapInt32) Push(el int32) {
	h.root = h.merge(h.root, &pairingHeapNodeInt32{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapInt64) Push(el int64) {
	h.root = h.merge(h.root, &pairingHeapNodeInt64{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapUint) Push(el uint) {
	h.root = h.merge(h.root, &pairingHeapNodeUint{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapUint8) Push(el uint8) {
	h.root = h.merge(h.root, &pairingHeapNodeUint8{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapUint16) Push(el uint16) {
	h.root = h.merge(h.root, &pairingHeapNodeUint16{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapUint32) Push(el uint32) {
	h.root = h.merge(h.root, &pairingHeapNodeUint32{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapUint64) Push(el uint64) {
	h.root = h.merge(h.root, &pairingHeapNodeUint64{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapFloat32) Push(el float32) {
	h.root = h.merge(h.root, &pairingHeapNodeFloat32{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapFloat64) Push(el float64) {
	h.root = h.merge(h.root, &pairingHeapNodeFloat64{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapComplex64) Push(el complex64) {
	h.root = h.merge(h.root, &pairingHeapNodeComplex64{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapComplex128) Push(el complex128) {
	h.root = h.merge(h.root, &pairingHeapNodeComplex128{
		elem: el,
	})
}

// Push adds the element to the heap
// O(1)
func (h *PairingHeapString) Push(el string) {
	h.root = h.merge(h.root, &pairingHeapNodeString{
		elem: el,
	})
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapInt) Merge(h2 *PairingHeapInt) *PairingHeapInt {
	return &PairingHeapInt{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapInt8) Merge(h2 *PairingHeapInt8) *PairingHeapInt8 {
	return &PairingHeapInt8{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapInt16) Merge(h2 *PairingHeapInt16) *PairingHeapInt16 {
	return &PairingHeapInt16{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapInt32) Merge(h2 *PairingHeapInt32) *PairingHeapInt32 {
	return &PairingHeapInt32{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapInt64) Merge(h2 *PairingHeapInt64) *PairingHeapInt64 {
	return &PairingHeapInt64{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapUint) Merge(h2 *PairingHeapUint) *PairingHeapUint {
	return &PairingHeapUint{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapUint8) Merge(h2 *PairingHeapUint8) *PairingHeapUint8 {
	return &PairingHeapUint8{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapUint16) Merge(h2 *PairingHeapUint16) *PairingHeapUint16 {
	return &PairingHeapUint16{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapUint32) Merge(h2 *PairingHeapUint32) *PairingHeapUint32 {
	return &PairingHeapUint32{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapUint64) Merge(h2 *PairingHeapUint64) *PairingHeapUint64 {
	return &PairingHeapUint64{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapFloat32) Merge(h2 *PairingHeapFloat32) *PairingHeapFloat32 {
	return &PairingHeapFloat32{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapFloat64) Merge(h2 *PairingHeapFloat64) *PairingHeapFloat64 {
	return &PairingHeapFloat64{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapComplex64) Merge(h2 *PairingHeapComplex64) *PairingHeapComplex64 {
	return &PairingHeapComplex64{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapComplex128) Merge(h2 *PairingHeapComplex128) *PairingHeapComplex128 {
	return &PairingHeapComplex128{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

// Merge merges two heaps
// O(1)
func (h *PairingHeapString) Merge(h2 *PairingHeapString) *PairingHeapString {
	return &PairingHeapString{
		less: h.less,
		root: h.merge(h.root, h2.root),
	}
}

func (h *PairingHeapInt) merge(n1, n2 *pairingHeapNodeInt) *pairingHeapNodeInt {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapInt8) merge(n1, n2 *pairingHeapNodeInt8) *pairingHeapNodeInt8 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapInt16) merge(n1, n2 *pairingHeapNodeInt16) *pairingHeapNodeInt16 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapInt32) merge(n1, n2 *pairingHeapNodeInt32) *pairingHeapNodeInt32 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapInt64) merge(n1, n2 *pairingHeapNodeInt64) *pairingHeapNodeInt64 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapUint) merge(n1, n2 *pairingHeapNodeUint) *pairingHeapNodeUint {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapUint8) merge(n1, n2 *pairingHeapNodeUint8) *pairingHeapNodeUint8 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapUint16) merge(n1, n2 *pairingHeapNodeUint16) *pairingHeapNodeUint16 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapUint32) merge(n1, n2 *pairingHeapNodeUint32) *pairingHeapNodeUint32 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapUint64) merge(n1, n2 *pairingHeapNodeUint64) *pairingHeapNodeUint64 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapFloat32) merge(n1, n2 *pairingHeapNodeFloat32) *pairingHeapNodeFloat32 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapFloat64) merge(n1, n2 *pairingHeapNodeFloat64) *pairingHeapNodeFloat64 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapComplex64) merge(n1, n2 *pairingHeapNodeComplex64) *pairingHeapNodeComplex64 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapComplex128) merge(n1, n2 *pairingHeapNodeComplex128) *pairingHeapNodeComplex128 {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapString) merge(n1, n2 *pairingHeapNodeString) *pairingHeapNodeString {
	switch {
	case n1 == nil:
		return n2
	case n2 == nil:
		return n1
	case h.less(n1.elem, n2.elem):
		c1 := n1.child
		n1.child = n2
		n2.parent = n1
		n2.sibling = c1
		return n1
	default:
		c2 := n2.child
		n2.child = n1
		n1.parent = n2
		n1.sibling = c2
		return n2
	}
}

func (h *PairingHeapInt) mergePairs(n *pairingHeapNodeInt) *pairingHeapNodeInt {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapInt8) mergePairs(n *pairingHeapNodeInt8) *pairingHeapNodeInt8 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapInt16) mergePairs(n *pairingHeapNodeInt16) *pairingHeapNodeInt16 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapInt32) mergePairs(n *pairingHeapNodeInt32) *pairingHeapNodeInt32 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapInt64) mergePairs(n *pairingHeapNodeInt64) *pairingHeapNodeInt64 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapUint) mergePairs(n *pairingHeapNodeUint) *pairingHeapNodeUint {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapUint8) mergePairs(n *pairingHeapNodeUint8) *pairingHeapNodeUint8 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapUint16) mergePairs(n *pairingHeapNodeUint16) *pairingHeapNodeUint16 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapUint32) mergePairs(n *pairingHeapNodeUint32) *pairingHeapNodeUint32 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapUint64) mergePairs(n *pairingHeapNodeUint64) *pairingHeapNodeUint64 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapFloat32) mergePairs(n *pairingHeapNodeFloat32) *pairingHeapNodeFloat32 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapFloat64) mergePairs(n *pairingHeapNodeFloat64) *pairingHeapNodeFloat64 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapComplex64) mergePairs(n *pairingHeapNodeComplex64) *pairingHeapNodeComplex64 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapComplex128) mergePairs(n *pairingHeapNodeComplex128) *pairingHeapNodeComplex128 {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}

func (h *PairingHeapString) mergePairs(n *pairingHeapNodeString) *pairingHeapNodeString {
	switch {
	case n == nil:
		return nil
	case n.sibling == nil:
		return n
	default:
		cousin := n.sibling.sibling
		return h.merge(h.merge(n, n.sibling), h.mergePairs(cousin))
	}
}
