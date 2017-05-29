package sparse

import (
	"reflect"
	"testing"

	"github.com/badgerodon/goreify/generics"
)

func Test(t *testing.T) {
	t.Run("8", func(t *testing.T) {
		arr := NewArray(8)
		for i := 0; i < 8; i++ {
			arr.Set(i, i)
		}

		if !reflect.DeepEqual([]generics.T1{0, 1, 2, 3, 4, 5, 6, 7}, arr.Slice()) {
			t.Errorf("expected 8 items to be in the sparse array, got %v", arr.Slice())
		}
		if arr.Len() != 8 {
			t.Errorf("expected length to be 8, got %v", arr.Len())
		}

		arr.Delete(2)
		arr.Delete(4)

		if !reflect.DeepEqual([]generics.T1{0, 1, 3, 5, 6, 7}, arr.Slice()) {
			t.Errorf("expected 6 items to be in the sparse array, got %v", arr.Slice())
		}
		if arr.Len() != 6 {
			t.Errorf("expected length to be 6, got %v", arr.Len())
		}
	})

	t.Run("64", func(t *testing.T) {
		arr := NewArray(64)
		arr.Set(7, 17)
		arr.Set(63, 163)

		for ek, ev := range map[int]int{7: 17, 63: 163} {
			v, ok := arr.Get(ek)
			if !ok {
				t.Errorf("expected to get back %v for %v", ev, ek)
			}
			vi, ok := v.(int)
			if !ok {
				t.Errorf("expected to get back an int, got %T", v)
			}
			if vi != ev {
				t.Errorf("expected %v, got %v", ev, vi)
			}
		}

		_, ok := arr.Get(51)
		if ok {
			t.Errorf("expected get on a non-existent key to return false")
		}

		if !reflect.DeepEqual([]generics.T1{17, 163}, arr.Slice()) {
			t.Errorf("expected 2 items to be in the sparse array, got %v", arr.Slice())
		}
	})
}
