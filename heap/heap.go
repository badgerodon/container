package heap

import "github.com/badgerodon/goreify/generics"

//go:generate goreify github.com/badgerodon/container/heap.Heap numeric,string

// A Heap is an effient data structure to find the minimum element in a collection
type Heap interface {
	Peek() (generics.T1, bool)
	Pop() (generics.T1, bool)
	Push(generics.T1)
}
