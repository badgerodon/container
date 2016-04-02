package container

import "testing"

func TestUnrolledList(t *testing.T) {
	u := NewUnrolledList(10)

	for i := 0; i < 100; i++ {
		u.Add(i)
	}

	if u.Len() != 100 {
		t.Fatalf("length should be 100, got: %v", u.Len())
	}
	if !u.Contains(16) {
		t.Fatalf("list should contain 16")
	}
	if u.IndexOf(16) != 16 {
		t.Fatalf("index of 16 should be 16, got: %v, in: %v", u.IndexOf(16), u)
	}
	if u.Get(16) != 16 {
		t.Fatalf("expected 16, got: %v", u.Get(16))
	}

	u.Clear()

	if u.Get(16) != nil {
		t.Fatalf("expected nil, got: %v", u.Get(16))
	}

	u.Insert(11, 1)
	if u.Get(11) != 1 {
		t.Fatalf("expected 1, got: %v", u.Get(11))
	}

	u.Remove(1)
	if u.Get(1) != nil {
		t.Fatalf("expected nil, got: %v", u.Get(1))
	}
	if u.Len() != 11 {
		t.Fatalf("expected len to be 11, got: %v", u.Len())
	}
}
