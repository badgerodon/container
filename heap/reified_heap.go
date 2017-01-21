package heap

type

// A HeapInt is an effient data structure to find the minimum element in a collection
HeapInt interface {
	Peek() (int, bool)
	Pop() (int, bool)
	Push(int)
}

type

// A HeapInt8 is an effient data structure to find the minimum element in a collection
HeapInt8 interface {
	Peek() (int8, bool)
	Pop() (int8, bool)
	Push(int8)
}

type

// A HeapInt16 is an effient data structure to find the minimum element in a collection
HeapInt16 interface {
	Peek() (int16, bool)
	Pop() (int16, bool)
	Push(int16)
}

type

// A HeapInt32 is an effient data structure to find the minimum element in a collection
HeapInt32 interface {
	Peek() (int32, bool)
	Pop() (int32, bool)
	Push(int32)
}

type

// A HeapInt64 is an effient data structure to find the minimum element in a collection
HeapInt64 interface {
	Peek() (int64, bool)
	Pop() (int64, bool)
	Push(int64)
}

type

// A HeapUint is an effient data structure to find the minimum element in a collection
HeapUint interface {
	Peek() (uint, bool)
	Pop() (uint, bool)
	Push(uint)
}

type

// A HeapUint8 is an effient data structure to find the minimum element in a collection
HeapUint8 interface {
	Peek() (uint8, bool)
	Pop() (uint8, bool)
	Push(uint8)
}

type

// A HeapUint16 is an effient data structure to find the minimum element in a collection
HeapUint16 interface {
	Peek() (uint16, bool)
	Pop() (uint16, bool)
	Push(uint16)
}

type

// A HeapUint32 is an effient data structure to find the minimum element in a collection
HeapUint32 interface {
	Peek() (uint32, bool)
	Pop() (uint32, bool)
	Push(uint32)
}

type

// A HeapUint64 is an effient data structure to find the minimum element in a collection
HeapUint64 interface {
	Peek() (uint64, bool)
	Pop() (uint64, bool)
	Push(uint64)
}

type

// A HeapFloat32 is an effient data structure to find the minimum element in a collection
HeapFloat32 interface {
	Peek() (float32, bool)
	Pop() (float32, bool)
	Push(float32)
}

type

// A HeapFloat64 is an effient data structure to find the minimum element in a collection
HeapFloat64 interface {
	Peek() (float64, bool)
	Pop() (float64, bool)
	Push(float64)
}

type

// A HeapComplex64 is an effient data structure to find the minimum element in a collection
HeapComplex64 interface {
	Peek() (complex64, bool)
	Pop() (complex64, bool)
	Push(complex64)
}

type

// A HeapComplex128 is an effient data structure to find the minimum element in a collection
HeapComplex128 interface {
	Peek() (complex128, bool)
	Pop() (complex128, bool)
	Push(complex128)
}

type

// A HeapString is an effient data structure to find the minimum element in a collection
HeapString interface {
	Peek() (string, bool)
	Pop() (string, bool)
	Push(string)
}
