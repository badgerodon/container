package heap

import "testing"

func testLess(x, y int) bool {
	return x < y
}

func TestHeaps(t *testing.T) {

	heapMakers := []func() HeapInt{
		func() HeapInt { return NewPairingHeapInt(testLess) },
		func() HeapInt { return NewBinaryHeapInt(testLess) },
	}

	for _, heapMaker := range heapMakers {
		h := heapMaker()

		h.Push(3)
		h.Push(8)
		h.Push(11)
		h.Push(10000)
		h.Push(17)
		h.Push(6)
		h.Push(1)

		for _, expected := range []int{1, 3, 6, 8, 11, 17, 10000} {
			if actual, _ := h.Peek(); actual != expected {
				t.Errorf("[%T] expected %v to be returned, but got %v", h, expected, actual)
			}
			if actual, _ := h.Pop(); actual != expected {
				t.Errorf("[%T] expected %v to be returned, but got %v", h, expected, actual)
			}
		}

		if _, ok := h.Pop(); ok {
			t.Errorf("[%T] expected Pop to return false when the collection is empty", h)
		}
	}

	{
		h1, h2 := NewPairingHeapInt(testLess), NewPairingHeapInt(testLess)

		h1.Push(3)
		h1.Push(1)
		h2.Push(2)
		h2.Push(4)

		h3 := h1.Merge(h2)

		for _, expected := range []int{1, 2, 3, 4} {
			if actual, _ := h3.Peek(); actual != expected {
				t.Errorf("expected %v to be returned, but got %v", expected, actual)
			}
			if actual, _ := h3.Pop(); actual != expected {
				t.Errorf("expected %v to be returned, but got %v", expected, actual)
			}
		}
	}
}

const testHeapSize = 10 * 1024

func benchmarkPushPop(h HeapInt, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < testHeapSize; j++ {
			h.Push(j)
		}
		for j := 0; j < testHeapSize; j++ {
			h.Pop()
		}
	}
}
func benchmarkPushPeek(h HeapInt, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < testHeapSize; j++ {
			h.Push(j)
		}
		for j := 0; j < testHeapSize; j++ {
			h.Peek()
		}
	}
}

func BenchmarkBinaryHeap(b *testing.B) {
	b.Run("Push and Peek", func(b *testing.B) {
		b.SetBytes(testHeapSize * 8)
		benchmarkPushPeek(NewBinaryHeapInt(testLess), b.N)
	})
	b.Run("Push and Pop", func(b *testing.B) {
		b.SetBytes(testHeapSize * 8)
		benchmarkPushPop(NewBinaryHeapInt(testLess), b.N)
	})
}

func BenchmarkPairingHeap(b *testing.B) {
	b.Run("Push and Peek", func(b *testing.B) {
		b.SetBytes(testHeapSize * 8)
		benchmarkPushPeek(NewPairingHeapInt(testLess), b.N)
	})
	b.Run("Push and Pop", func(b *testing.B) {
		b.SetBytes(testHeapSize * 8)
		benchmarkPushPop(NewPairingHeapInt(testLess), b.N)
	})
}
