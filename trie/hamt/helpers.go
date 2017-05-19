package hamt

func getIndex(hash uint32, level uint32) uint32 {
	offset := w * level
	mask := uint32((1<<w - 1) << offset)
	idx := (hash & mask) >> offset
	return idx
}

// Copyright 2011 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

func fnv32a(data []byte) uint32 {
	offset := uint32(2166136261)
	prime := uint32(16777619)
	hash := offset
	for _, c := range data {
		hash ^= uint32(c)
		hash *= prime
	}
	return hash
}

// From Hacker's Delight, fig 5.2.
func popcountHD(x uint32) int {
	x -= (x >> 1) & 0x55555555
	x = (x & 0x33333333) + ((x >> 2) & 0x33333333)
	x = (x + (x >> 4)) & 0x0f0f0f0f
	x = x + (x >> 8)
	x = x + (x >> 16)
	return int(x & 0x0000003f)
}
