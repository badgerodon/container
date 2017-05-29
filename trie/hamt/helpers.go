package hamt

type (
	word uint64
)

const (
	bitsPerWord      = 64
	maxword     word = (1 << 64) - 1
)

func getArrayIndex(bitmap word, idx word) word {
	return popcount(bitmap & (maxword << idx))
}

func getBitmapIndexes(bitmap word) []word {
	var idxs []word
	for i := word(0); i < 64; i++ {
		if bitmap&(1<<i) > 0 {
			idxs = append(idxs, i)
		}
	}
	return idxs
}

func getHashIndex(hash word, w word, lvl word) word {
	offset := w * lvl
	mask := word(((1 << w) - 1) << offset)
	idx := (hash & mask) >> offset
	return idx
}

func maxLevel(w word) word {
	if 64%w == 0 {
		return 64 / w
	}
	return (64 / w) + 1
}

// func getIndex(hash uint32, level uint32) uint32 {
// 	offset := w * level
// 	mask := uint32((1<<w - 1) << offset)
// 	idx := (hash & mask) >> offset
// 	return idx
// }

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
