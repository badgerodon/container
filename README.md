# container

[![GoDoc](https://godoc.org/github.com/badgerodon/container?status.png)](http://godoc.org/github.com/badgerodon/container)

This is a collection of type-safe data structures in Go. Each data structure contains a default, `interface{}` implementation which is used as a template to generate type-safe versions by the `generate` program.

The convention is: the type `MyContainer` will contain `interface{}` items, `MyContainerOfT` will contain `T` items. The type `T` will be upper-cased and there is a special rule for `[]byte` to be named `Bytes`.

For containers that behave similar to a map, the convention is `MyContainerOfTKeyToTValue`, for example `MyContainerOfIntToString`.

You can use the `generate` program to generate your own versions:

    //go:generate go run github.com/badgerodon/container/generate/ -package yourpackage -template arraylist -destination myarraylist.go -types MyType

If you're familiar with generics you're probably already aware that there are limitations to this approach which can make it hard to generalize a data structure. For example equality and comparisons may not be well defined for the type in question. (though it does work well for basic structs and builtin types) I have provided one override which converts `a == b` for `[]byte` to `bytes.Equal(a, b)`. Theoretically more could be supported.

Wherever possible I've tried to name methods consistently so it's easy to abstract from a particular implementation.

## Lists

### ArrayList

An `ArrayList` is a dynamic, growable array. This is basically equivalent to a slice.

### LinkedList

A `LinkedList` stores each item in a node which contains forward/backwards pointers to other nodes.

### UnrolledList

An `UnrolledList` is dynamic, growable array of arrays. (Basically a `LinkedList` of arrays)

Only sections (size defined by arraySize) of the list will be contiguous. The amortized O(1) cost of appending an element to an ArrayList degrades if the ArrayList is large, since copies and re-allocation can become very expensive and the memory usage can grow out of hand. An `UnrolledList` avoids this issue by always allocating fixed-size chunks, and not having to copy existing data. The downside is an added indirection cost, and incompatibility with anything that needs a real array. 